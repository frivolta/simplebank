package token

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

// Contain payload data of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json"issued_at"`
	ExpiresAt time.Time `json:"expire_at"`
}

var (
	ErrExpiredToken   = errors.New("token has expired")
	ErrorInvalidToken = errors.New("token is invalid")
)

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}
	return payload, nil
}

// checks if the token payload is valid or not
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiresAt) {
		return ErrExpiredToken
	}
	return nil
}
