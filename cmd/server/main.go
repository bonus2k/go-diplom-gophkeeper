package main

import (
	"github.com/bonus2k/go-diplom-gophkeeper/internal/database"
	"github.com/bonus2k/go-diplom-gophkeeper/internal/logger"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	log *logger.Logger
)

func main() {
	initLogger()
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	store := database.NewDataStore(log, db)
	err = store.Migrate()
	if err != nil {
		log.Fatal("failed to migrate database", err)
	}
}

func initLogger() {
	l := logrus.New()
	log = logger.NewLogger(l)
}
