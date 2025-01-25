package repository

import (
	"context"
	sqlc "readly/db/sqlc"
)

type UserRepository interface {
	CreateUser(ctx context.Context, req CreateUserRequest) (*CreateUserResponse, error)
	DeleteUser(ctx context.Context, id int64) error
	GetUserByEmail(ctx context.Context, email string) (*GetUserResponse, error)
	GetUserByID(ctx context.Context, id int64) (*GetUserResponse, error)
	UpdateUser(ctx context.Context, req UpdateRequest) (*UpdateResponse, error)
}

type UserRepositoryImpl struct {
	querier sqlc.Querier
}

func NewUserRepository(q sqlc.Querier) UserRepositoryImpl {
	return UserRepositoryImpl{
		querier: q,
	}
}

type CreateUserRequest struct {
	Name     string
	Email    string
	Password string
}

type CreateUserResponse struct {
	ID    int64
	Name  string
	Email string
}

func (r UserRepositoryImpl) CreateUser(ctx context.Context, req CreateUserRequest) (*CreateUserResponse, error) {
	args := sqlc.CreateUserParams{
		Name:           req.Name,
		Email:          req.Email,
		HashedPassword: req.Password,
	}
	res, err := r.querier.CreateUser(ctx, args)
	if err != nil {
		return nil, err
	}
	u := &CreateUserResponse{
		ID:    res.ID,
		Name:  res.Name,
		Email: res.Email,
	}
	return u, nil
}

func (r UserRepositoryImpl) DeleteUser(ctx context.Context, id int64) error {
	err := r.querier.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

type GetUserResponse struct {
	ID    int64
	Name  string
	Email string
}

func (r UserRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*GetUserResponse, error) {
	res, err := r.querier.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	u := &GetUserResponse{
		ID:    res.ID,
		Name:  res.Name,
		Email: res.Email,
	}
	return u, nil
}

func (r UserRepositoryImpl) GetUserByID(ctx context.Context, id int64) (*GetUserResponse, error) {
	res, err := r.querier.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	u := &GetUserResponse{
		ID:    res.ID,
		Name:  res.Name,
		Email: res.Email,
	}
	return u, nil
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

func (r UserRepositoryImpl) UpdateUser(ctx context.Context, req UpdateRequest) (*UpdateResponse, error) {
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
