package auth

import "time"

type Maker interface {
	Generate(email string, duration time.Duration) (string, error)
	Verify(token string) (*Payload, error)
}
