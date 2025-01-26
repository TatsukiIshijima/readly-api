package entity

type ReadingStatus int

const (
	Unread ReadingStatus = iota
	Reading
	Done
	Unknown
)
