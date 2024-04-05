package service

import (
	"context"
	"database/sql"

	"github.com/kamwawrzak/jwt-auth-service/internal/model"
)

type userReader interface {
	GetUserByEmail(ctx context.Context, db *sql.DB, email string,) (*model.User, error)
}

type LoginService struct {
	db *sql.DB
	userR userReader
}

func NewLoginService(dbConn *sql.DB, userR userReader) *LoginService{
	return &LoginService{
		db: dbConn,
		userR: userR,
	}
}

func (ls *LoginService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := ls.userR.GetUserByEmail(ctx, ls.db, email)
	if err != nil {
		return "", err
	}

	err = verifyPassword(password, user.Password)
	if err != nil {
		return "", err
	}

	return "success authentication", nil
}