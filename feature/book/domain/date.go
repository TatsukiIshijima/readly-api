package domain

import (
	"database/sql"
	"readly/pb/readly/v1"
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

// TODO: methods on both value and pointer receivers
// ポインターレシーバーではなく、値レシーバーにする
// インスタンスの値を直接変更しないため
func (d *Date) ToProto() *pb.Date {
	if d == nil {
		return nil
	}
	return &pb.Date{
		Year:  d.Year,
		Month: d.Month,
		Day:   d.Day,
	}
}

// TODO: methods on both value and pointer receivers
// ポインターレシーバーではなく、値レシーバーにする
// インスタンスの値を直接変更しないため
func (d *Date) ToTime() *time.Time {
	if d == nil {
		return nil
	}
	t := time.Date(int(d.Year), time.Month(d.Month), int(d.Day), 0, 0, 0, 0, time.UTC)
	return &t
}

func (d Date) Before(target Date) bool {
	thisTime := d.ToTime()
	targetTime := target.ToTime()

	if thisTime == nil || targetTime == nil {
		return false
	}

	return thisTime.Unix() < targetTime.Unix()
}
