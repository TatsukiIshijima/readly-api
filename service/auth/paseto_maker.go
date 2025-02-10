package auth

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
	"time"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (*PasetoMaker, error) {
	if len(symmetricKey) < chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid secret key size")
	}
	return &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}, nil
}

func (p *PasetoMaker) Generate(userID int64, duration time.Duration) (*Payload, error) {
	claims, err := NewClaims(userID, duration)
	if err != nil {
		return nil, err
	}
	token, err := p.paseto.Encrypt(p.symmetricKey, claims, nil)
	if err != nil {
		return nil, err
	}
	id, err := uuid.Parse(claims.ID)
	if err != nil {
		return nil, err
	}
	return &Payload{
		ID:        id,
		Token:     token,
		ExpiredAt: claims.ExpiresAt.Time,
	}, nil
}

func (p *PasetoMaker) Verify(token string) (*Claims, error) {
	claims := &Claims{}
	err := p.paseto.Decrypt(token, p.symmetricKey, claims, nil)
	if err != nil {
		return nil, err
	}
	err = claims.IsExpired()
	if err != nil {
		return nil, err
	}
	return claims, nil
}
