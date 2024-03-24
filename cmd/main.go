package main

import (
	"github.com/sirupsen/logrus"

	"github.com/kamwawrzak/jwt-auth-service/internal/config"
	"github.com/kamwawrzak/jwt-auth-service/internal/server"
)

var defaultConfigPath = "./config.local.yaml"


func main() {
	log := logrus.New()
	cfg, err := config.NewConfig(defaultConfigPath)
	if err != nil {
		log.WithError(err).Fatal("Failed to read configuration file: ", defaultConfigPath)
	}
	handler := server.NewAuthHandler()
	s := server.NewServer(cfg.ServerCfg, log, handler)

	err = s.Start()
	if err != nil {
		log.WithError(err).Fatal("Failed to start server")
	}
}
