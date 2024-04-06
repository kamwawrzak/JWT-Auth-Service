package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/kamwawrzak/jwt-auth-service/internal/config"	
)

type JWTService struct {
	secretKey string
	ttl time.Duration
}

func NewJWTService(cfg config.JWTCfg) *JWTService{
	return &JWTService{
		secretKey: cfg.SecretKey,
		ttl: cfg.TimeToLive,
	}
}

func (j *JWTService) CreateToken(id string) (string, *time.Time, error) {
	expireAt := getExpirationTime(time.Now(), j.ttl)
	claims := j.createClaims(expireAt, id)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", nil, err
	}

	return signedToken, &expireAt, nil
}

func (j *JWTService) createClaims(expireAt time.Time, id string) (jwt.Claims) {
	return jwt.MapClaims{ 
		"sub": id, 
		"exp": time.Now().Add(j.ttl).Unix(), 
	}
}

func getExpirationTime(now time.Time, expireAfter time.Duration) time.Time {
	return now.Add(expireAfter)
}
