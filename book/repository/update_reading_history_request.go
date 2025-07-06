package repository

import (
	"database/sql"
	"readly/book/domain"
	sqlc "readly/db/sqlc"
	"time"
)

type UpdateReadingHistoryRequest struct {
	UserID    int64
	BookID    int64
	Status    domain.ReadingStatus
	StartDate *domain.Date
	EndDate   *domain.Date
}

func (r UpdateReadingHistoryRequest) toSQLC() sqlc.UpdateReadingHistoryParams {
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
	return sqlc.UpdateReadingHistoryParams{
		UserID:    r.UserID,
		BookID:    r.BookID,
		Status:    r.Status.ToSQLC(),
		StartDate: sd,
		EndDate:   ed,
	}
}
