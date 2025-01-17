package repository

import (
	"context"
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
	querier sqlc.Container
}

func NewReadingHistoryRepository(q sqlc.Container) ReadingHistoryRepository {
	return ReadingHistoryRepositoryImpl{
		querier: q,
	}
}

type CreateReadingHistoryRequest struct {
}

type CreateReadingHistoryResponse struct {
}

func (r ReadingHistoryRepositoryImpl) Create(ctx context.Context, req CreateReadingHistoryRequest) (*CreateReadingHistoryResponse, error) {

}

type DeleteReadingHistoryRequest struct {
}

func (r ReadingHistoryRepositoryImpl) Delete(ctx context.Context, req DeleteReadingHistoryRequest) error {

}

type GetReadingHistoryByUserRequest struct {
}

type GetReadingHistoryByUserResponse struct {
}

func (r ReadingHistoryRepositoryImpl) GetByUser(ctx context.Context, req GetReadingHistoryByUserRequest) ([]GetReadingHistoryByUserResponse, error) {

}

type GetReadingHistoryByUserAndBookRequest struct {
}

type GetReadingHistoryByUserAndBookResponse struct {
}

func (r ReadingHistoryRepositoryImpl) GetByUserAndBook(ctx context.Context, req GetReadingHistoryByUserAndBookRequest) (*GetReadingHistoryByUserAndBookResponse, error) {

}

type GetReadingHistoryByUserAndStatusRequest struct {
}

type GetReadingHistoryByUserAndStatusResponse struct {
}

func (r ReadingHistoryRepositoryImpl) GetByUserAndStatus(ctx context.Context, req GetReadingHistoryByUserAndStatusRequest) ([]GetReadingHistoryByUserAndStatusResponse, error) {

}

type UpdateReadingHistoryRequest struct {
}

type UpdateReadingHistoryResponse struct {
}

func (r ReadingHistoryRepositoryImpl) Update(ctx context.Context, req UpdateReadingHistoryRequest) (*UpdateReadingHistoryResponse, error) {

}
