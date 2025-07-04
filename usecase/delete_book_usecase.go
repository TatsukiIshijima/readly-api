package usecase

import (
	"context"
	"errors"
	"readly/repository"
	userRepo "readly/user/repository"
)

type DeleteBookUseCase interface {
	DeleteBook(ctx context.Context, req DeleteBookRequest) error
}

type DeleteBookUseCaseImpl struct {
	transactor         repository.Transactor
	bookRepo           repository.BookRepository
	readingHistoryRepo repository.ReadingHistoryRepository
	userRepo           userRepo.UserRepository
}

func NewDeleteBookUseCase(
	transactor repository.Transactor,
	bookRepo repository.BookRepository,
	readingHistoryRepo repository.ReadingHistoryRepository,
	userRepo userRepo.UserRepository,
) DeleteBookUseCase {
	return &DeleteBookUseCaseImpl{
		transactor:         transactor,
		bookRepo:           bookRepo,
		readingHistoryRepo: readingHistoryRepo,
		userRepo:           userRepo,
	}
}

func (u *DeleteBookUseCaseImpl) DeleteBook(ctx context.Context, req DeleteBookRequest) error {
	err := u.transactor.Exec(ctx, func() error {
		deleteHistoryArgs := repository.DeleteReadingHistoryRequest{
			UserID: req.UserID,
			BookID: req.BookID,
		}
		err := u.readingHistoryRepo.Delete(ctx, deleteHistoryArgs)
		if err != nil {
			if errors.Is(err, repository.ErrNoRowsDeleted) {
				return newError(BadRequest, NotFoundBookError, "reading history not found")
			}
			return err
		}
		err = u.deleteBookGenres(ctx, req.BookID)
		if err != nil {
			if errors.Is(err, repository.ErrNoRowsDeleted) {
				return newError(BadRequest, NotFoundBookError, "genre not found")
			}
			return err
		}
		err = u.bookRepo.DeleteBook(ctx, repository.NewDeleteBookRequest(req.BookID))
		if err != nil {
			if errors.Is(err, repository.ErrNoRowsDeleted) {
				return newError(BadRequest, NotFoundBookError, "book not found")
			}
			return err
		}
		return nil
	})
	return handle(err)
}

func (u *DeleteBookUseCaseImpl) deleteBookGenres(ctx context.Context, bookID int64) error {
	res, err := u.bookRepo.GetGenresByBookID(ctx, repository.NewGetGenresByBookIDRequest(bookID))
	if err != nil {
		return err
	}
	for _, g := range res.Genres {
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
