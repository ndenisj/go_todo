package token

import "time"

// Maker: an interface for managing tokens
type Maker interface {
	CreateToken(userID int64, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}
