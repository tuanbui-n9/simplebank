package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token has expired")
)

type Payload struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	IssueAt  time.Time `json:"issueAt"`
	ExpireAt time.Time `json:"expireAt"`
	Role     string    `json:"role"`
}

func NewPayload(username string, role string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &Payload{
		ID:       tokenID,
		Username: username,
		IssueAt:  time.Now(),
		ExpireAt: time.Now().Add(duration),
		Role:     role,
	}, nil
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpireAt) {
		return ErrExpiredToken
	}

	return nil
}
