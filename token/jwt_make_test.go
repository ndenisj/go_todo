package token

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/ndenisj/go_todo/utils"
	"github.com/stretchr/testify/require"
)

func TestJwtToken(t *testing.T) {
	maker, err := NewJwtMaker(utils.RandomString(32))
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

func TestExpiredJWTToken(t *testing.T) {
	maker, err := NewJwtMaker(utils.RandomString(32))
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

func TestInvalidJWTTokenAlgNone(t *testing.T) {
	payload, err := NewPayload(utils.RandomInt(1, 10), time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := NewJwtMaker(utils.RandomString(32))
	require.NoError(t, err)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)
}
