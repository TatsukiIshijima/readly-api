package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"readly/env"
	"readly/service/auth"
	"readly/usecase"
)

type UserController interface {
	SignUp(ctx *gin.Context)
	SignIn(ctx *gin.Context)
}

type UserControllerImpl struct {
	config        env.Config
	maker         auth.TokenMaker
	signUpUseCase usecase.SignUpUseCase
	signInUseCase usecase.SignInUseCase
}

func NewUserController(config env.Config, maker auth.TokenMaker, signUpUseCase usecase.SignUpUseCase, signInUseCase usecase.SignInUseCase) UserControllerImpl {
	return UserControllerImpl{
		config:        config,
		maker:         maker,
		signUpUseCase: signUpUseCase,
		signInUseCase: signInUseCase,
	}
}

type SignUpRequest struct {
	Name     string `json:"name" binding:"required,min=1"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type SignUpResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserID       int64  `json:"user_id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
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

	result, err := s.signUpUseCase.SignUp(ctx, args)
	if err != nil {
		code, e := toHttpStatusCode(err)
		ctx.JSON(code, errorResponse(e))
		return
	}

	res := SignUpResponse{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
		UserID:       result.UserID,
		Name:         result.Name,
		Email:        result.Email,
	}

	ctx.JSON(http.StatusOK, res)
}

type SignInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type SignInResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserID       int64  `json:"user_id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
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

	result, err := s.signInUseCase.SignIn(ctx, args)
	if err != nil {
		code, e := toHttpStatusCode(err)
		ctx.JSON(code, errorResponse(e))
		return
	}

	res := SignInResponse{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
		UserID:       result.UserID,
		Name:         result.Name,
		Email:        result.Email,
	}

	ctx.JSON(http.StatusOK, res)
}
