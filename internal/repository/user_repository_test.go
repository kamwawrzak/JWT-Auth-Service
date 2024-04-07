package repository

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert" 
	"github.com/stretchr/testify/require"

	"github.com/kamwawrzak/jwt-auth-service/internal/model"
)

func TestGetUser(t *testing.T) {
	// arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	repo := NewUserRepository()
	expectedUser := &model.User{
		Id: "123",
		Email: "test@gmail.com",
		Password: "hash-pass-123",
	}
	
	query := "SELECT id, email, password_hash, created_at FROM users WHERE id = ?"

	rows := sqlmock.NewRows([]string{"id", "email", "password_hash", "created_at"}).
		AddRow(expectedUser.Id, expectedUser.Email, expectedUser.Password, expectedUser.CreatedAt)

	mock.ExpectQuery(query).WithArgs(expectedUser.Id).WillReturnRows(rows)

	// act
	actualUser, err := repo.GetUser(context.Background(), db, expectedUser.Id)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, actualUser)
}

func TestGetUserByEmail(t *testing.T) {
	// arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	repo := NewUserRepository()
	expectedUser := &model.User{
		Id: "123",
		Email: "test@gmail.com",
		Password: "hash-pass-123",
	}
	
	query := "SELECT id, email, password_hash, created_at FROM users WHERE email = ?"

	rows := sqlmock.NewRows([]string{"id", "email", "password_hash", "created_at"}).
		AddRow(expectedUser.Id, expectedUser.Email, expectedUser.Password, expectedUser.CreatedAt)

	mock.ExpectQuery(query).WithArgs(expectedUser.Email).WillReturnRows(rows)

	// act
	actualUser, err := repo.GetUserByEmail(context.Background(), db, expectedUser.Email)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, actualUser)
}

func TestCreateUser(t *testing.T){
	// arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	repo := NewUserRepository()
	expectedUser := &model.User{
		Id: "1",
		Email: "test@gmail.com",
		Password: "hash-pass-123",
	}

	mock.ExpectExec("INSERT INTO `users`").WithArgs(
		expectedUser.Email,
		expectedUser.Password).
		WillReturnResult(sqlmock.NewResult(1, 1))
    mock.ExpectQuery("SELECT").WithArgs("1").
		WillReturnRows(sqlmock.
			NewRows([]string{"id", "email", "password_hash", "created_at"}).
				AddRow(1, expectedUser.Email, expectedUser.Password, expectedUser.CreatedAt))

	// act
	actualUser, err := repo.CreateUser(context.Background(), db, expectedUser)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, actualUser)
}