package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomInt(t *testing.T) {
	randomInt := RandomInt(1, 3)

	require.NotEmpty(t, randomInt)
	require.Positive(t, randomInt)
}

func TestRandomString(t *testing.T) {
	randomStr := RandomString(int(RandomInt(1, 6)))

	require.NotEmpty(t, randomStr)
}
