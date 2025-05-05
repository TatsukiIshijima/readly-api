package repository

import (
	sqlc "readly/db/sqlc"
	"readly/entity"
)

type CreateReadingHistoryResponse struct {
	BookID    int64
	Status    entity.ReadingStatus
	StartDate *entity.Date
	EndDate   *entity.Date
}

func newCreateReadingHistoryResponseFromSQLC(r sqlc.ReadingHistory) *CreateReadingHistoryResponse {
	return &CreateReadingHistoryResponse{
		BookID:    r.BookID,
		Status:    entity.NewReadingStatusFromSQLC(r.Status),
		StartDate: entity.NewDateEntityFromNullTime(r.StartDate),
		EndDate:   entity.NewDateEntityFromNullTime(r.EndDate),
	}
}
