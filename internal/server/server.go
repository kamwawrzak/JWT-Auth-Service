package server

import (
	"fmt"
	"net/http"
)

type authenticator interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

type server struct {
	port int
	mux *http.ServeMux
	handler authenticator
}

func NewServer(port int, handler authenticator) *server {
	return &server{
		port: port,
		mux: http.NewServeMux(),
		handler: handler,
	}
}

func (s *server) Start() error {
	s.registerEndpoints()
	err := http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.mux)
	if err != nil {
		return err
	}
	return nil
}

func (s *server) registerEndpoints() {
	s.mux.HandleFunc("GET /ping", s.handler.Ping)
}
