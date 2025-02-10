package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"readly/usecase"
)

type ErrorResponse struct {
	Code    int    `json:"error_code"`
	Message string `json:"error_message"`
}

func errorResponse(err error) gin.H {
	var e *usecase.Error
	if !errors.As(err, &e) {
		return gin.H{"code": -1, "message": err.Error()}
	}
	return gin.H{"code": e.ErrorCode, "message": e.Message}
}

func toHttpStatusCode(err error) (int, error) {
	var e *usecase.Error
	if !errors.As(err, &e) {
		return http.StatusInternalServerError, err
	}
	var sc int
	switch e.StatusCode {
	case usecase.BadRequest:
		sc = http.StatusBadRequest
	case usecase.UnAuthorized:
		sc = http.StatusUnauthorized
	case usecase.Forbidden:
		sc = http.StatusForbidden
	case usecase.NotFound:
		sc = http.StatusNotFound
	case usecase.Conflict:
		sc = http.StatusConflict
	case usecase.Internal:
		sc = http.StatusInternalServerError
	default:
		sc = http.StatusInternalServerError
	}
	return sc, e
}
