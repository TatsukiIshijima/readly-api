package repository

import (
	"readly/book/domain"
	sqlc "readly/db/sqlc"
)

type UpdateReadingHistoryResponse struct {
	BookID    int64
	Status    domain.ReadingStatus
	StartDate *domain.Date
	EndDate   *domain.Date
}

func newUpdateReadingHistoryResponseFromSQLC(r sqlc.ReadingHistory) *UpdateReadingHistoryResponse {
	return &UpdateReadingHistoryResponse{
		BookID:    r.BookID,
		Status:    domain.NewReadingStatusFromSQLC(r.Status),
		StartDate: domain.NewDateEntityFromNullTime(r.StartDate),
		EndDate:   domain.NewDateEntityFromNullTime(r.EndDate),
	}
}
