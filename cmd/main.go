package main

import (
	"github.com/sirupsen/logrus"

	"github.com/kamwawrzak/jwt-auth-service/internal/server"
)

var port = 8098

func main() {
	log := logrus.New()
	handler := server.NewAuthHandler()
	s := server.NewServer(port, handler)
	
	err := s.Start()
	if err != nil {
		log.WithError(err).Fatal("Failed to start server")
	}
}
