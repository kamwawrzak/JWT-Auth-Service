package main

import (
	"github.com/sirupsen/logrus"

	"github.com/kamwawrzak/jwt-auth-service/internal/config"
	"github.com/kamwawrzak/jwt-auth-service/internal/app"
)

var defaultConfigPath = "./config.local.yaml"


func main() {
	log := logrus.New()

	cfg, err := config.NewConfig(defaultConfigPath)
	if err != nil {
		log.WithError(err).Fatal("Failed to read configuration file: ", defaultConfigPath)
	}

	db, err := app.DbConn(&cfg.DbCfg)
	if err != nil {
		log.WithError(err).Fatal("Failed connect to DB")
	}
	defer db.Close()

	s := app.Setup(cfg, log, db)
	err = s.Start()
	if err != nil {
		log.WithError(err).Fatal("Failed to start server")
	}
}
