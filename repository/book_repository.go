package repository

import (
	"context"
	"database/sql"
	db "readly/db/sqlc"
	"time"
)

type BookRepository interface {
	Register(ctx context.Context, args RegisterBookParams) error
}

type BookRepositoryImpl struct {
	store *db.Store
}

func NewBookRepository(store *db.Store) BookRepository {
	return &BookRepositoryImpl{store: store}
}

type RegisterBookParams struct {
	UserID        int64
	BookID        int64
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

func (repo BookRepositoryImpl) Register(ctx context.Context, args RegisterBookParams) error {
	err := repo.store.ExecTx(ctx, func(q *db.Queries) error {
		// 1. Author保存
		_, err := q.CreateAuthor(ctx, args.AuthorName)
		if err != nil {
			return err
		}
		// 2. Publisher保存
		_, err = q.CreatePublisher(ctx, args.PublisherName)
		if err != nil {
			return err
		}
		// 3. Genre保存（同時実行されるかも）
		for _, name := range args.Genres {
			_, err = q.GetGenreByName(ctx, name)
			// すでに存在する場合はスキップ
			if err == nil {
				continue
			}
			_, err = q.CreateGenre(ctx, name)
			if err != nil {
				return err
			}
		}
		// 4. Book保存
		_, err = q.CreateBook(ctx, db.CreateBookParams{
			Title:         sql.NullString{String: args.Title, Valid: true},
			Description:   sql.NullString{String: args.Description, Valid: true},
			CoverImageUrl: sql.NullString{String: args.CoverImageURL, Valid: true},
			Url:           sql.NullString{String: args.URL, Valid: true},
			AuthorName:    args.AuthorName,
			PublisherName: args.PublisherName,
			PublishedDate: sql.NullTime{Time: args.PublishDate, Valid: true},
			Isbn:          sql.NullString{String: args.ISBN, Valid: true},
		})
		// 5. BookGenre保存
		for _, name := range args.Genres {
			_, err = q.CreateBookGenre(ctx, db.CreateBookGenreParams{
				BookID:    args.BookID,
				GenreName: name,
			})
			if err != nil {
				return err
			}
		}
		// 6. ReadingHistory保存
		_, err = q.CreateReadingHistory(ctx, db.CreateReadingHistoryParams{
			UserID:    args.UserID,
			BookID:    args.BookID,
			Status:    db.ReadingStatusUnread,
			StartDate: sql.NullTime{Time: time.Time{}, Valid: true},
			EndDate:   sql.NullTime{Time: time.Time{}, Valid: false},
		})
		if err != nil {
			return err
		}

		return nil
	})
	return err
}
