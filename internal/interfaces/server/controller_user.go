package server

import (
	"context"
	"errors"

	pb "github.com/bonus2k/go-diplom-gophkeeper/internal/interfaces/proto"
	"github.com/bonus2k/go-diplom-gophkeeper/internal/models"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (s *Controller) Register(ctx context.Context, user *pb.User) (*pb.JwtToken, error) {
	log := log.WithFields(logrus.Fields{
		"method": "Register",
		"user":   user.Email,
	})

	uc := &models.UserCtx{
		Username: "not register",
		Email:    user.Email,
	}
	ctx = context.WithValue(ctx, "UserCtx", uc)
	newUser, err := s.db.AddUser(ctx, &models.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, status.Error(codes.AlreadyExists, "user already exists")
		}
		log.WithError(err).Error("could not add user")
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.JwtToken{Token: newUser.Username}, nil
}
func (s *Controller) Login(ctx context.Context, user *pb.User) (*pb.JwtToken, error) {
	log := log.WithFields(logrus.Fields{
		"method": "Register",
		"user":   user.Email,
	})

	uc := &models.UserCtx{
		Username: "not register",
		Email:    user.Email,
	}
	ctx = context.WithValue(ctx, "UserCtx", uc)
	getUser, err := s.db.GetUser(ctx, uc.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "User not found")
		}
		log.WithError(err).Error("Could not get user")
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.JwtToken{Token: getUser.Username}, nil
}
