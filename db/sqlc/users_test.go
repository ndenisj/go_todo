package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/ndenisj/go_todo/utils"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := utils.HashedPassword(utils.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		FullName: utils.RandomFullname(),
		Phone: sql.NullString{
			String: utils.RandomPhone(),
			Valid:  true,
		},
		Username:       utils.RandomUsername(),
		Email:          utils.RandomEmail(),
		HashedPassword: hashedPassword,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, user.FullName, arg.FullName)
	require.Equal(t, user.Username, arg.Username)
	require.Equal(t, user.Email, arg.Email)
	require.Equal(t, user.HashedPassword, arg.HashedPassword)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.False(t, user.IsAdmin)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	// create todo
	user1 := createRandomUser(t)
	// get the todo
	user2, err := testQueries.GetUser(context.Background(), user1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)

	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}
