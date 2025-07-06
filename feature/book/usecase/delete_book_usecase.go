package usecase

import (
	"context"
	"errors"
	"readly/db/transaction"
	bookRepo "readly/feature/book/repository"
)

type DeleteBookUseCase interface {
	DeleteBook(ctx context.Context, req DeleteBookRequest) error
}

type DeleteBookUseCaseImpl struct {
	transactor         transaction.Transactor
	bookRepo           bookRepo.BookRepository
	readingHistoryRepo bookRepo.ReadingHistoryRepository
}

func NewDeleteBookUseCase(
	transactor transaction.Transactor,
	bookRepo bookRepo.BookRepository,
	readingHistoryRepo bookRepo.ReadingHistoryRepository,
) DeleteBookUseCase {
	return &DeleteBookUseCaseImpl{
		transactor:         transactor,
		bookRepo:           bookRepo,
		readingHistoryRepo: readingHistoryRepo,
	}
}

func (u *DeleteBookUseCaseImpl) DeleteBook(ctx context.Context, req DeleteBookRequest) error {
	err := u.transactor.Exec(ctx, func() error {
		deleteHistoryArgs := bookRepo.DeleteReadingHistoryRequest{
			UserID: req.UserID,
			BookID: req.BookID,
		}
		err := u.readingHistoryRepo.Delete(ctx, deleteHistoryArgs)
		if err != nil {
			if errors.Is(err, bookRepo.ErrNoRowsDeleted) {
				return newError(BadRequest, NotFoundBookError, "reading history not found")
			}
			return err
		}
		err = u.deleteBookGenres(ctx, req.BookID)
		if err != nil {
			if errors.Is(err, bookRepo.ErrNoRowsDeleted) {
				return newError(BadRequest, NotFoundBookError, "genre not found")
			}
			return err
		}
		err = u.bookRepo.DeleteBook(ctx, bookRepo.NewDeleteBookRequest(req.BookID))
		if err != nil {
			if errors.Is(err, bookRepo.ErrNoRowsDeleted) {
				return newError(BadRequest, NotFoundBookError, "book not found")
			}
			return err
		}
		return nil
	})
	return handle(err)
}

func (u *DeleteBookUseCaseImpl) deleteBookGenres(ctx context.Context, bookID int64) error {
	res, err := u.bookRepo.GetGenresByBookID(ctx, bookRepo.NewGetGenresByBookIDRequest(bookID))
	if err != nil {
		return err
	}
	for _, g := range res.Genres {
		args := bookRepo.DeleteBookGenreRequest{
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
