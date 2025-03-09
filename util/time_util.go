package util

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func ToTimeOrNil(value *timestamppb.Timestamp) *time.Time {
	if value != nil {
		t := value.AsTime()
		return &t
	}
	return nil
}

func ToTimestampOrNil(value *time.Time) *timestamppb.Timestamp {
	if value != nil {
		t := timestamppb.New(*value)
		return t
	}
	return nil
}
