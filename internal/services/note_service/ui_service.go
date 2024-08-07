package note_service

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	pb "github.com/bonus2k/go-diplom-gophkeeper/internal/interfaces/proto"
	"github.com/bonus2k/go-diplom-gophkeeper/internal/logger"
	"github.com/bonus2k/go-diplom-gophkeeper/internal/models"
	"github.com/bonus2k/go-diplom-gophkeeper/internal/util"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	once sync.Once
	log  *logger.Logger
	sn   *UIService
)

type UIService struct {
	storage map[uuid.UUID]*models.Noteable
	uc      pb.UserServicesClient
	nc      pb.NoteServicesClient
	jwt     string
	hash    []byte
}

func NewUIService(logger *logger.Logger, conn *grpc.ClientConn) *UIService {
	once.Do(func() {
		log = logger
		sn = &UIService{storage: make(map[uuid.UUID]*models.Noteable), uc: pb.NewUserServicesClient(conn), nc: pb.NewNoteServicesClient(conn)}
	})
	return sn
}

func (cn *UIService) AddNote(note models.Noteable) (*[]models.Noteable, error) {
	log := log.WithFields(logrus.Fields{
		"method": "AddNote",
	})

	if cn.jwt == "" {
		log.Warning("AddNote: jwt not found")
		return nil, fmt.Errorf("You need sigin to app")
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2000*time.Millisecond)
	defer cancelFunc()

	marshal, err := json.Marshal(note)
	if err != nil {
		log.WithError(err).Error("Error marshalling note")
		return nil, err
	}

	encrypt, err := util.Encrypt(ctx, cn.hash, marshal)
	if err != nil {
		log.WithError(err).Error("Error encrypting note")
		return nil, err
	}

	noteDto := &pb.Note{
		Id:         note.GetId().String(),
		Name:       note.GetName(),
		Type:       note.GetType().String(),
		SecretData: encrypt,
	}

	ctx = cn.addToken(ctx)
	switch _, ok := cn.storage[note.GetId()]; ok {
	case true:
		_, err = cn.nc.UpdateNote(ctx, noteDto)
		if err != nil {
			log.WithError(err).Error("Error updating note")
			return nil, err
		}
	case false:
		_, err = cn.nc.AddNote(ctx, noteDto)
		if err != nil {
			log.WithError(err).Error("Error adding note")
			return nil, err
		}
	}

	cn.storage[note.GetId()] = &note
	log.WithField("note", note.GetName()).Info("Added note")
	return toNotableList(cn.storage), nil
}

func (cn *UIService) LoadNote() (*[]models.Noteable, error) {
	log := log.WithFields(logrus.Fields{
		"method": "AddNote",
	})

	if cn.jwt == "" {
		log.Warning("AddNote: jwt not found")
		return nil, fmt.Errorf("You need sigin to app")
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2000*time.Millisecond)
	defer cancelFunc()
	ctx = cn.addToken(ctx)
	notes, err := cn.nc.GetNotes(ctx, &pb.NoteRequest{})
	if err != nil {
		log.WithError(err).Error("Error getting notes")
		return nil, err
	}
	for _, noteDto := range notes.Notes {
		note, err := unmarshalNote(ctx, cn.hash, noteDto)
		if err != nil {
			log.WithError(err).Error("Error unmarshalling note")
			continue
		}
		cn.storage[note.GetId()] = &note
	}

	return toNotableList(cn.storage), nil
}

func (cn *UIService) DeleteNote(id uuid.UUID) (*[]models.Noteable, error) {
	log := log.WithFields(logrus.Fields{
		"method": "DeleteNote",
	})

	if cn.jwt == "" {
		log.Warning("DeleteNote: jwt not found")
		return nil, fmt.Errorf("You need sigin to app")
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 2000*time.Millisecond)
	defer cancelFunc()
	ctx = cn.addToken(ctx)
	_, err := cn.nc.DeleteNote(ctx, &pb.NoteRequest{IdNote: id.String()})
	if err != nil {
		log.WithError(err).Error("Error deleting note")
		return nil, err
	}
	delete(cn.storage, id)
	return toNotableList(cn.storage), nil
}

func (cn *UIService) Register(user *pb.User) error {
	log := log.WithFields(logrus.Fields{
		"method": "Register",
	})
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2000*time.Millisecond)
	defer cancelFunc()
	token, err := cn.uc.Register(ctx, user)
	if err != nil {
		return err
	}
	cn.jwt = token.Token
	cn.hash = getHash(user)
	log.Infof("registered new user: %s", user)
	return nil
}

func (cn *UIService) Login(user *pb.User) error {
	log := log.WithFields(logrus.Fields{
		"method": "Login",
	})
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2000*time.Millisecond)
	defer cancelFunc()
	token, err := cn.uc.Login(ctx, user)
	if err != nil {
		return err
	}
	cn.jwt = token.Token
	cn.hash = getHash(user)
	log.Infof("user sign in: %s", user.Username)
	return nil
}

func getHash(user *pb.User) []byte {
	bytes := sha256.Sum256([]byte(user.Email + user.Password))
	return bytes[:]
}

func (cn *UIService) addToken(ctx context.Context) context.Context {
	md := metadata.New(map[string]string{"token": cn.jwt})
	return metadata.NewOutgoingContext(ctx, md)

}

func toNotableList(notes map[uuid.UUID]*models.Noteable) *[]models.Noteable {
	var l []models.Noteable
	for _, u := range notes {
		l = append(l, *u)
	}
	return &l
}

func unmarshalNote(ctx context.Context, key []byte, noteDto *pb.Note) (models.Noteable, error) {

	decrypt, err := util.Decrypt(ctx, key, noteDto.SecretData)
	if err != nil {
		return nil, err
	}

	switch noteDto.Type {
	case models.CARD.String():
		note := &models.BankCardNote{}
		err = json.Unmarshal(decrypt, note)
		if err != nil {
			return nil, err
		}
		return note, nil

	case models.CREDENTIAL.String():
		note := &models.CredentialNote{}
		err = json.Unmarshal(decrypt, note)
		if err != nil {
			return nil, err
		}
		return note, nil

	case models.BINARY.String():
		note := &models.BinaryNote{}
		err = json.Unmarshal(decrypt, note)
		if err != nil {
			return nil, err
		}
		return note, nil

	case models.TEXT.String():
		note := &models.TextNote{}
		err = json.Unmarshal(decrypt, note)
		if err != nil {
			return nil, err
		}
		return note, nil

	default:
		return nil, errors.New("unknown note type")
	}

}
