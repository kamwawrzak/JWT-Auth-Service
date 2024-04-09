package repository

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert" 
	"github.com/stretchr/testify/require"

	"github.com/kamwawrzak/jwt-auth-service/internal/model"
)

func TestGetUser(t *testing.T) {
	// arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	repo := NewUserRepository()
	expectedUser := prepareExampleUser()
	
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
	expectedUser := prepareExampleUser()
	
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

func TestCreateUserSuccess(t *testing.T){
	// arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	repo := NewUserRepository()
	expectedUser := prepareExampleUser()

	mock.ExpectExec("INSERT INTO `users`").
		WithArgs(expectedUser.Email,expectedUser.Password).
		WillReturnResult(sqlmock.NewResult(1, 1))
    mock.ExpectQuery("SELECT").WithArgs("1").
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password_hash", "created_at"}).
		AddRow(123, expectedUser.Email, expectedUser.Password, expectedUser.CreatedAt))

	// act
	actualUser, err := repo.CreateUser(context.Background(), db, expectedUser)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, actualUser)
}

func TestCreateUserAlreadyExistsError(t *testing.T){
	// arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	repo := NewUserRepository()
	user := prepareExampleUser()

	mock.ExpectExec("INSERT INTO `users`").
		WithArgs(user.Email, user.Password).
		WillReturnError(&mysql.MySQLError{Number: 1062, Message: "Duplicate entry 'test@example.com' for key 'email'"})

	// act
	_, err = repo.CreateUser(context.Background(), db, user)

	// assert
	assert.Equal(t, duplicatedErr, err.Error())
}

func prepareExampleUser() *model.User {
	return &model.User{
		Id: "123",
		Email: "test@gmail.com",
		Password: "hash-pass-123",
	}
}
