package repository

import "github.com/google/uuid"

type GetSessionByIDRequest struct {
	ID uuid.UUID
}

type GetSessionByUserIDRequest struct {
	UserID int64
}
