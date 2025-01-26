package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"readly/usecase"
)

type UserService interface {
	SignUp(ctx *gin.Context)
	SignIn(ctx *gin.Context)
}

type UserServiceImpl struct {
	signUpUseCase usecase.SignUpUseCase
	signInUseCase usecase.SignInUseCase
}

func NewUserService(signUpUseCase usecase.SignUpUseCase, signInUseCase usecase.SignInUseCase) UserServiceImpl {
	return UserServiceImpl{
		signUpUseCase: signUpUseCase,
		signInUseCase: signInUseCase,
	}
}

type SignUpRequest struct {
	Name     string `json:"name" binding:"required,min=1"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func (s UserServiceImpl) SignUp(ctx *gin.Context) {
	var req SignUpRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := usecase.SignUpRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	res, err := s.signUpUseCase.SignUp(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

type SignInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func (s UserServiceImpl) SignIn(ctx *gin.Context) {
	var req SignInRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := usecase.SignInRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	res, err := s.signInUseCase.SignIn(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}
