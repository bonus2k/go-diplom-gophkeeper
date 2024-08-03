package main

import (
	"os"
	"time"

	"github.com/bonus2k/go-diplom-gophkeeper/internal/logger"
	"github.com/bonus2k/go-diplom-gophkeeper/internal/models"
	"github.com/bonus2k/go-diplom-gophkeeper/internal/mvc"
	"github.com/bonus2k/go-diplom-gophkeeper/internal/services/note_service"
	"github.com/sirupsen/logrus"
)

var (
	log *logger.Logger
)

func main() {
	done := make(chan os.Signal, 1)
	initLogger()
	serviceNote := note_service.NewServiceNote(log)
	ui := mvc.NewControllerUI(log, serviceNote)
	ui.AddItemInfoList("The application has started successfully. Hello! 😁")
	go func() {
		err := ui.Run()
		if err != nil {
			log.Fatal("start controller ui is fail", err)
		}
	}()

	storage := createTestStorage()
	for _, note := range storage {
		ui.AddNote(note)
	}

	go func() {
		storage := createTestStorage()
		for _, note := range storage {
			ui.AddNote(note)
		}
	}()
	<-done
}

func initLogger() {
	l := logrus.New()
	log = logger.NewLogger(l)
}

func createTestStorage() []models.Noteable {
	storage := make([]models.Noteable, 0)
	note1 := models.BankCardNote{
		Bank:         "Bank",
		Number:       "5556 4655 4655 4655 4655",
		Expiration:   "2022-04-01",
		Cardholder:   "John Smith",
		SecurityCode: "123",
	}
	note1.BaseNote = models.BaseNote{
		NameRecord: "Note1",
		Created:    time.Now().Unix(),
		Type:       models.CARD,
		MetaInfo:   []string{"test1 = test1\n", "test2 = test2\n", "сайт: www.test.com\n"},
	}

	note2 := models.BankCardNote{
		Bank:         "Bank",
		Number:       "5556 4611 4655 4655 4655",
		Expiration:   "2022-05-01",
		Cardholder:   "John Smith1",
		SecurityCode: "124",
	}
	note2.BaseNote = models.BaseNote{
		NameRecord: "Note2",
		Created:    time.Now().Unix(),
		Type:       models.CARD,
		MetaInfo:   []string{"test2 = test2\n", "test1 = test1\n", "test3 = test3\n"},
	}
	storage = append(storage, note1)
	storage = append(storage, note2)
	return storage
}
