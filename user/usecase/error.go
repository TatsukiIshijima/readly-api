package usecase

import (
	"errors"
	"log/slog"
)

type Error struct {
	StatusCode StatusCode
	ErrorCode  ErrorCode
	Message    string
}

type StatusCode int16

const (
	BadRequest   StatusCode = 400
	UnAuthorized StatusCode = 401
	Forbidden    StatusCode = 403
	NotFound     StatusCode = 404
	Conflict     StatusCode = 409
	Internal     StatusCode = 500
	Maintenance  StatusCode = 503
)

type ErrorCode int

const (
	// common
	InternalServerError ErrorCode = 1000
	InvalidTokenError   ErrorCode = 1001
	InvalidRequestError ErrorCode = 1002

	// user
	EmailAlreadyRegisteredError ErrorCode = 2000
	NotFoundUserError           ErrorCode = 2001
	InvalidPasswordError        ErrorCode = 2002
)

func newError(statusCode StatusCode, errorCode ErrorCode, message string) *Error {
	return &Error{
		StatusCode: statusCode,
		ErrorCode:  errorCode,
		Message:    message,
	}
}

func (err *Error) Error() string {
	return err.Message
}

func handle(err error) error {
	if err == nil {
		return nil
	}
	var e *Error
	if errors.As(err, &e) {
		return e
	} else {
		slog.Error(err.Error())
		return newError(Internal, InternalServerError, "internal server error")
	}
}
