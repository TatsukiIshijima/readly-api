package repository

import (
	"database/sql"
	"errors"
	"github.com/lib/pq"
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

func NewError(message string, code ErrorCode) *Error {
	return &Error{
		Message: message,
		Code:    code,
	}
}

func (err *Error) Error() string {
	return err.Message
}

func handle(err error) error {
	var code ErrorCode

	if errors.Is(err, sql.ErrNoRows) {
		code = NotFound
	}

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
	}
	if code == "" {
		return nil
	} else {
		return NewError(err.Error(), code)
	}
}
