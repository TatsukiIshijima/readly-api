package usecase

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"readly/entity"
	"readly/repository"
)

type SignUpUseCase struct {
	Executor
	userRepo repository.UserRepository
}

func NewSignUpUseCase(userRepo repository.UserRepository) SignUpUseCase {
	return SignUpUseCase{
		userRepo: userRepo,
	}
}

type SignUpRequest struct {
	Name     string
	Email    string
	Password string
}

func (u SignUpUseCase) SignUp(ctx context.Context, req SignUpRequest) (*entity.User, error) {
	hashedPassword, err := generateHashedPassword(req.Password)
	if err != nil {
		return nil, handle(err)
	}
	user, err := u.userRepo.CreateUser(ctx, repository.CreateUserRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	})
	if err != nil {
		return nil, handle(err)
	}
	return &entity.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func generateHashedPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
