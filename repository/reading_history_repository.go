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
	id := newInt64(r.ID)
	t := newString(r.Title)
	g := newGenres(r.Genres)
	desc := newString(r.Description)
	coverImgURL := newString(r.CoverImageUrl)
	URL := newString(r.Url)
	a := newString(r.AuthorName)
	p := newString(r.PublisherName)
	pd := newTime(r.PublishedDate)
	ISBN := newString(r.Isbn)
	s := newReadingStatus(r.Status)
	sd := newTime(r.StartDate)
	ed := newTime(r.EndDate)
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

func (r ReadingHistoryRepositoryImpl) GetByUser(ctx context.Context, req GetReadingHistoryByUserRequest) ([]GetReadingHistoryByUserResponse, error) {
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
	id := newInt64(r.ID)
	t := newString(r.Title)
	g := newGenres(r.Genres)
	desc := newString(r.Description)
	coverImgURL := newString(r.CoverImageUrl)
	URL := newString(r.Url)
	a := newString(r.AuthorName)
	p := newString(r.PublisherName)
	pd := newTime(r.PublishedDate)
	ISBN := newString(r.Isbn)
	s := newReadingStatus(r.Status)
	sd := newTime(r.StartDate)
	ed := newTime(r.EndDate)
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

func (r ReadingHistoryRepositoryImpl) GetByUserAndBook(ctx context.Context, req GetReadingHistoryByUserAndBookRequest) (*GetReadingHistoryByUserAndBookResponse, error) {
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
		Status: r.Status.value(),
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
	id := newInt64(r.ID)
	t := newString(r.Title)
	g := newGenres(r.Genres)
	desc := newString(r.Description)
	coverImgURL := newString(r.CoverImageUrl)
	URL := newString(r.Url)
	a := newString(r.AuthorName)
	p := newString(r.PublisherName)
	pd := newTime(r.PublishedDate)
	ISBN := newString(r.Isbn)
	s := newReadingStatus(r.Status)
	sd := newTime(r.StartDate)
	ed := newTime(r.EndDate)
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

func (r ReadingHistoryRepositoryImpl) GetByUserAndStatus(ctx context.Context, req GetReadingHistoryByUserAndStatusRequest) ([]GetReadingHistoryByUserAndStatusResponse, error) {
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

type UpdateReadingHistoryRequest struct {
	UserID    int64
	BookID    int64
	Status    ReadingStatus
	StartDate *time.Time
	EndDate   *time.Time
}

func (r UpdateReadingHistoryRequest) toParams() sqlc.UpdateReadingHistoryParams {
	sd := sql.NullTime{Time: time.Time{}, Valid: false}
	ed := sql.NullTime{Time: time.Time{}, Valid: false}
	if r.StartDate != nil {
		sd = sql.NullTime{Time: *r.StartDate, Valid: true}
	}
	if r.EndDate != nil {
		ed = sql.NullTime{Time: *r.EndDate, Valid: true}
	}
	return sqlc.UpdateReadingHistoryParams{
		UserID:    r.UserID,
		BookID:    r.BookID,
		Status:    r.Status.value(),
		StartDate: sd,
		EndDate:   ed,
	}
}

type UpdateReadingHistoryResponse struct {
	BookID    int64
	Status    ReadingStatus
	StartDate *time.Time
	EndDate   *time.Time
}

func newUpdateReadingHistoryResponse(r sqlc.ReadingHistory) *UpdateReadingHistoryResponse {
	bid := r.BookID
	s := newReadingStatus(r.Status)
	sd := newTime(r.StartDate)
	ed := newTime(r.EndDate)
	return &UpdateReadingHistoryResponse{
		BookID:    bid,
		Status:    s,
		StartDate: sd,
		EndDate:   ed,
	}
}

func (r ReadingHistoryRepositoryImpl) Update(ctx context.Context, req UpdateReadingHistoryRequest) (*UpdateReadingHistoryResponse, error) {
	h, err := r.querier.UpdateReadingHistory(ctx, req.toParams())
	if err != nil {
		return nil, err
	}
	return newUpdateReadingHistoryResponse(h), nil
}
