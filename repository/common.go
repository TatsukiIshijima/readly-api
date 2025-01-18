package repository

import (
	"database/sql"
	"strings"
	"time"
)

// FIXME:移動
func newInt64(ni sql.NullInt64) *int64 {
	if !ni.Valid {
		return nil
	}
	return &ni.Int64
}

// FIXME:移動
func newString(ns sql.NullString) *string {
	if !ns.Valid {
		return nil
	}
	return &ns.String
}

// FIXME:移動
func newTime(nt sql.NullTime) *time.Time {
	if !nt.Valid {
		return nil
	}
	return &nt.Time
}

// FIXME:移動
func newGenres(bytes []byte) []string {
	return strings.Split(string(bytes), ", ")
}
