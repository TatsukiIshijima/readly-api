package usecase

import (
	"context"
	"readly/repository"
)

type DeleteBookUseCase struct {
	Executor
	transactor         repository.Transactor
	bookRepo           repository.BookRepository
	readingHistoryRepo repository.ReadingHistoryRepository
	userRepo           repository.UserRepository
}

func NewDeleteBookUseCase(transactor repository.Transactor, bookRepo repository.BookRepository, readingHistoryRepo repository.ReadingHistoryRepository, userRepo repository.UserRepository) DeleteBookUseCase {
	return DeleteBookUseCase{
		transactor:         transactor,
		bookRepo:           bookRepo,
		readingHistoryRepo: readingHistoryRepo,
		userRepo:           userRepo,
	}
}

type DeleteBookRequest struct {
	UserID int64
	BookID int64
}

func (u DeleteBookUseCase) DeleteBook(ctx context.Context, req DeleteBookRequest) error {
	err := u.transactor.Exec(ctx, func() error {
		err := u.bookRepo.DeleteBook(ctx, req.BookID)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
