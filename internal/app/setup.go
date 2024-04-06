package app

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"

	"github.com/kamwawrzak/jwt-auth-service/internal/config"
	"github.com/kamwawrzak/jwt-auth-service/internal/repository"
	"github.com/kamwawrzak/jwt-auth-service/internal/service"
	"github.com/kamwawrzak/jwt-auth-service/internal/server"
)

func Setup(cfg *config.Config, log *logrus.Logger, db *sql.DB) *server.Server {
	userR := repository.NewUserRepository()
	jwtSvc := service.NewJWTService(cfg.JWTCfg)
	signupSvc := service.NewSignupService(db, userR)
	loginSvc := service.NewLoginService(db, userR, jwtSvc)
	h := server.NewAuthHandler(signupSvc, loginSvc)
	return server.NewServer(cfg.ServerCfg, log, h)
}


func DbConn(cfg *config.DbCfg) (*sql.DB, error) {
	writerCfg := cfg.Writer
	connsCfg := cfg.Connections
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
	 	writerCfg.User,
	 	writerCfg.Password,
	 	writerCfg.Host,
	 	writerCfg.Port,
	 	writerCfg.Database,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * connsCfg.MaxConnLifetime)
	db.SetMaxOpenConns(connsCfg.MaxOpenConns)
	db.SetMaxIdleConns(connsCfg.MaxIdleConns)

	if err := db.Ping(); err != nil {
        return nil, err
    }

	return db, nil
}