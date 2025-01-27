package auth

import "time"

type Maker interface {
	Generate(userID int64, duration time.Duration) (string, error)
	Verify(token string) (*Claims, error)
}
