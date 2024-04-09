package service

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/kamwawrzak/jwt-auth-service/internal/model"
	"github.com/kamwawrzak/jwt-auth-service/mocks"
)

func TestLoginSuccess(t *testing.T) {
	// arrange
	repoMock := &mocks.MockUserRepo{}
	jwtMock := &mocks.MockJwtService{}
	dbMock := &sql.DB{}
	ctx := context.Background()
	expectedJWT := "aaa.bbb.ccc"
	
	loginSvc := NewLoginService(dbMock, repoMock, jwtMock)
	expectedUser := &model.User{
		Id: "1",
		Email: "test@example.com",
		Password : "$2a$10$HLaOi/54zWvNZWzODldKceh8OGerdIvLBLfzf48Z6kJd3v0fcgH1u",
	}

	repoMock.On("GetUserByEmail", ctx, dbMock, mock.Anything).Return(expectedUser, nil)
	jwtMock.On("CreateToken", mock.Anything).Return(expectedJWT, time.Now().Add(time.Hour), nil)

	// act
	actualJWT, expiredAt, err := loginSvc.Login(ctx, expectedUser.Email, "123")

	// assert
	assert.NoError(t, err)
	assert.Equal(t, expectedJWT, actualJWT)
	assert.True(t, expiredAt.After(time.Now()))
	repoMock.AssertExpectations(t)
	jwtMock.AssertExpectations(t)
}
