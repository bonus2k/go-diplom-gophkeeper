package note_service

import (
	"sync"

	"github.com/bonus2k/go-diplom-gophkeeper/internal/logger"
	"github.com/bonus2k/go-diplom-gophkeeper/internal/models"
)

var (
	once sync.Once
	log  *logger.Logger
	sn   *ServiceNote
)

type ServiceNote struct {
	storage []models.Noteable
}

func NewServiceNote(logger *logger.Logger) *ServiceNote {
	once.Do(func() {
		log = logger
		sn = &ServiceNote{storage: make([]models.Noteable, 0)}
	})
	return sn
}

func (cn *ServiceNote) AddNote(note models.Noteable) []models.Noteable {
	cn.storage = append(cn.storage, note)
	log.Infof("added new note: %s", note)
	return cn.storage
}
