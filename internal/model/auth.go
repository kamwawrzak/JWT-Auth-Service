package model

import "time"

type LoginInput struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Jwt string `json:"jwt"`
	ExpireAt time.Time `json:"expire_at"`
}

type SignupInput struct {
	Email string `json:"email"`
	Password string `json:"password"`
	PasswordRepeat string `json:"password_repeat"`
}

type SignupResponse struct {
	Email string `json:"email"`
}
