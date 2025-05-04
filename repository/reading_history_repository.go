package repository

import (
	"context"
	sqlc "readly/db/sqlc"
	"readly/entity"
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
	return &ReadingHistoryRepositoryImpl{
		querier: q,
	}
}

type ReadingStatus int

const (
	Unread ReadingStatus = iota
	Reading
	Done
	Unknown
)

type Convertible interface {
	entity.ReadingStatus | sqlc.ReadingStatus
}

func (status ReadingStatus) toSqlc() sqlc.ReadingStatus {
	switch status {
	case Unread:
		return sqlc.ReadingStatusUnread
	case Reading:
		return sqlc.ReadingStatusReading
	case Done:
		return sqlc.ReadingStatusDone
	default:
		return sqlc.ReadingStatusUnknown
	}
}

func (status ReadingStatus) ToEntity() entity.ReadingStatus {
	switch status {
	case Unread:
		return entity.Unread
	case Reading:
		return entity.Reading
	case Done:
		return entity.Done
	default:
		return entity.Unknown
	}
}

func newFromSqlc(rs sqlc.ReadingStatus) ReadingStatus {
	switch rs {
	case sqlc.ReadingStatusUnread:
		return Unread
	case sqlc.ReadingStatusReading:
		return Reading
	case sqlc.ReadingStatusDone:
		return Done
	default:
		return Unknown
	}
}

func newFromEntity(e entity.ReadingStatus) ReadingStatus {
	switch e {
	case entity.Unread:
		return Unread
	case entity.Reading:
		return Reading
	case entity.Done:
		return Done
	default:
		return Unknown
	}
}

func NewReadingStatus[T Convertible](src T) ReadingStatus {
	switch v := any(src).(type) {
	case entity.ReadingStatus:
		return newFromEntity(v)
	case sqlc.ReadingStatus:
		return newFromSqlc(v)
	default:
		return Unknown
	}

}

func (r *ReadingHistoryRepositoryImpl) Create(ctx context.Context, req CreateReadingHistoryRequest) (*CreateReadingHistoryResponse, error) {
	res, err := r.querier.CreateReadingHistory(ctx, req.toSQLC())
	if err != nil {
		return nil, err
	}
	return newCreateReadingHistoryResponseFromSQLC(res), nil
}

func (r *ReadingHistoryRepositoryImpl) Delete(ctx context.Context, req DeleteReadingHistoryRequest) error {
	rowsAffected, err := r.querier.DeleteReadingHistory(ctx, req.toSQLC())
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrNoRowsDeleted
	}
	return nil
}

type GetReadingHistoryByUserRequest struct {
	UserID int64
	Limit  int32
	Offset int32
}

