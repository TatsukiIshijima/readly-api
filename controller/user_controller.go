package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"readly/usecase"
)

type UserController interface {
	SignUp(ctx *gin.Context)
	SignIn(ctx *gin.Context)
}

type UserControllerImpl struct {
	signUpUseCase usecase.SignUpUseCase
	signInUseCase usecase.SignInUseCase
}

func NewUserController(signUpUseCase usecase.SignUpUseCase, signInUseCase usecase.SignInUseCase) UserControllerImpl {
	return UserControllerImpl{
		signUpUseCase: signUpUseCase,
		signInUseCase: signInUseCase,
	}
}

type SignUpRequest struct {
	Name     string `json:"name" binding:"required,min=1"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func (s UserControllerImpl) SignUp(ctx *gin.Context) {
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

func (s UserControllerImpl) SignIn(ctx *gin.Context) {
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
