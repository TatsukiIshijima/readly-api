package entity

import (
	"database/sql"
	"readly/pb"
	"time"
)

type Date struct {
	Year  int32 `json:"year"`
	Month int32 `json:"month"`
	Day   int32 `json:"day"`
}

func Now() Date {
	t := time.Now()
	return Date{
		Year:  int32(t.Year()),
		Month: int32(t.Month()),
		Day:   int32(t.Day()),
	}
}

func NewDateEntityFromProto(proto *pb.Date) *Date {
	if proto == nil {
		return nil
	}
	t := Date{
		Year:  proto.GetYear(),
		Month: proto.GetMonth(),
		Day:   proto.GetDay(),
	}
	return &t
}

func NewDateEntityFromNullTime(nt sql.NullTime) *Date {
	if !nt.Valid {
		return nil
	}
	return &Date{
		Year:  int32(nt.Time.Year()),
		Month: int32(nt.Time.Month()),
		Day:   int32(nt.Time.Day()),
	}
}

func (d *Date) ToProto() *pb.Date {
	return &pb.Date{
		Year:  d.Year,
		Month: d.Month,
		Day:   d.Day,
	}
}

func (d *Date) ToTime() *time.Time {
	t := time.Date(int(d.Year), time.Month(d.Month), int(d.Day), 0, 0, 0, 0, time.UTC)
	return &t
}
