package repository

import sqlc "readly/db/sqlc"

type DeleteSessionByUserIDRequest struct {
	UserID int64
	Limit  int32
}

func (r DeleteSessionByUserIDRequest) toSQLC() sqlc.DeleteSessionByUserIDParams {
	return sqlc.DeleteSessionByUserIDParams{
		UserID: r.UserID,
		Limit:  r.Limit,
	}
}
