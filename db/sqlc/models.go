// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type ReadingStatus string

const (
	ReadingStatusUnread  ReadingStatus = "unread"
	ReadingStatusReading ReadingStatus = "reading"
	ReadingStatusDone    ReadingStatus = "done"
)

func (e *ReadingStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ReadingStatus(s)
	case string:
		*e = ReadingStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for ReadingStatus: %T", src)
	}
	return nil
}

type NullReadingStatus struct {
	ReadingStatus ReadingStatus `json:"reading_status"`
	Valid         bool          `json:"valid"` // Valid is true if ReadingStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullReadingStatus) Scan(value interface{}) error {
	if value == nil {
		ns.ReadingStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ReadingStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullReadingStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ReadingStatus), nil
}

// Stores author data.
type Author struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// Stores book data.
type Book struct {
	ID            int64          `json:"id"`
	Title         sql.NullString `json:"title"`
	Description   sql.NullString `json:"description"`
	CoverImageUrl sql.NullString `json:"cover_image_url"`
	Url           sql.NullString `json:"url"`
	AuthorName    string         `json:"author_name"`
	PublisherName string         `json:"publisher_name"`
	PublishedDate sql.NullTime   `json:"published_date"`
	Isbn          sql.NullString `json:"isbn"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

// Stores book and genre. Normalize using intermediate tables because of the many-to-many relationship between books and genres.
type BookGenre struct {
	BookID    int64  `json:"book_id"`
	GenreName string `json:"genre_name"`
}

// Stores genre data.
type Genre struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// Stores publisher data.
type Publisher struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// Stores reading history.
type ReadingHistory struct {
	UserID    int64         `json:"user_id"`
	BookID    int64         `json:"book_id"`
	Status    ReadingStatus `json:"status"`
	StartDate sql.NullTime  `json:"start_date"`
	EndDate   sql.NullTime  `json:"end_date"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

// Stores user data.
type User struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"hashed_password"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
