package server

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/kamwawrzak/jwt-auth-service/internal/config"
)

type authenticator interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

type server struct {
	log *logrus.Logger
	port int
	mux *http.ServeMux
	handler authenticator
}

func NewServer(cfg config.ServerCfg, log *logrus.Logger, handler authenticator) *server {
	return &server{
		log: log,
		port: cfg.Port,
		mux: http.NewServeMux(),
		handler: handler,
	}
}

func (s *server) Start() error {
	s.registerEndpoints()
	s.log.WithField("port", s.port).Info("Starting http server")
	err := http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.mux)
	if err != nil {
		return err
	}
	return nil
}

func (s *server) registerEndpoints() {
	s.mux.HandleFunc("GET /ping", s.handler.Ping)
}
