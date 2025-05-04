package repository

import (
	sqlc "readly/db/sqlc"
	"readly/entity"
)

type UpdateReadingHistoryResponse struct {
	BookID    int64
	Status    entity.ReadingStatus
	StartDate *entity.Date
	EndDate   *entity.Date
}

func newUpdateReadingHistoryResponseFromSQLC(r sqlc.ReadingHistory) *UpdateReadingHistoryResponse {
	return &UpdateReadingHistoryResponse{
		BookID:    r.BookID,
		Status:    entity.NewReadingStatusFromSQLC(r.Status),
		StartDate: entity.NewDateEntityFromNullTime(r.StartDate),
		EndDate:   entity.NewDateEntityFromNullTime(r.EndDate),
	}
}
