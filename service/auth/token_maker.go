package auth

import "time"

type TokenMaker interface {
	Generate(userID int64, duration time.Duration) (*Payload, error)
	Verify(token string) (*Claims, error)
}
