package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/kamwawrzak/jwt-auth-service/internal/model"
)

type userReader interface {
	GetUserByEmail(ctx context.Context, db *sql.DB, email string,) (*model.User, error)
}

type jwtCreator interface {
	CreateToken(id string) (string, *time.Time, error)
}

type LoginService struct {
	db *sql.DB
	userR userReader
	jwtSvc jwtCreator
}

func NewLoginService(dbConn *sql.DB, userR userReader, jwtSvc jwtCreator) *LoginService{
	return &LoginService{
		db: dbConn,
		userR: userR,
		jwtSvc: jwtSvc,
	}
}

func (ls *LoginService) Login(ctx context.Context, email, password string) (string, *time.Time, error) {
	user, err := ls.userR.GetUserByEmail(ctx, ls.db, email)
	if err != nil {
		return "", nil, err
	}

	err = verifyPassword(password, user.Password)
	if err != nil {
		return "", nil, err
	}

	return ls.jwtSvc.CreateToken(user.Id)
}
