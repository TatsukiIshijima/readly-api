package usecase

import (
	"context"
	"errors"
	"github.com/lib/pq"
	"readly/db/transaction"
	bookRepo "readly/feature/book/repository"
)

type CreateGenresUseCase interface {
	CreateGenres(ctx context.Context, request CreateGenresRequest) error
}

type CreateGenresUseCaseImpl struct {
	transactor transaction.Transactor
	bookRepo   bookRepo.BookRepository
}

func NewCreateGenresUseCase(
	transactor transaction.Transactor,
	bookRepo bookRepo.BookRepository,
) CreateGenresUseCase {
	return &CreateGenresUseCaseImpl{
		transactor: transactor,
		bookRepo:   bookRepo,
	}
}

func (u *CreateGenresUseCaseImpl) CreateGenres(ctx context.Context, request CreateGenresRequest) error {
	err := u.transactor.Exec(ctx, func() error {
		for _, name := range request.Names {
			_, err := u.bookRepo.CreateGenre(ctx, bookRepo.NewCreateGenreRequest(name))
			if err != nil {
				// Check if it's a duplicate key error (genre already exists)
				// If so, skip this genre and continue with the next one
				if u.checkDuplicateKeyError(err) == nil {
					continue
				}
				return err
			}
		}
		return nil
	})
	return handle(err)
}

// checkDuplicateKeyError checks if the error is a duplicate key error (PostgreSQL error code 23505)
// If it is, returns nil (ignoring the error), otherwise returns the original error
func (u *CreateGenresUseCaseImpl) checkDuplicateKeyError(err error) error {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) && pqErr.Code != "23505" {
		return err
	}
	return nil
}
