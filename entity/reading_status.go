package entity

import pb "readly/pb/readly/v1"

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
