package usecase

import (
	"context"
	"readly/entity"
)

type Executor interface {
	SignUp(ctx context.Context, req SignUpRequest) (*entity.User, error)
	RegisterBook(ctx context.Context, req RegisterBookRequest) (*entity.Book, error)
	DeleteBook(ctx context.Context, req DeleteBookRequest) error
}
