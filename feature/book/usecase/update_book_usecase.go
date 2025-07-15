package usecase

import (
	"context"
	"errors"
	"readly/db/transaction"
	bookRepo "readly/feature/book/repository"
	"readly/util"
)

type UpdateBookUseCase interface {
	UpdateBook(ctx context.Context, request UpdateBookRequest) (*UpdateBookResponse, error)
}

type UpdateBookUseCaseImpl struct {
	transactor         transaction.Transactor
	bookRepository     bookRepo.BookRepository
	readingHistoryRepo bookRepo.ReadingHistoryRepository
}

func NewUpdateBookUseCase(
	transactor transaction.Transactor,
	bookRepository bookRepo.BookRepository,
	readingHistoryRepo bookRepo.ReadingHistoryRepository,
) UpdateBookUseCase {
	return &UpdateBookUseCaseImpl{
		transactor:         transactor,
		bookRepository:     bookRepository,
		readingHistoryRepo: readingHistoryRepo,
	}
}

func (u UpdateBookUseCaseImpl) UpdateBook(ctx context.Context, request UpdateBookRequest) (*UpdateBookResponse, error) {
	var res *UpdateBookResponse
	err := u.transactor.Exec(ctx, func() error {
		err := u.updateBook(ctx, request)
		if err != nil {
			return err
		}
		err = u.updateGenresIfNeed(ctx, request)
		if err != nil {
			return err
		}
		err = u.updateReadingHistory(ctx, request)
		if err != nil {
			return err
		}
		res = &UpdateBookResponse{
			BookID: request.BookID,
		}
		return nil
	})
	return res, handle(err)
}

func (u UpdateBookUseCaseImpl) updateBook(ctx context.Context, req UpdateBookRequest) error {
	if err := req.Validate(); err != nil {
		return handle(err)
	}
	updateBookReq := req.toBookRepoRequest()
	_, err := u.bookRepository.UpdateBook(ctx, updateBookReq)
	if err != nil {
		if errors.Is(err, bookRepo.ErrNoRows) {
			return newError(BadRequest, InvalidRequestError, "book not found")
		}
		return err
	}
	return nil
}

func (u UpdateBookUseCaseImpl) updateGenresIfNeed(ctx context.Context, req UpdateBookRequest) error {
	getBookGenresReq := bookRepo.GetGenresByBookIDRequest{
		ID: req.BookID,
	}
	bgRes, err := u.bookRepository.GetGenresByBookID(ctx, getBookGenresReq)
	if err != nil {
		return err
	}
	if len(req.Genres) == 0 && len(bgRes.Genres) == 0 {
		return nil
	}
	if util.EqualSet(bgRes.Genres, req.Genres) {
		return nil
	}
	// 差分がある場合は一度Bookに紐づくgenreを削除してから追加する
	for _, genre := range bgRes.Genres {
		err := u.bookRepository.DeleteBookGenre(ctx, bookRepo.DeleteBookGenreRequest{
			BookID:    req.BookID,
			GenreName: genre,
		})
		if err != nil {
			return err
		}
	}
	for _, genre := range req.Genres {
		_, err := u.bookRepository.CreateBookGenre(ctx, bookRepo.CreateBookGenreRequest{
			BookID:    req.BookID,
			GenreName: genre,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (u UpdateBookUseCaseImpl) updateReadingHistory(ctx context.Context, req UpdateBookRequest) error {
	updateReadingHistoryReq := req.toReadingHistoryRepoRequest()
	_, err := u.readingHistoryRepo.Update(ctx, updateReadingHistoryReq)
	if err != nil {
		if errors.Is(err, bookRepo.ErrNoRows) {
			return newError(BadRequest, InvalidRequestError, "user not found")
		}
		return err
	}
	return nil
}
