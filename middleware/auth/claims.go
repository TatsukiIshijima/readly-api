package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func NewClaims(userID int64, duration time.Duration) (Claims, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return Claims{}, err
	}
	return Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        id.String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}, nil
}

func (c Claims) IsExpired() error {
	if time.Now().After(c.ExpiresAt.Time) {
		// FIXME: エラー定義
		return jwt.ErrTokenExpired
	}
	return nil
}

type Payload struct {
	ID        uuid.UUID
	Token     string
	ExpiredAt time.Time
}
