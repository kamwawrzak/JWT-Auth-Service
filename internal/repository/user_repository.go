package repository

import (
	"context"
	"fmt"
	"database/sql"

	mysql "github.com/go-sql-driver/mysql"

	"github.com/kamwawrzak/jwt-auth-service/internal/model"
)

var duplicatedErr = "already exist"


type userRepository struct {}

func NewUserRepository()*userRepository{
	return &userRepository{}
}

func (ur *userRepository) GetUser(ctx context.Context, db *sql.DB, id string) (*model.User, error) {
	var user model.User
	query := "SELECT id, email, password_hash, created_at FROM users WHERE id = ?"
	err := db.QueryRow(query, id).Scan(
		&user.Id,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
    if err != nil {
        return nil, err
    }

	return &user, nil
}

func (ur *userRepository) CreateUser(ctx context.Context, dbConn *sql.DB, newUser *model.User) (*model.User, error) {
	query := "INSERT INTO `users` (`email`, `password_hash`) VALUES (?, ?)"
	result, err := dbConn.ExecContext(ctx, query, newUser.Email, newUser.Password)
	if err != nil {
		if isDuplicateEntryError(err) {
			return nil, fmt.Errorf(duplicatedErr)
		}
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user, err := ur.GetUser(ctx, dbConn, fmt.Sprint(id))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func isDuplicateEntryError(err error) bool {
    mysqlErr, ok := err.(*mysql.MySQLError)
    if !ok {
        return false
    }
    return mysqlErr.Number == 1062
}
