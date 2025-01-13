package repository

import (
	"context"
	sqlc "readly/db/sqlc"
	"readly/entity"
)

type Repository interface {
	CreateUser(ctx context.Context, req CreateUserRequest) (*entity.User, error)
	DeleteUser(ctx context.Context, id int64) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserByID(ctx context.Context, id int64) (*entity.User, error)
	UpdateUser(ctx context.Context, req UpdateRequest) (*entity.User, error)
}

type RepositoryImpl struct {
	container *sqlc.Container
}

func NewUserRepository(container *sqlc.Container) RepositoryImpl {
	return RepositoryImpl{
		container: container,
	}
}

type CreateUserRequest struct {
	Name     string
	Email    string
	Password string
}

func (r RepositoryImpl) CreateUser(ctx context.Context, req CreateUserRequest) (*entity.User, error) {
	args := sqlc.CreateUserParams{
		Name:           req.Name,
		Email:          req.Email,
		HashedPassword: req.Password,
	}
	res, err := r.container.Querier.CreateUser(ctx, args)
	if err != nil {
		return nil, err
	}
	u := &entity.User{
		ID:    res.ID,
		Name:  res.Name,
		Email: res.Email,
	}
	return u, nil
}

func (r RepositoryImpl) DeleteUser(ctx context.Context, id int64) error {
	err := r.container.Querier.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (r RepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	res, err := r.container.Querier.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	u := &entity.User{
		ID:    res.ID,
		Name:  res.Name,
		Email: res.Email,
	}
	return u, nil
}

func (r RepositoryImpl) GetUserByID(ctx context.Context, id int64) (*entity.User, error) {
	res, err := r.container.Querier.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	u := &entity.User{
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

func (r RepositoryImpl) UpdateUser(ctx context.Context, req UpdateRequest) (*entity.User, error) {
	args := sqlc.UpdateUserParams{
		ID:             req.ID,
		Name:           req.Name,
		Email:          req.Email,
		HashedPassword: req.Password,
	}
	res, err := r.container.Querier.UpdateUser(ctx, args)
	if err != nil {
		return nil, err
	}
	u := &entity.User{
		ID:    res.ID,
		Name:  res.Name,
		Email: res.Email,
	}
	return u, nil
}
