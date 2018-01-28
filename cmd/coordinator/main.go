package main

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/vastness-io/coordinator/pkg/server"
	"github.com/vastness-io/coordinator/pkg/service/webhook"
	"github.com/vastness-io/linguist-svc"
	toolkit "github.com/vastness-io/toolkit/pkg/grpc"
	"github.com/vastness-io/vcs-webhook-svc/webhook/github"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

const (
	name        = "coordinator"
	description = "Ensures components work together."
)

var (
	log         = logrus.WithField("component", "coordinator")
	commit      string
	version     string
	addr        string
	port        int
	linguistSrv string
	debugMode   bool
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	app := cli.NewApp()
	app.Name = name
	app.Usage = description
	app.Version = fmt.Sprintf("%s (%s)", version, commit)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "addr,a",
			Usage:       "TCP address to listen on",
			Value:       "127.0.0.1",
			Destination: &addr,
		},
		cli.IntFlag{
			Name:        "port,p",
			Usage:       "Port to listen on",
			Value:       8080,
			Destination: &port,
		},
		cli.StringFlag{
			Name:        "linguist,s",
			Usage:       "Linguist Service address.",
			Value:       "127.0.0.1:8082",
			Destination: &linguistSrv,
		},
		cli.BoolFlag{
			Name:        "debug,d",
			Usage:       "Debug mode",
			Destination: &debugMode,
		},
	}
	app.Action = func(_ *cli.Context) { run() }
	app.Run(os.Args)
}

func run() {

	if debugMode {
		logrus.SetLevel(logrus.DebugLevel)
	}

	log.Info("Starting coordinator")

	var (
		mux      = http.NewServeMux()
		address  = net.JoinHostPort(addr, strconv.Itoa(port))
		tracer   = opentracing.GlobalTracer()
		lis, err = net.Listen("tcp", address)
		srv      = toolkit.NewGRPCServer(tracer, log)
	)

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		log.Infof("GRPC server listening on %s", address)
		if err := srv.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	linguistConn, err := toolkit.NewGRPCClient(tracer, log, grpc.WithInsecure())(linguistSrv)

	if err != nil {
		log.Fatal(err)
	}

	defer linguistConn.Close()

	var (
		linguistClient       = linguist.NewLinguistClient(linguistConn)
		githubWebhookService = webhook.NewGithubWebhookService(log, linguistClient)
		githubWebhookServer  = server.NewGithubWebhookServer(githubWebhookService, log)
	)

	github.RegisterGithubWebhookServer(srv, githubWebhookServer)

	grpc_prometheus.Register(srv)

	mux.Handle("/metrics", promhttp.Handler())

	go func() {
		log.Infof("HTTP server listening on %s", address)
		http.Serve(lis, mux)
	}()

	signalChan := make(chan os.Signal, 1)

	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-signalChan:
			log.Info("Exiting coordinator")
			srv.GracefulStop()
			os.Exit(0)
		}
	}
}
