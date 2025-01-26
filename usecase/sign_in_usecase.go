package usecase

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"readly/entity"
	"readly/repository"
)

type SignInUseCase struct {
	Executor
	userRepo repository.UserRepository
}

func NewSignInUseCase(userRepo repository.UserRepository) SignInUseCase {
	return SignInUseCase{
		userRepo: userRepo,
	}
}

type SignInRequest struct {
	Email    string
	Password string
}

func (u SignInUseCase) SignIn(ctx context.Context, req SignInRequest) (*entity.User, error) {
	user, err := u.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, handle(err)
	}
	err = checkPasswordHash(req.Password, user.Password)
	if err != nil {
		return nil, handle(err)
	}
	return &entity.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func checkPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
