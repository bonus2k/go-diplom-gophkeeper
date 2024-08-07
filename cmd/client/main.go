package main

import (
	"github.com/bonus2k/go-diplom-gophkeeper/internal/logger"
	"github.com/bonus2k/go-diplom-gophkeeper/internal/mvc"
	"github.com/bonus2k/go-diplom-gophkeeper/internal/services/note_service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	log *logger.Logger
)

func main() {

	initLogger()
	conn, err := grpc.NewClient(":3200", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Error(err)
		}
	}(conn)

	serviceNote := note_service.NewUIService(log, conn)
	ui := mvc.NewUIController(log, serviceNote)
	ui.AddItemInfoList("The application has started successfully. Hello! 😁")
	err = ui.Run()
	if err != nil {
		log.Fatal("start controller ui is fail", err)
	}
}

func initLogger() {
	l := logrus.New()
	log = logger.NewLogger(l)
}
