package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func NewClaims(userID int64, duration time.Duration) Claims {
	return Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
}

func (c Claims) IsExpired() error {
	if time.Now().After(c.ExpiresAt.Time) {
		// FIXME: エラー定義
		return jwt.ErrTokenExpired
	}
	return nil
}

type Payload struct {
	Token     string
	ExpiredAt time.Time
}
