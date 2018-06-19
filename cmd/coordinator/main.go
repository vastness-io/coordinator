package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mattes/migrate"
	_ "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
	"github.com/mattes/migrate/source/go-bindata"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/vastness-io/coordinator-svc/project"
	"github.com/vastness-io/coordinator/db/migration"
	"github.com/vastness-io/coordinator/pkg/repository"
	project_server "github.com/vastness-io/coordinator/pkg/server/project"
	event_server "github.com/vastness-io/coordinator/pkg/server/vcs_event"
	project_service "github.com/vastness-io/coordinator/pkg/service/project"
	event_service "github.com/vastness-io/coordinator/pkg/service/vcs_event"
	"github.com/vastness-io/gormer"
	"github.com/vastness-io/linguist-svc"
	toolkit "github.com/vastness-io/toolkit/pkg/grpc"
	vcswebhook "github.com/vastness-io/vcs-webhook-svc/webhook"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

const (
	name                     = "coordinator"
	description              = "Ensures components work together."
	supportedDb              = "postgres"
	dbName                   = "postgres"
	connectionStringTemplate = "host=%s port=%v user=%s dbname=%s password=%s sslmode=disable"
	migrationStringTemplate  = "postgres://%s:%s@%s:%v/%s?sslmode=disable"
)

var (
	log          = logrus.WithField("component", name)
	commit       string
	version      string
	addr         string
	port         int
	databaseHost string
	databasePort int
	databaseUser string
	databasePass string
	linguistSrv  string
	debugMode    bool
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

	log.Infof("Starting %s", name)

	var (
		address          = net.JoinHostPort(addr, strconv.Itoa(port))
		tracer           = opentracing.GlobalTracer()
		lis, err         = net.Listen("tcp", address)
		srv              = toolkit.NewGRPCServer(tracer, log)
		connectionString = fmt.Sprintf(connectionStringTemplate, databaseHost, databasePort, databaseUser, dbName, databasePass)
		migrationString  = fmt.Sprintf(migrationStringTemplate, databaseUser, databasePass, databaseHost, databasePort, dbName)
	)

	if err != nil {
		log.Fatal(err)
	}

	linguistConn, err := toolkit.NewGRPCClient(tracer, log, grpc.WithInsecure())(linguistSrv)

	if err != nil {
		log.Fatal(err)
	}

	defer linguistConn.Close()

	gormDB, err := gorm.Open(supportedDb, connectionString)
	defer gormDB.Close()

	if err != nil {
		log.Fatal(err)
	}

	s := bindata.Resource(migration.AssetNames(),
		func(name string) ([]byte, error) {
			return migration.Asset(name)
		})

	d, err := bindata.WithInstance(s)

	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithSourceInstance("go-bindata", d, migrationString)

	if err != nil {
		log.Fatal(err)
	}

	log.Infof("Running db migrations")

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	err1, err2 := m.Close()

	if err1 != nil {
		log.Fatalf("could not close migrate source: %v", err1)
	}
	if err2 != nil {
		log.Fatalf("could not close migrate database: %v", err2)
	}

	var (
		db                       = gormer.Wrap(gormDB)
		projectRepository        = repository.NewProjectRepository(db)
		linguistClient           = linguist.NewLinguistClient(linguistConn)
		vcsEventService          = event_service.NewVcsEventService(log.WithField("service", "vcs_event"), linguistClient, projectRepository)
		projectService           = project_service.NewProjectService(log.WithField("service", "project"), projectRepository)
		vcsEventServer           = event_server.NewVcsEventServer(vcsEventService, log.WithField("server", "vcs_event"))
		projectInformationServer = project_server.NewProjectInformationServer(projectService, log.WithField("server", "project"))
	)

	vcswebhook.RegisterVcsEventServer(srv, vcsEventServer)

	project.RegisterProjectsServer(srv, projectInformationServer)

	go func() {
		log.Infof("GRPC server listening on %s", address)
		if err := srv.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	signalChan := make(chan os.Signal, 1)

	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-signalChan:
			log.Infof("Exiting %s", name)
			srv.GracefulStop()
			os.Exit(0)
		}
	}
}
