package repository

import (
	"context"
	"database/sql"
	"errors"
	db "readly/db/sqlc"
	"time"
)

type BookRepository interface {
	Register(ctx context.Context, args RegisterBookParams) error
}

type BookRepositoryImpl struct {
	store *Store
}

func NewBookRepository(store *Store) BookRepository {
	return &BookRepositoryImpl{store: store}
}

type RegisterBookParams struct {
	UserID        int64
	Title         string
	Genres        []string
	Description   string
	CoverImageURL string
	URL           string
	AuthorName    string
	PublisherName string
	PublishDate   time.Time
	ISBN          string
}

func (r BookRepositoryImpl) Register(ctx context.Context, args RegisterBookParams) error {
	err := r.store.execTx(ctx, func(q *db.Queries) error {
		if err := r.registerAuthorIfNotExist(ctx, q, args.AuthorName); err != nil {
			return err
		}
		if err := r.registerPublisherIfNotExist(ctx, q, args.PublisherName); err != nil {
			return err
		}
		for _, genre := range args.Genres {
			if err := r.registerGenreIfNotExist(ctx, q, genre); err != nil {
				return err
			}
		}
		book, err := q.CreateBook(ctx, db.CreateBookParams{
			Title:         sql.NullString{String: args.Title, Valid: true},
			Description:   sql.NullString{String: args.Description, Valid: true},
			CoverImageUrl: sql.NullString{String: args.CoverImageURL, Valid: true},
			Url:           sql.NullString{String: args.URL, Valid: true},
			AuthorName:    args.AuthorName,
			PublisherName: args.PublisherName,
			PublishedDate: sql.NullTime{Time: args.PublishDate, Valid: true},
			Isbn:          sql.NullString{String: args.ISBN, Valid: true},
		})
		if err != nil {
			return err
		}
		for _, genre := range args.Genres {
			if _, err := q.CreateBookGenre(ctx, db.CreateBookGenreParams{
				BookID:    book.ID,
				GenreName: genre,
			}); err != nil {
				return err
			}
		}
		if _, err := q.CreateReadingHistory(ctx, db.CreateReadingHistoryParams{
			UserID:    args.UserID,
			BookID:    book.ID,
			Status:    db.ReadingStatusUnread,
			StartDate: sql.NullTime{Time: time.Time{}, Valid: true},
			EndDate:   sql.NullTime{Time: time.Time{}, Valid: false},
		}); err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r BookRepositoryImpl) registerAuthorIfNotExist(ctx context.Context, q *db.Queries, name string) error {
	var err error
	_, err = q.GetAuthorByName(ctx, name)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
		_, err = q.CreateAuthor(ctx, name)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r BookRepositoryImpl) registerPublisherIfNotExist(ctx context.Context, q *db.Queries, name string) error {
	_, err := q.GetPublisherByName(ctx, name)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
		_, err = q.CreatePublisher(ctx, name)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r BookRepositoryImpl) registerGenreIfNotExist(ctx context.Context, q *db.Queries, name string) error {
	_, err := q.GetGenreByName(ctx, name)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
		_, err = q.CreateGenre(ctx, name)
		if err != nil {
			return err
		}
	}
	return nil
}
