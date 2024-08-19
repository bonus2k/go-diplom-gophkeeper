package main

import (
	"log"
	"os"

	"github.com/bonus2k/go-diplom-gophkeeper/internal/logger"
	"github.com/bonus2k/go-diplom-gophkeeper/internal/mvc"
	"github.com/bonus2k/go-diplom-gophkeeper/internal/services/ui"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	logg *logger.Logger
)

func main() {
	parseFlags()
	initLogger()
	conn, err := grpc.NewClient(connAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			logg.Error(err)
		}
	}(conn)

	serviceNote := ui.NewUIService(logg, conn)
	controller := mvc.NewUIController(logg, serviceNote)
	controller.AddItemInfoList("The application has started successfully. Hello! 😁")
	err = controller.Run()
	if err != nil {
		log.Fatal("start controller controller is fail", err)
	}
}

func initLogger() {
	var file *os.File
	if f, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err != nil {
		file = os.Stdout
	} else {
		file = f
	}

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		log.Fatal(err)
	}

	l := &logrus.Logger{
		Out:   file,
		Level: level,
		Formatter: &logrus.TextFormatter{
			DisableColors:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
		},
	}

	logg = logger.NewLogger(l)
}
