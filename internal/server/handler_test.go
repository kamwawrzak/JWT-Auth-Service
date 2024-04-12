package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/kamwawrzak/jwt-auth-service/internal/model"
	"github.com/kamwawrzak/jwt-auth-service/mocks"
)

func TestCreateUser(t *testing.T){
	// arrange
	handler, signMock, _ := getHandler()
	newUser := prepareExampleUser()
	input := &model.SignupInput{
		Email: newUser.Email,
		Password: newUser.Password,
		PasswordRepeat: newUser.Password,
	}

	signMock.On("CreateUser", mock.Anything, mock.Anything).Return(newUser, nil)

	jsonBody, err := json.Marshal(input)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonBody))
	if err != nil {
		require.NoError(t, err)
	}

	rr := httptest.NewRecorder()
	handlerWrapper := http.HandlerFunc(handler.SignUp)

	expected := &model.SignupResponse{
		Email: newUser.Email,
	}

	// act
	handlerWrapper.ServeHTTP(rr, req)

	var actual *model.SignupResponse
	err = json.NewDecoder(rr.Body).Decode(&actual)
	require.NoError(t, err)

	// assert
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, expected, actual)
	
}

func TestLoginSuccess(t *testing.T){
	// arrange
	handler, _, loginMock := getHandler()
	user := prepareExampleUser()
	input := &model.LoginInput{
		Email: user.Email,
		Password: user.Password,
	}
	expectedResult := &model.LoginResponse{
		Jwt: "aaa.bbb.ccc",
		ExpireAt: time.Now(),
	}

	loginMock.
	On("Login", mock.Anything, mock.Anything, mock.Anything).
	Return(expectedResult.Jwt, expectedResult.ExpireAt, nil)

	jsonBody, err := json.Marshal(input)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
	if err != nil {
		require.NoError(t, err)
	}

	rr := httptest.NewRecorder()
	handlerWrapper := http.HandlerFunc(handler.Login)

	// act
	handlerWrapper.ServeHTTP(rr, req)

	var actual *model.LoginResponse
	err = json.NewDecoder(rr.Body).Decode(&actual)
	require.NoError(t, err)

	// assert
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, expectedResult.Jwt, actual.Jwt)
	assert.True(t, actual.ExpireAt.Before(time.Now().Add(time.Hour)))
}

func getHandler() (*AuthHandler, *mocks.MockSignupService, *mocks.MockLoginService) {
	mockLoginSvc := &mocks.MockLoginService{}
	mockSignupSvc := &mocks.MockSignupService{}
	handler := NewAuthHandler(mockSignupSvc, mockLoginSvc)

	return handler, mockSignupSvc, mockLoginSvc
}

func prepareExampleUser() *model.User {
	return &model.User{
		Id: "123",
		Email: "test@example.com",
		Password: "hash-pass-123",
	}
}