func (r GetReadingHistoryByUserRequest) toParams() sqlc.GetReadingHistoryByUserParams {
	return sqlc.GetReadingHistoryByUserParams{
		UserID: r.UserID,
		Limit:  r.Limit,
		Offset: r.Offset,
	}
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

func newGetReadingHistoryByUserResponse(r sqlc.GetReadingHistoryByUserRow) GetReadingHistoryByUserResponse {
	id := nilInt64(r.ID)
	t := nilString(r.Title)
	g := newGenres(r.Genres)
	desc := nilString(r.Description)
	coverImgURL := nilString(r.CoverImageUrl)
	URL := nilString(r.Url)
	a := nilString(r.AuthorName)
	p := nilString(r.PublisherName)
	pd := nilTime(r.PublishedDate)
	ISBN := nilString(r.Isbn)
	s := NewReadingStatus[sqlc.ReadingStatus](r.Status)
	sd := nilTime(r.StartDate)
	ed := nilTime(r.EndDate)
	return GetReadingHistoryByUserResponse{
		BookID:        *id,
		Title:         *t,
		Genres:        g,
		Description:   desc,
		CoverImageURL: coverImgURL,
		URL:           URL,
		AuthorName:    a,
		PublisherName: p,
		PublishDate:   pd,
		ISBN:          ISBN,
		Status:        s,
		StartDate:     sd,
		EndDate:       ed,
	}
}

func (r *ReadingHistoryRepositoryImpl) GetByUser(ctx context.Context, req GetReadingHistoryByUserRequest) ([]GetReadingHistoryByUserResponse, error) {
	rows, err := r.querier.GetReadingHistoryByUser(ctx, req.toParams())
	if err != nil {
		return nil, err
	}
	res := make([]GetReadingHistoryByUserResponse, len(rows))
	for i := 0; i < len(rows); i++ {
		getResponse := newGetReadingHistoryByUserResponse(rows[i])
		res[i] = getResponse
	}
	return res, nil
}

type GetReadingHistoryByUserAndBookRequest struct {
	UserID int64
	BookID int64
}

func (r GetReadingHistoryByUserAndBookRequest) toParams() sqlc.GetReadingHistoryByUserAndBookParams {
	return sqlc.GetReadingHistoryByUserAndBookParams{
		UserID: r.UserID,
		BookID: r.BookID,
	}
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

func newGetReadingHistoryByUserAndBookResponse(r sqlc.GetReadingHistoryByUserAndBookRow) *GetReadingHistoryByUserAndBookResponse {
	id := nilInt64(r.ID)
	t := nilString(r.Title)
	g := newGenres(r.Genres)
	desc := nilString(r.Description)
	coverImgURL := nilString(r.CoverImageUrl)
	URL := nilString(r.Url)
	a := nilString(r.AuthorName)
	p := nilString(r.PublisherName)
	pd := nilTime(r.PublishedDate)
	ISBN := nilString(r.Isbn)
	s := NewReadingStatus[sqlc.ReadingStatus](r.Status)
	sd := nilTime(r.StartDate)
	ed := nilTime(r.EndDate)
	return &GetReadingHistoryByUserAndBookResponse{
		BookID:        *id,
		Title:         *t,
		Genres:        g,
		Description:   desc,
		CoverImageURL: coverImgURL,
		URL:           URL,
		AuthorName:    a,
		PublisherName: p,
		PublishDate:   pd,
		ISBN:          ISBN,
		Status:        s,
		StartDate:     sd,
		EndDate:       ed,
	}
}

func (r *ReadingHistoryRepositoryImpl) GetByUserAndBook(ctx context.Context, req GetReadingHistoryByUserAndBookRequest) (*GetReadingHistoryByUserAndBookResponse, error) {
	row, err := r.querier.GetReadingHistoryByUserAndBook(ctx, req.toParams())
	if err != nil {
		return nil, err
	}
	return newGetReadingHistoryByUserAndBookResponse(row), nil
}

type GetReadingHistoryByUserAndStatusRequest struct {
	UserID int64
	Status ReadingStatus
	Limit  int32
	Offset int32
}

func (r GetReadingHistoryByUserAndStatusRequest) toParams() sqlc.GetReadingHistoryByUserAndStatusParams {
	return sqlc.GetReadingHistoryByUserAndStatusParams{
		UserID: r.UserID,
		Status: r.Status.toSqlc(),
		Limit:  r.Limit,
		Offset: r.Offset,
	}
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

func newGetReadingHistoryByUserAndStatusResponse(r sqlc.GetReadingHistoryByUserAndStatusRow) GetReadingHistoryByUserAndStatusResponse {
	id := nilInt64(r.ID)
	t := nilString(r.Title)
	g := newGenres(r.Genres)
	desc := nilString(r.Description)
	coverImgURL := nilString(r.CoverImageUrl)
	URL := nilString(r.Url)
	a := nilString(r.AuthorName)
	p := nilString(r.PublisherName)
	pd := nilTime(r.PublishedDate)
	ISBN := nilString(r.Isbn)
	s := NewReadingStatus[sqlc.ReadingStatus](r.Status)
	sd := nilTime(r.StartDate)
	ed := nilTime(r.EndDate)
	return GetReadingHistoryByUserAndStatusResponse{
		BookID:        *id,
		Title:         *t,
		Genres:        g,
		Description:   desc,
		CoverImageURL: coverImgURL,
		URL:           URL,
		AuthorName:    a,
		PublisherName: p,
		PublishDate:   pd,
		ISBN:          ISBN,
		Status:        s,
		StartDate:     sd,
		EndDate:       ed,
	}
}

func (r *ReadingHistoryRepositoryImpl) GetByUserAndStatus(ctx context.Context, req GetReadingHistoryByUserAndStatusRequest) ([]GetReadingHistoryByUserAndStatusResponse, error) {
	rows, err := r.querier.GetReadingHistoryByUserAndStatus(ctx, req.toParams())
	if err != nil {
		return nil, err
	}
	res := make([]GetReadingHistoryByUserAndStatusResponse, len(rows))
	for i := 0; i < len(rows); i++ {
		getResponse := newGetReadingHistoryByUserAndStatusResponse(rows[i])
		res[i] = getResponse
	}
	return res, nil
}

func (r *ReadingHistoryRepositoryImpl) Update(ctx context.Context, req UpdateReadingHistoryRequest) (*UpdateReadingHistoryResponse, error) {
	h, err := r.querier.UpdateReadingHistory(ctx, req.toSQLC())
	if err != nil {
		return nil, err
	}
	return newUpdateReadingHistoryResponseFromSQLC(h), nil
}
