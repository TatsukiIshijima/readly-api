package repository

import (
	"readly/book/domain"
	sqlc "readly/db/sqlc"
)

type CreateReadingHistoryResponse struct {
	BookID    int64
	Status    domain.ReadingStatus
	StartDate *domain.Date
	EndDate   *domain.Date
}

func newCreateReadingHistoryResponseFromSQLC(r sqlc.ReadingHistory) *CreateReadingHistoryResponse {
	return &CreateReadingHistoryResponse{
		BookID:    r.BookID,
		Status:    domain.NewReadingStatusFromSQLC(r.Status),
		StartDate: domain.NewDateEntityFromNullTime(r.StartDate),
		EndDate:   domain.NewDateEntityFromNullTime(r.EndDate),
	}
}
