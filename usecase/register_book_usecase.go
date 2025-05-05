package usecase

import (
	"context"
	"errors"
	"github.com/lib/pq"
	"readly/entity"
	"readly/repository"
)

type RegisterBookUseCase interface {
	RegisterBook(ctx context.Context, req RegisterBookRequest) (*RegisterBookResponse, error)
}

type RegisterBookUseCaseImpl struct {
	transactor         repository.Transactor
	bookRepo           repository.BookRepository
	readingHistoryRepo repository.ReadingHistoryRepository
	userRepo           repository.UserRepository
}

func NewRegisterBookUseCase(
	transactor repository.Transactor,
	bookRepo repository.BookRepository,
	readingHistoryRepo repository.ReadingHistoryRepository,
	userRepo repository.UserRepository,
) RegisterBookUseCase {
	return &RegisterBookUseCaseImpl{
		transactor:         transactor,
		bookRepo:           bookRepo,
		readingHistoryRepo: readingHistoryRepo,
		userRepo:           userRepo,
	}
}

func (u *RegisterBookUseCaseImpl) RegisterBook(ctx context.Context, req RegisterBookRequest) (*RegisterBookResponse, error) {
	var book *entity.Book
	err := u.transactor.Exec(ctx, func() error {
		err := u.createAuthorIfNeed(ctx, req.AuthorName)
		if err != nil {
			return err
		}

		err = u.createPublisherIfNeed(ctx, req.PublisherName)
		if err != nil {
			return err
		}

		createArgs := repository.CreateBookRequest{
			Title:         req.Title,
			Description:   req.Description,
			CoverImageURL: req.CoverImageURL,
			URL:           req.URL,
			Author:        req.AuthorName,
			Publisher:     req.PublisherName,
			PublishDate:   req.PublishDate,
			ISBN:          req.ISBN,
		}
		b, err := u.bookRepo.CreateBook(ctx, createArgs)
		if err != nil {
			return err
		}
		for _, genre := range req.Genres {
			err := u.createGenreIfNeed(ctx, genre)
			if err != nil {
				return err
			}
			args := repository.CreateBookGenreRequest{
				BookID:    b.ID,
				GenreName: genre,
			}
			_, err = u.bookRepo.CreateBookGenre(ctx, args)
			if err != nil {
				return err
			}
		}
		createHistoryArgs := repository.CreateReadingHistoryRequest{
			UserID:    req.UserID,
			BookID:    b.ID,
			Status:    req.Status,
			StartDate: req.StartDate,
			EndDate:   req.EndDate,
		}
		rh, err := u.readingHistoryRepo.Create(ctx, createHistoryArgs)
		if err != nil {
			return err
		}
		book = &entity.Book{
			ID:            b.ID,
			Title:         b.Title,
			Genres:        req.Genres,
			Description:   b.Description,
			CoverImageURL: b.CoverImageURL,
			URL:           b.URL,
			AuthorName:    b.Author,
			PublisherName: b.Publisher,
			PublishDate:   b.PublishDate,
			ISBN:          b.ISBN,
			Status:        rh.Status,
			StartDate:     rh.StartDate,
			EndDate:       rh.EndDate,
		}
		return nil
	})
	return NewRegisterBookResponse(book), handle(err)
}

func (u *RegisterBookUseCaseImpl) createAuthorIfNeed(ctx context.Context, author *string) error {
	if author == nil {
		return nil
	}
	if len(*author) == 0 {
		return nil
	}
	_, err := u.bookRepo.CreateAuthor(ctx, repository.NewCreateAuthorRequest(*author))
	if err != nil {
		return u.checkDuplicateKeyError(err)
	}
	return nil
}

func (u *RegisterBookUseCaseImpl) createPublisherIfNeed(ctx context.Context, publisher *string) error {
	if publisher == nil {
		return nil
	}
	if len(*publisher) == 0 {
		return nil
	}
	_, err := u.bookRepo.CreatePublisher(ctx, repository.NewCreatePublisherRequest(*publisher))
	if err != nil {
		return u.checkDuplicateKeyError(err)
	}
	return err
}

func (u *RegisterBookUseCaseImpl) createGenreIfNeed(ctx context.Context, genre string) error {
	if len(genre) == 0 {
		return nil
	}
	_, err := u.bookRepo.CreateGenre(ctx, repository.NewCreateGenreRequest(genre))
	if err != nil {
		return u.checkDuplicateKeyError(err)
	}
	return nil
}

func (u *RegisterBookUseCaseImpl) checkDuplicateKeyError(err error) error {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) && pqErr.Code != "23505" {
		return err
	}
	return nil
}
