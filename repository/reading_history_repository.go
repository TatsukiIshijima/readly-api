package repository

import (
	"context"
	"database/sql"
	sqlc "readly/db/sqlc"
	"time"
)

type ReadingHistoryRepository interface {
	Create(ctx context.Context, req CreateReadingHistoryRequest) (*CreateReadingHistoryResponse, error)
	Delete(ctx context.Context, req DeleteReadingHistoryRequest) error
	GetByUser(ctx context.Context, req GetReadingHistoryByUserRequest) ([]GetReadingHistoryByUserResponse, error)
	GetByUserAndBook(ctx context.Context, req GetReadingHistoryByUserAndBookRequest) (*GetReadingHistoryByUserAndBookResponse, error)
	GetByUserAndStatus(ctx context.Context, req GetReadingHistoryByUserAndStatusRequest) ([]GetReadingHistoryByUserAndStatusResponse, error)
	Update(ctx context.Context, req UpdateReadingHistoryRequest) (*UpdateReadingHistoryResponse, error)
}

type ReadingHistoryRepositoryImpl struct {
	querier sqlc.Querier
}

func NewReadingHistoryRepository(q sqlc.Querier) ReadingHistoryRepository {
	return ReadingHistoryRepositoryImpl{
		querier: q,
	}
}

type ReadingStatus int

const (
	Unread ReadingStatus = iota
	Reading
	Done
)

func (status ReadingStatus) value() sqlc.ReadingStatus {
	switch status {
	case Unread:
		return sqlc.ReadingStatusUnread
	case Reading:
		return sqlc.ReadingStatusReading
	case Done:
		return sqlc.ReadingStatusDone
	default:
		panic("invalid reading status")
	}
}

func newReadingStatus(rs sqlc.ReadingStatus) ReadingStatus {
	switch rs {
	case sqlc.ReadingStatusUnread:
		return Unread
	case sqlc.ReadingStatusReading:
		return Reading
	case sqlc.ReadingStatusDone:
		return Done
	default:
		panic("invalid reading status")
	}
}

type CreateReadingHistoryRequest struct {
	UserID    int64
	BookID    int64
	Status    ReadingStatus
	StartDate *time.Time
	EndDate   *time.Time
}

func (r CreateReadingHistoryRequest) toParams() sqlc.CreateReadingHistoryParams {
	sd := sql.NullTime{Time: time.Time{}, Valid: false}
	ed := sql.NullTime{Time: time.Time{}, Valid: false}
	if r.StartDate != nil {
		sd = sql.NullTime{Time: *r.StartDate, Valid: true}
	}
	if r.EndDate != nil {
		ed = sql.NullTime{Time: *r.EndDate, Valid: true}
	}
	return sqlc.CreateReadingHistoryParams{
		UserID:    r.UserID,
		BookID:    r.BookID,
		Status:    r.Status.value(),
		StartDate: sd,
		EndDate:   ed,
	}
}

type CreateReadingHistoryResponse struct {
	BookID    int64
	Status    ReadingStatus
	StartDate *time.Time
	EndDate   *time.Time
}

func newCreateReadingHistoryResponse(r sqlc.ReadingHistory) *CreateReadingHistoryResponse {
	return &CreateReadingHistoryResponse{
		BookID:    r.BookID,
		Status:    newReadingStatus(r.Status),
		StartDate: &r.StartDate.Time,
		EndDate:   &r.EndDate.Time,
	}
}

func (r ReadingHistoryRepositoryImpl) Create(ctx context.Context, req CreateReadingHistoryRequest) (*CreateReadingHistoryResponse, error) {
	h, err := r.querier.CreateReadingHistory(ctx, req.toParams())
	if err != nil {
		return nil, err
	}
	return newCreateReadingHistoryResponse(h), nil
}

type DeleteReadingHistoryRequest struct {
	UserID int64
	BookID int64
}

func (r ReadingHistoryRepositoryImpl) Delete(ctx context.Context, req DeleteReadingHistoryRequest) error {
	err := r.querier.DeleteReadingHistory(ctx, sqlc.DeleteReadingHistoryParams{
		UserID: req.UserID,
		BookID: req.BookID,
	})
	if err != nil {
		return err
	}
	return nil
}

type GetReadingHistoryByUserRequest struct {
	UserID int64
	Limit  int32
	Offset int32
}

type GetReadingHistoryByUserResponse struct {
	BookID        int64
	Title         string
	Genres        []string
	Description   *string
	CoverImageURL *string
	URL           *string
	AuthorName    *string
	PublisherName *string
	PublishDate   *time.Time
	ISBN          *string
	Status        ReadingStatus
	StartDate     *time.Time
	EndDate       *time.Time
}

func (r ReadingHistoryRepositoryImpl) GetByUser(ctx context.Context, req GetReadingHistoryByUserRequest) ([]GetReadingHistoryByUserResponse, error) {

}

type GetReadingHistoryByUserAndBookRequest struct {
	UserID int64
	BookID int64
}

type GetReadingHistoryByUserAndBookResponse struct {
	BookID        int64
	Title         string
	Genres        []string
	Description   *string
	CoverImageURL *string
	URL           *string
	AuthorName    *string
	PublisherName *string
	PublishDate   *time.Time
	ISBN          *string
	Status        ReadingStatus
	StartDate     *time.Time
	EndDate       *time.Time
}

func (r ReadingHistoryRepositoryImpl) GetByUserAndBook(ctx context.Context, req GetReadingHistoryByUserAndBookRequest) (*GetReadingHistoryByUserAndBookResponse, error) {

}

type GetReadingHistoryByUserAndStatusRequest struct {
	UserID int64
	Status ReadingStatus
	Limit  int32
	Offset int32
}

type GetReadingHistoryByUserAndStatusResponse struct {
	BookID        int64
	Title         string
	Genres        []string
	Description   *string
	CoverImageURL *string
	URL           *string
	AuthorName    *string
	PublisherName *string
	PublishDate   *time.Time
	ISBN          *string
	Status        ReadingStatus
	StartDate     *time.Time
	EndDate       *time.Time
}

func (r ReadingHistoryRepositoryImpl) GetByUserAndStatus(ctx context.Context, req GetReadingHistoryByUserAndStatusRequest) ([]GetReadingHistoryByUserAndStatusResponse, error) {

}

type UpdateReadingHistoryRequest struct {
	UserID    int64
	BookID    int64
	Status    ReadingStatus
	StartDate *time.Time
	EndDate   *time.Time
}

type UpdateReadingHistoryResponse struct {
	BookID        int64
	Title         string
	Genres        []string
	Description   *string
	CoverImageURL *string
	URL           *string
	AuthorName    *string
	PublisherName *string
	PublishDate   *time.Time
	ISBN          *string
	Status        ReadingStatus
	StartDate     *time.Time
	EndDate       *time.Time
}

func (r ReadingHistoryRepositoryImpl) Update(ctx context.Context, req UpdateReadingHistoryRequest) (*UpdateReadingHistoryResponse, error) {

}
