package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(t *testing.T, password string) (string, error) {
	return HashedPassword(password)
}

func TestHashedPassword(t *testing.T) {
	password, err := hashPassword(t, RandomString(6))

	require.NoError(t, err)
	require.NotEmpty(t, password)
}

func TestCheckPassword(t *testing.T) {
	password := RandomString(5)

	hashedPwd, err := hashPassword(t, password)
	require.NoError(t, err)

	err = CheckPassword(password, hashedPwd)
	require.NoError(t, err)
}

func TestPassword(t *testing.T) {
	password := RandomString(6)

	hashedPsssword1, err := HashedPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPsssword1)

	err = CheckPassword(password, hashedPsssword1)
	require.NoError(t, err)

	wrongPassword := RandomString(6)
	err = CheckPassword(wrongPassword, hashedPsssword1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedPsssword2, err := HashedPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPsssword2)
	require.NotEqual(t, hashedPsssword1, hashedPsssword2)
}
