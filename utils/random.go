package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "qwertyuiopasdfghjklzxcvbnm"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt: generate random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString: generate random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner: generate a random owner
func RandomOwner() string {
	return RandomString(6)
}

// RandomTitle: generate a random title
func RandomTitle() string {
	return RandomString(9)
}

// RandomContent: generate a random content
func RandomContent() string {
	return RandomString(30)
}

func RandomFullname() string {
	return RandomString(5) + " " + RandomString(6)
}

func RandomPhone() string {
	return RandomString(11)
}

func RandomUsername() string {
	return RandomString(6)
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
