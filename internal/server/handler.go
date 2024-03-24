package server

import "net/http"

type authHandler struct {}

func NewAuthHandler() *authHandler{
	return &authHandler{}
}

func (a *authHandler) Ping(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
}
