package main

import (
	"net"

	"github.com/bonus2k/go-diplom-gophkeeper/internal/database"
	pb "github.com/bonus2k/go-diplom-gophkeeper/internal/interfaces/proto"
	"github.com/bonus2k/go-diplom-gophkeeper/internal/interfaces/server"
	"github.com/bonus2k/go-diplom-gophkeeper/internal/logger"
	"github.com/bonus2k/go-diplom-gophkeeper/internal/services/auth_service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	log *logger.Logger
)

func main() {
	initLogger()
	db, err := gorm.Open(sqlite.Open("dev_db.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	store := *database.NewDataStore(log, db)
	err = store.Migrate()
	if err != nil {
		log.Fatal("failed to migrate database", err)
	}

	authService, err := auth_service.NewAuthService(log, "private.pem")
	if err != nil {
		log.Fatal("failed to initialize auth service", err)
	}

	controller := server.NewController(log, store, *authService)
	listen, err := net.Listen("tcp", ":3200")
	if err != nil {
		log.Fatal("failed start listener", err)
	}
	gServer := grpc.NewServer(grpc.UnaryInterceptor(server.TokenInterceptor))
	pb.RegisterNoteServicesServer(gServer, controller)
	pb.RegisterUserServicesServer(gServer, controller)
	log.Info("server started")
	if err := gServer.Serve(listen); err != nil {
		log.Fatal("failed start server", err)
	}
}

func initLogger() {
	l := logrus.New()
	log = logger.NewLogger(l)
}
