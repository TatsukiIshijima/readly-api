package repository

import (
	"context"
	"database/sql"
	"errors"
	sqlc "readly/db/sqlc"
)

type ReadingHistoryRepository interface {
	Create(ctx context.Context, req CreateReadingHistoryRequest) (*CreateReadingHistoryResponse, error)
	Delete(ctx context.Context, req DeleteReadingHistoryRequest) error
	GetByUser(ctx context.Context, req GetReadingHistoryByUserRequest) ([]GetReadingHistoryByUserResponse, error)
	GetByUserAndBook(ctx context.Context, req GetReadingHistoryByUserAndBookRequest) (*GetReadingHistoryByUserAndBookResponse, error)
	GetByUserAndStatus(ctx context.Context, req GetReadingHistoryByUserAndStatusRequest) ([]GetReadingHistoryByUserAndStatusResponse, error)
	Update(ctx context.Context, req UpdateReadingHistoryRequest) (*UpdateReadingHistoryResponse, error)
}

type ReadingHistoryRepositoryImpl struct {
	querier sqlc.Querier
}

func NewReadingHistoryRepository(q sqlc.Querier) ReadingHistoryRepository {
	return &ReadingHistoryRepositoryImpl{
		querier: q,
	}
}

func (r *ReadingHistoryRepositoryImpl) Create(ctx context.Context, req CreateReadingHistoryRequest) (*CreateReadingHistoryResponse, error) {
	res, err := r.querier.CreateReadingHistory(ctx, req.toSQLC())
	if err != nil {
		return nil, err
	}
	return newCreateReadingHistoryResponseFromSQLC(res), nil
}

func (r *ReadingHistoryRepositoryImpl) Delete(ctx context.Context, req DeleteReadingHistoryRequest) error {
	rowsAffected, err := r.querier.DeleteReadingHistory(ctx, req.toSQLC())
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrNoRowsDeleted
	}
	return nil
}

func (r *ReadingHistoryRepositoryImpl) GetByUser(ctx context.Context, req GetReadingHistoryByUserRequest) ([]GetReadingHistoryByUserResponse, error) {
	rows, err := r.querier.GetReadingHistoryByUser(ctx, req.toSQLC())
	if err != nil {
		return nil, err
	}
	res := make([]GetReadingHistoryByUserResponse, len(rows))
	for i := 0; i < len(rows); i++ {
		getResponse := newGetReadingHistoryByUserResponseFromSQLC(rows[i])
		res[i] = getResponse
	}
	return res, nil
}

func (r *ReadingHistoryRepositoryImpl) GetByUserAndBook(ctx context.Context, req GetReadingHistoryByUserAndBookRequest) (*GetReadingHistoryByUserAndBookResponse, error) {
	row, err := r.querier.GetReadingHistoryByUserAndBook(ctx, req.toSQLC())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRows
		}
		return nil, err
	}
	return newGetReadingHistoryByUserAndBookResponseFromSQLC(row), nil
}

func (r *ReadingHistoryRepositoryImpl) GetByUserAndStatus(ctx context.Context, req GetReadingHistoryByUserAndStatusRequest) ([]GetReadingHistoryByUserAndStatusResponse, error) {
	rows, err := r.querier.GetReadingHistoryByUserAndStatus(ctx, req.toSQLC())
	if err != nil {
		return nil, err
	}
	res := make([]GetReadingHistoryByUserAndStatusResponse, len(rows))
	for i := 0; i < len(rows); i++ {
		getResponse := newGetReadingHistoryByUserAndStatusResponseFromSQLC(rows[i])
		res[i] = getResponse
	}
	return res, nil
}

func (r *ReadingHistoryRepositoryImpl) Update(ctx context.Context, req UpdateReadingHistoryRequest) (*UpdateReadingHistoryResponse, error) {
	h, err := r.querier.UpdateReadingHistory(ctx, req.toSQLC())
	if err != nil {
		return nil, err
	}
	return newUpdateReadingHistoryResponseFromSQLC(h), nil
}
