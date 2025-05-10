package repository

import (
	"database/sql"
	sqlc "readly/db/sqlc"
	"readly/entity"
	"time"
)

type UpdateReadingHistoryRequest struct {
	UserID    int64
	BookID    int64
	Status    entity.ReadingStatus
	StartDate *entity.Date
	EndDate   *entity.Date
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
