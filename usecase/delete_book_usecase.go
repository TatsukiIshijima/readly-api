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
		deleteHistoryArgs := repository.DeleteReadingHistoryRequest{
			UserID: req.UserID,
			BookID: req.BookID,
		}
		err := u.readingHistoryRepo.Delete(ctx, deleteHistoryArgs)
		if err != nil {
			return err
		}
		err = u.deleteBookGenres(ctx, req.BookID)
		if err != nil {
			return err
		}
		err = u.bookRepo.DeleteBook(ctx, req.BookID)
		if err != nil {
			return err
		}
		return nil
	})
	return handle(err)
}

func (u DeleteBookUseCase) deleteBookGenres(ctx context.Context, bookID int64) error {
	genres, err := u.bookRepo.GetGenresByBookID(ctx, bookID)
	if err != nil {
		return err
	}
	for _, g := range genres {
		args := repository.DeleteBookGenreRequest{
			BookID:    bookID,
			GenreName: g,
		}
		err := u.bookRepo.DeleteBookGenre(ctx, args)
		if err != nil {
			return err
		}
	}
	return nil
}
