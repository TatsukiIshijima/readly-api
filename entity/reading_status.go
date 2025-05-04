package entity

import (
	sqlc "readly/db/sqlc"
	pb "readly/pb/readly/v1"
)

type ReadingStatus int

const (
	Unread ReadingStatus = iota
	Reading
	Done
	Unknown
)

func NewReadingStatusFromProto(proto pb.ReadingStatus) ReadingStatus {
	switch proto {
	case pb.ReadingStatus_UNREAD:
		return Unread
	case pb.ReadingStatus_READING:
		return Reading
	case pb.ReadingStatus_DONE:
		return Done
	default:
		return Unknown
	}
}

func NewReadingStatusFromSQLC(sqlcStatus sqlc.ReadingStatus) ReadingStatus {
	switch sqlcStatus {
	case sqlc.ReadingStatusUnread:
		return Unread
	case sqlc.ReadingStatusReading:
		return Reading
	case sqlc.ReadingStatusDone:
		return Done
	default:
		return Unknown
	}
}
