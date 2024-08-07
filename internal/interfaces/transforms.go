package interfaces

import (
	pb "github.com/bonus2k/go-diplom-gophkeeper/internal/interfaces/proto"
	"github.com/bonus2k/go-diplom-gophkeeper/internal/models"
	"github.com/google/uuid"
)

func DtoToEntity(note *pb.Note) (models.SecretData, error) {
	uid, err := uuid.Parse(note.Id)
	if err != nil {
		return models.SecretData{}, err
	}
	return models.SecretData{
		ID:     uid,
		Type:   note.Type,
		Name:   note.Name,
		Secret: note.SecretData,
	}, nil
}
