package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert" 
)

func TestCreateToken(t *testing.T) {
	// arrange
	userId := "123"
	secretKey := "test_secret"
	jwtSvc := JWTService{
		secretKey: secretKey, 
		ttl: 1 * time.Hour,
	}

	// act
	actual, expireAt, err := jwtSvc.CreateToken(userId)

	// assert
	assert.NoError(t, err)
	assert.True(t, expireAt.After(time.Now()))
	assert.NotEqual(t, "", actual)
}
