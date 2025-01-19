package usecase

import (
	"context"
	"readly/entity"
)

type Executor interface {
	RegisterBook(ctx context.Context, req RegisterBookRequest) (*entity.Book, error)
	DeleteBook(ctx context.Context, req DeleteBookRequest) error
}

type DeleteBookRequest struct {
}
