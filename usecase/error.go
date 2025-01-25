package usecase

import (
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"readly/repository"
)

type Error struct {
	Message string
	Code    ErrorCode
}

type ErrorCode string

const (
	BadRequest ErrorCode = "BAD_REQUEST"
	NotFound   ErrorCode = "NOT_FOUND"
	Forbidden  ErrorCode = "FORBIDDEN"
	Conflict   ErrorCode = "CONFLICT"
	Internal   ErrorCode = "INTERNAL"
)

func newError(message string, code ErrorCode) *Error {
	return &Error{
		Message: message,
		Code:    code,
	}
}

func (err *Error) Error() string {
	return err.Message
}

func handle(err error) error {
	if err == nil {
		return nil
	}
	var code ErrorCode

	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		switch pqErr.Code {
		case "23503":
			code = BadRequest
		case "23505":
			code = Conflict
		default:
			code = Internal
		}
		return newError(pqErr.Message, code)
	}

	if errors.Is(err, repository.ErrNoRowsDeleted) {
		return newError(err.Error(), BadRequest)
	}
	if errors.Is(err, sql.ErrNoRows) {
		return newError(err.Error(), NotFound)
	}

	return newError(err.Error(), Internal)
}
