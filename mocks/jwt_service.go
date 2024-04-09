package mocks

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type MockJwtService struct {
	mock.Mock
}

func (m *MockJwtService) CreateToken(id string) (string, *time.Time, error) {
	args := m.Called(id)
	jwt, ok := args.Get(0).(string)
	if !ok {
		return "", &time.Time{}, nil
	}
	expireAt, ok := args.Get(1).(time.Time)
	if !ok {
		return jwt, &time.Time{}, nil
	}

	return jwt, &expireAt, args.Error(2)
}
