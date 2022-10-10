package token

import "time"

type Maker interface {
	// Create and sign a new token
	CreateToken(username string, duration time.Duration) (string, *Payload, error)
	// Check if the input is valid or not
	VerifyToken(token string) (*Payload, error)
}
