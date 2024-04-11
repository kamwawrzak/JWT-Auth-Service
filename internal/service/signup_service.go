package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/kamwawrzak/jwt-auth-service/internal/model"
)

var errNotMatchedPass = errors.New("passwords don't match")

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

func (s *SignupService) CreateUser(ctx context.Context, input *model.SignupInput) (*model.User, error){
	if input.Password != input.PasswordRepeat {
		return nil, errNotMatchedPass
	}

	hashedPassword, err := hashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	newUser := &model.User{
		Email: input.Email,
		Password: hashedPassword,
	}

	user, err := s.userR.CreateUser(ctx, s.db, newUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}
