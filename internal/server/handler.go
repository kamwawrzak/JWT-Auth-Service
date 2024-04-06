package server

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/kamwawrzak/jwt-auth-service/internal/model"
)

type userCreator interface {
	CreateUser(ctx context.Context, u *model.User) (*model.User, error)
}

type userVerifier interface {
	Login(ctx context.Context, email, password string) (string, *time.Time, error)
}

type AuthHandler struct {
	SignupSvc userCreator
	LoginSvc userVerifier
}


func NewAuthHandler(signupSvc userCreator, loginSvc userVerifier) *AuthHandler{
	return &AuthHandler{
		SignupSvc: signupSvc,
		LoginSvc: loginSvc,
	}
}

func (a *AuthHandler) Ping(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
}

func (a *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()

	var user model.User
	
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = a.SignupSvc.CreateUser(ctx, &user)
	if err != nil {
		switch err.Error() {
		case "already exist":
			http.Error(w, "User already exists", http.StatusConflict)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (a *AuthHandler) Login(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()

	var user model.User
	
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	jwt, expireAt, err := a.LoginSvc.Login(ctx, user.Email, user.Password)
	if err != nil {
		http.Error(w, "Authorization failed", http.StatusUnauthorized)
		return
	}

	payload := model.LoginResponse{
		Jwt: jwt,
		ExpireAt: *expireAt,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(payload)
}
