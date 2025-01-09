package user

import (
	"context"
	sqlc "readly/db/sqlc"
	"readly/domain"
)

type Repository interface {
	Register(ctx context.Context, req RegisterRequest) (*domain.User, error)
	Login(ctx context.Context, req LoginRequest) (domain.User, error)
	GetByID(ctx context.Context, id int64) (domain.User, error)
	GetByEmail(ctx context.Context, email string) (domain.User, error)
	Update(ctx context.Context, req UpdateRequest) (domain.User, error)
	Delete(ctx context.Context, id int64) error
}

type RepositoryImpl struct {
	container sqlc.Container
}

func NewRepository(db sqlc.DBTX, q sqlc.Querier) RepositoryImpl {
	return RepositoryImpl{
		container: sqlc.NewContainer(db, q),
	}
}

type RegisterRequest struct {
	Name     string
	Email    string
	Password string
}

type LoginRequest struct {
	Email    string
	Password string
}

type UpdateRequest struct {
	ID       int64
	Name     string
	Email    string
	Password string
}

func (r RepositoryImpl) Register(ctx context.Context, req RegisterRequest) (*domain.User, error) {
	args := sqlc.CreateUserParams{
		Name:           req.Name,
		Email:          req.Email,
		HashedPassword: req.Password,
	}
	user, err := r.container.Querier.CreateUser(ctx, args)
	if err != nil {
		return nil, err
	}
	res := &domain.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	return res, nil
}

func (r RepositoryImpl) Login(ctx context.Context, req LoginRequest) (domain.User, error) {
	// TODO: Implement
	return domain.User{}, nil
}

func (r RepositoryImpl) GetByID(ctx context.Context, id int64) (domain.User, error) {
	// TODO: Implement
	return domain.User{}, nil
}

func (r RepositoryImpl) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	// TODO: Implement
	return domain.User{}, nil
}

func (r RepositoryImpl) Update(ctx context.Context, req UpdateRequest) (domain.User, error) {
	// TODO: Implement
	return domain.User{}, nil
}

func (r RepositoryImpl) Delete(ctx context.Context, id int64) error {
	// TODO: Implement
	return nil
}
