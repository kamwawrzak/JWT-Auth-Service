package service

import (
	"context"
	"database/sql"

	"github.com/kamwawrzak/jwt-auth-service/internal/model"
)

type userCreator interface {
	CreateUser(ctx context.Context, db *sql.DB, u *model.User) (*model.User, error)
}

type SignupService struct {
	db *sql.DB
	userR userCreator
}

func NewSignupService(dbConn *sql.DB, userR userCreator) *SignupService{
	return &SignupService{
		db: dbConn,
		userR: userR,
	}
}

func (s *SignupService) CreateUser(ctx context.Context, u *model.User) (*model.User, error){
	// hash password
	u.Password = "hashed"
	user, err := s.userR.CreateUser(ctx, s.db, u)
	if err != nil {
		return nil, err
	}

	return user, nil
}
