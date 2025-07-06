package repository

import (
	sqlc "readly/db/sqlc"
	"readly/feature/book/domain"
)

type GetReadingHistoryByUserRequest struct {
	UserID int64
	Limit  int32
	Offset int32
}

func (r GetReadingHistoryByUserRequest) toSQLC() sqlc.GetReadingHistoryByUserParams {
	return sqlc.GetReadingHistoryByUserParams{
		UserID: r.UserID,
		Limit:  r.Limit,
		Offset: r.Offset,
	}
}

type GetReadingHistoryByUserAndBookRequest struct {
	UserID int64
	BookID int64
}

func (r GetReadingHistoryByUserAndBookRequest) toSQLC() sqlc.GetReadingHistoryByUserAndBookParams {
	return sqlc.GetReadingHistoryByUserAndBookParams{
		UserID: r.UserID,
		BookID: r.BookID,
	}
}

type GetReadingHistoryByUserAndStatusRequest struct {
	UserID int64
	Status domain.ReadingStatus
	Limit  int32
	Offset int32
}

func (r GetReadingHistoryByUserAndStatusRequest) toSQLC() sqlc.GetReadingHistoryByUserAndStatusParams {
	return sqlc.GetReadingHistoryByUserAndStatusParams{
		UserID: r.UserID,
		Status: r.Status.ToSQLC(),
		Limit:  r.Limit,
		Offset: r.Offset,
	}
}
