package main

import (
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/vastness-io/coordinator/pkg/repository"
	"github.com/vastness-io/coordinator/pkg/server"
	"github.com/vastness-io/coordinator/pkg/service/event"
	"github.com/vastness-io/linguist-svc"
	toolkit "github.com/vastness-io/toolkit/pkg/grpc"
	vcswebhook "github.com/vastness-io/vcs-webhook-svc/webhook"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

const (
	name             = "coordinator"
	description      = "Ensures components work together."
	supportedDb      = "postgres"
	dbName           = "postgres"
	connectionString = "host=%s port=%v user=%s dbname=%s password=%s sslmode=disable"
)

var (
	log                   = logrus.WithField("component", name)
	commit                string
	version               string
	addr                  string
	port                  int
	databaseHost          string
	databasePort          int
	databaseUser          string
	databasePass          string
	migrationFileLocation string
	linguistSrv           string
	debugMode             bool
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
			EnvVar:      "COORDINATOR_ADDRESS",
		},
		cli.IntFlag{
			Name:        "port,p",
			Usage:       "Port to listen on",
			Value:       8080,
			Destination: &port,
			EnvVar:      "COORDINATOR_PORT",
		},
		cli.StringFlag{
			Name:        "database-host, db-host",
			Usage:       "Database connection host",
			Value:       "127.0.0.1",
			Destination: &databaseHost,
			EnvVar:      "DATABASE_HOST",
		},
		cli.IntFlag{
			Name:        "database-port, db-port",
			Usage:       "Database connection port",
			Value:       5432,
			Destination: &databasePort,
			EnvVar:      "DATABASE_PORT",
		},
		cli.StringFlag{
			Name:        "database-user, db-user",
			Usage:       "Database connection user",
			Destination: &databaseUser,
			EnvVar:      "DATABASE_USER",
		},
		cli.StringFlag{
			Name:        "database-pass, db-pass",
			Usage:       "Database connection password",
			Destination: &databasePass,
			EnvVar:      "DATABASE_PASS",
		},
		cli.StringFlag{
			Name:        "migration-file-location, migrate",
			Usage:       "Location of Database migration files",
			Value:       "/migration",
			Destination: &migrationFileLocation,
			EnvVar:      "MIGRATION_FILE_LOCATION",
		},
		cli.StringFlag{
			Name:        "linguist,s",
			Usage:       "Linguist Service address.",
			Value:       "127.0.0.1:8082",
			Destination: &linguistSrv,
			EnvVar:      "LINGUIST_ADDRESS",
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

	linguistConn, err := toolkit.NewGRPCClient(tracer, log, grpc.WithInsecure())(linguistSrv)

	if err != nil {
		log.Fatal(err)
	}

	defer linguistConn.Close()

	gormDB, err := gorm.Open(supportedDb, fmt.Sprintf(connectionString, databaseHost, databasePort, databaseUser, dbName, databasePass))
	defer gormDB.Close()

	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(gormDB.DB(), &postgres.Config{})

	log.Infof("Running migrations from %s", migrationFileLocation)

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationFileLocation),
		dbName, driver)

	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	var (
		db                = repository.NewDB(gormDB)
		projectRepository = repository.NewProjectRepository(db)
		linguistClient    = linguist.NewLinguistClient(linguistConn)
		vcsEventService   = event.NewVcsEventService(log, linguistClient, projectRepository)
		vcsEventServer    = server.NewVcsEventServer(vcsEventService, log)
	)

	vcswebhook.RegisterVcsEventServer(srv, vcsEventServer)

	grpc_prometheus.Register(srv)

	mux.Handle("/metrics", promhttp.Handler())

	go func() {
		log.Infof("GRPC server listening on %s", address)
		if err := srv.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

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
