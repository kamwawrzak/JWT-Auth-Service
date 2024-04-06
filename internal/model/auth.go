package model

import "time"

type LoginResponse struct {
	Jwt string `json:"jwt"`
	ExpireAt time.Time `json:"expire_at"`
}
