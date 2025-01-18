package repository

import (
	"database/sql"
	"strings"
	"time"
)

func nilInt64(ni sql.NullInt64) *int64 {
	if !ni.Valid {
		return nil
	}
	return &ni.Int64
}

func nilString(ns sql.NullString) *string {
	if !ns.Valid {
		return nil
	}
	return &ns.String
}

func nilTime(nt sql.NullTime) *time.Time {
	if !nt.Valid {
		return nil
	}
	return &nt.Time
}

func newGenres(bytes []byte) []string {
	return strings.Split(string(bytes), ", ")
}
