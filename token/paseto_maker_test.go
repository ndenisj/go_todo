package token

import (
	"testing"
	"time"

	"github.com/ndenisj/go_todo/utils"
	"github.com/stretchr/testify/require"
)

func TestPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(utils.RandomString(32))
	require.NoError(t, err)

	user_id := utils.RandomInt(1, 10)
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, payload, err := maker.CreateToken(user_id, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, user_id, payload.UserId)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(utils.RandomString(32))
	require.NoError(t, err)

	token, payload, err := maker.CreateToken(utils.RandomInt(1, 10), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}
