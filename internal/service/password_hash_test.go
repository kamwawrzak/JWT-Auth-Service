package service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashPassword(t *testing.T) {
	// arrange
	password := "password-123"
	expectedHashLen := 60

	// act
	actual, err := hashPassword(password)

	// assert
	require.NoError(t, err)
	require.Equal(t, expectedHashLen, len(actual))

}

func TestVerifyPassword(t *testing.T) {
	// arrange
	password := "password-123"
	hashedPassword, err := hashPassword(password)
	require.NoError(t, err)

	// act
	err = verifyPassword(password, hashedPassword)

	// assert
	require.NoError(t, err)
}