package user

import (
	"context"
	"readly/domain"
)

type Repository interface {
	Register(ctx context.Context, req RegisterRequest) (domain.User, error)
	Login(ctx context.Context, req LoginRequest) (domain.User, error)
	GetByID(ctx context.Context, id int64) (domain.User, error)
	GetByEmail(ctx context.Context, email string) (domain.User, error)
	Update(ctx context.Context, req UpdateRequest) (domain.User, error)
	Delete(ctx context.Context, id int64) error
}

type RegisterRequest struct {
	Name     string
	Email    string
	Password string
}

type LoginRequest struct {
	Email    string
	Password string
}

type UpdateRequest struct {
	ID       int64
	Name     string
	Email    string
	Password string
}
