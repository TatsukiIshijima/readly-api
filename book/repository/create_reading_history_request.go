package repository

import (
	"database/sql"
	"readly/book/domain"
	sqlc "readly/db/sqlc"
	"time"
)

type CreateReadingHistoryRequest struct {
	UserID    int64
	BookID    int64
	Status    domain.ReadingStatus
	StartDate *domain.Date
	EndDate   *domain.Date
}

func (r CreateReadingHistoryRequest) toSQLC() sqlc.CreateReadingHistoryParams {
	sd := sql.NullTime{Time: time.Time{}, Valid: false}
	ed := sql.NullTime{Time: time.Time{}, Valid: false}
	if r.StartDate != nil {
		t := r.StartDate.ToTime()
		sd = sql.NullTime{Time: *t, Valid: true}
	}
	if r.EndDate != nil {
		t := r.EndDate.ToTime()
		ed = sql.NullTime{Time: *t, Valid: true}
	}
	return sqlc.CreateReadingHistoryParams{
		UserID:    r.UserID,
		BookID:    r.BookID,
		Status:    r.Status.ToSQLC(),
		StartDate: sd,
		EndDate:   ed,
	}
}
