package repository

import (
	"database/sql"
	sqlc "readly/db/sqlc"
	"time"
)

type UpdateReadingHistoryRequest struct {
	UserID    int64
	BookID    int64
	Status    ReadingStatus
	StartDate *time.Time
	EndDate   *time.Time
}

func (r UpdateReadingHistoryRequest) toSQLC() sqlc.UpdateReadingHistoryParams {
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
		Status:    r.Status.toSqlc(),
		StartDate: sd,
		EndDate:   ed,
	}
}
