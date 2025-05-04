package repository

import sqlc "readly/db/sqlc"

type DeleteReadingHistoryRequest struct {
	BookID int64
	UserID int64
}

func (r DeleteReadingHistoryRequest) toSQLC() sqlc.DeleteReadingHistoryParams {
	return sqlc.DeleteReadingHistoryParams{
		BookID: r.BookID,
		UserID: r.UserID,
	}
}
