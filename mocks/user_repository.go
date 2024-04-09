package mocks

import (
	"context"
	"database/sql"

	"github.com/stretchr/testify/mock"

	"github.com/kamwawrzak/jwt-auth-service/internal/model"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) GetUserByEmail(ctx context.Context, db *sql.DB, email string) (*model.User, error){
	args := m.Called(ctx, db, email)
	user, ok := args.Get(0).(*model.User)
	if !ok {
		return &model.User{}, args.Error(1)
	}
    return user, args.Error(1)
}

func (m *MockUserRepo) CreateUser(ctx context.Context, db *sql.DB, user *model.User) (*model.User, error) {
    args := m.Called(ctx, db, user)
	user, ok := args.Get(0).(*model.User)
	if !ok {
		return &model.User{}, args.Error(1)
	}
    return user, args.Error(1)
}
