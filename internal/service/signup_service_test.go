package service

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/kamwawrzak/jwt-auth-service/internal/model"
	"github.com/kamwawrzak/jwt-auth-service/mocks"
)

func TestCreateUserSuccess(t *testing.T) {
	// arrange
	repoMock := &mocks.MockUserRepo{}
	dbMock := &sql.DB{}
	ctx := context.Background()

	signupSvc := NewSignupService(dbMock, repoMock)
	expectedUser := &model.User{
		Id: "123",
		Email: "test@example.com",
		Password : "hash123",
	}

	repoMock.On("CreateUser", ctx, dbMock, mock.Anything).Return(expectedUser, nil)

	// act
	actual, err := signupSvc.CreateUser(ctx, expectedUser)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, actual)
	repoMock.AssertExpectations(t)
}
