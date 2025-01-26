package usecase

import (
	"context"
	"errors"
	"github.com/lib/pq"
	"readly/entity"
	"readly/repository"
	"time"
)

type RegisterBookUseCase struct {
	Executor
	transactor         repository.Transactor
	bookRepo           repository.BookRepository
	readingHistoryRepo repository.ReadingHistoryRepository
	userRepo           repository.UserRepository
}

func NewRegisterBookUseCase(transactor repository.Transactor, bookRepo repository.BookRepository, readingHistoryRepo repository.ReadingHistoryRepository, userRepo repository.UserRepository) RegisterBookUseCase {
	return RegisterBookUseCase{
		transactor:         transactor,
		bookRepo:           bookRepo,
		readingHistoryRepo: readingHistoryRepo,
		userRepo:           userRepo,
	}
}

type RegisterBookRequest struct {
	UserID        int64
	Title         string
	Genres        []string
	Description   *string
	CoverImageURL *string
	URL           *string
	AuthorName    *string
	PublisherName *string
	PublishDate   *time.Time
	ISBN          *string
	Status        entity.ReadingStatus
	StartDate     *time.Time
	EndDate       *time.Time
}

func (u RegisterBookUseCase) RegisterBook(ctx context.Context, req RegisterBookRequest) (*entity.Book, error) {
	var res *entity.Book
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
			Status:    repository.NewReadingStatus[entity.ReadingStatus](req.Status),
			StartDate: req.StartDate,
			EndDate:   req.EndDate,
		}
		rh, err := u.readingHistoryRepo.Create(ctx, createHistoryArgs)
		if err != nil {
			return err
		}
		res = &entity.Book{
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
			Status:        rh.Status.ToEntity(),
			StartDate:     rh.StartDate,
			EndDate:       rh.EndDate,
		}
		return nil
	})
	return res, handle(err)
}

func (u RegisterBookUseCase) createAuthorIfNeed(ctx context.Context, author *string) error {
	if author == nil {
		return nil
	}
	if len(*author) == 0 {
		return nil
	}
	_, err := u.bookRepo.CreateAuthor(ctx, *author)
	if err != nil {
		return u.checkDuplicateKeyError(err)
	}
	return nil
}

func (u RegisterBookUseCase) createPublisherIfNeed(ctx context.Context, publisher *string) error {
	if publisher == nil {
		return nil
	}
	if len(*publisher) == 0 {
		return nil
	}
	_, err := u.bookRepo.CreatePublisher(ctx, *publisher)
	if err != nil {
		return u.checkDuplicateKeyError(err)
	}
	return err
}

func (u RegisterBookUseCase) createGenreIfNeed(ctx context.Context, genre string) error {
	if len(genre) == 0 {
		return nil
	}
	_, err := u.bookRepo.CreateGenre(ctx, genre)
	if err != nil {
		return u.checkDuplicateKeyError(err)
	}
	return nil
}

func (u RegisterBookUseCase) checkDuplicateKeyError(err error) error {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) && pqErr.Code != "23505" {
		return err
	}
	return nil
}
