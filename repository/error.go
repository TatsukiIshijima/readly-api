package repository

import (
	"errors"
)

var ErrNoRows = errors.New("no rows in result set")
var ErrNoRowsDeleted = errors.New("no rows were deleted")
