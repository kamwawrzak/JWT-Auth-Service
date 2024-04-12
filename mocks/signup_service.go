package mocks

import (
	"context"
	
	"github.com/stretchr/testify/mock"

	"github.com/kamwawrzak/jwt-auth-service/internal/model"
)

type MockSignupService struct {
	mock.Mock
}

func (m *MockSignupService) CreateUser(ctx context.Context, input *model.SignupInput) (*model.User, error){
	args := m.Called(ctx, input)
	user, ok := args.Get(0).(*model.User)
	if !ok {
		return &model.User{}, args.Error(1)
	}
	return user, args.Error(1)
}
