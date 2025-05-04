package repository

import (
	"context"
	sqlc "readly/db/sqlc"
)

type UserRepository interface {
	CreateUser(ctx context.Context, req CreateUserRequest) (*CreateUserResponse, error)
	DeleteUser(ctx context.Context, req DeleteUserRequest) error
	GetUserByEmail(ctx context.Context, req GetUserByEmailRequest) (*GetUserResponse, error)
	GetUserByID(ctx context.Context, req GetUserByIDRequest) (*GetUserResponse, error)
	UpdateUser(ctx context.Context, req UpdateRequest) (*UpdateResponse, error)
}

type UserRepositoryImpl struct {
	querier sqlc.Querier
}

func NewUserRepository(q sqlc.Querier) UserRepository {
	return &UserRepositoryImpl{
		querier: q,
	}
}

func (r *UserRepositoryImpl) CreateUser(ctx context.Context, req CreateUserRequest) (*CreateUserResponse, error) {
	res, err := r.querier.CreateUser(ctx, req.toSQLC())
	if err != nil {
		return nil, err
	}
	return newCreateUserResponseFromSQLC(res), nil
}

func (r *UserRepositoryImpl) DeleteUser(ctx context.Context, req DeleteUserRequest) error {
	err := r.querier.DeleteUser(ctx, req.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) GetUserByEmail(ctx context.Context, req GetUserByEmailRequest) (*GetUserResponse, error) {
	res, err := r.querier.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	return newGetUserResponseFromSQLC(res), nil
}

func (r *UserRepositoryImpl) GetUserByID(ctx context.Context, req GetUserByIDRequest) (*GetUserResponse, error) {
	res, err := r.querier.GetUserByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return newGetUserResponseFromSQLC(res), nil
}

type UpdateRequest struct {
	ID       int64
	Name     string
	Email    string
	Password string
}

type UpdateResponse struct {
	ID    int64
	Name  string
	Email string
}

func (r *UserRepositoryImpl) UpdateUser(ctx context.Context, req UpdateRequest) (*UpdateResponse, error) {
	args := sqlc.UpdateUserParams{
		ID:             req.ID,
		Name:           req.Name,
		Email:          req.Email,
		HashedPassword: req.Password,
	}
	res, err := r.querier.UpdateUser(ctx, args)
	if err != nil {
		return nil, err
	}
	u := &UpdateResponse{
		ID:    res.ID,
		Name:  res.Name,
		Email: res.Email,
	}
	return u, nil
}
