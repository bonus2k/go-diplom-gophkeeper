package main

import (
	"log"
	"net"
	"os"

	"github.com/bonus2k/go-diplom-gophkeeper/internal/database"
	pb "github.com/bonus2k/go-diplom-gophkeeper/internal/interfaces/proto"
	"github.com/bonus2k/go-diplom-gophkeeper/internal/interfaces/server"
	"github.com/bonus2k/go-diplom-gophkeeper/internal/logger"
	"github.com/bonus2k/go-diplom-gophkeeper/internal/services/auth"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	logg *logger.Logger
)

func main() {
	parseFlags()
	initLogger()
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	store := *database.NewDataStore(logg, db)
	err = store.Migrate()
	if err != nil {
		log.Fatal("failed to migrate database", err)
	}

	authService, err := auth.NewAuthService(logg, crtFile)
	if err != nil {
		log.Fatal("failed to initialize auth service", err)
	}

	controller := server.NewController(logg, store, *authService)
	listen, err := net.Listen("tcp", srvAddr)
	if err != nil {
		log.Fatal("failed start listener", err)
	}
	gServer := grpc.NewServer(grpc.UnaryInterceptor(server.TokenInterceptor))
	pb.RegisterNoteServicesServer(gServer, controller)
	pb.RegisterUserServicesServer(gServer, controller)
	logg.Info("server started")
	if err := gServer.Serve(listen); err != nil {
		log.Fatal("failed start server", err)
	}
}

func initLogger() {

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		log.Fatal(err)
	}

	l := &logrus.Logger{
		Out:   os.Stdout,
		Level: level,
		Formatter: &logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "@timestamp",
				logrus.FieldKeyLevel: "@level",
				logrus.FieldKeyMsg:   "@message",
				logrus.FieldKeyFunc:  "@caller",
			},
		},
	}
	logg = logger.NewLogger(l)
}
