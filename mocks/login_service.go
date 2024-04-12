package mocks

import (
	"context"
	"time"
	
	"github.com/stretchr/testify/mock"
)

type MockLoginService struct {
	mock.Mock
}

func (m *MockLoginService) Login(ctx context.Context, email, password string) (string, *time.Time, error) {
	args := m.Called(ctx, email, password)
	jwt, ok := args.Get(0).(string)
	if !ok {
		return "", &time.Time{}, args.Error(2)
	}
	expireAt, ok := args.Get(1).(time.Time)
	if !ok {
		return jwt, &time.Time{}, args.Error(2)
	}
	return jwt, &expireAt, args.Error(2)
}
