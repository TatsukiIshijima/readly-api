// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateAuthor(ctx context.Context, name string) (Author, error)
	CreateBook(ctx context.Context, arg CreateBookParams) (Book, error)
	CreateBookGenre(ctx context.Context, arg CreateBookGenreParams) (BookGenre, error)
	CreateGenre(ctx context.Context, name string) (Genre, error)
	CreatePublisher(ctx context.Context, name string) (Publisher, error)
	CreateReadingHistory(ctx context.Context, arg CreateReadingHistoryParams) (ReadingHistory, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteAuthor(ctx context.Context, name string) error
	DeleteBook(ctx context.Context, id int64) error
	DeleteBookGenre(ctx context.Context, arg DeleteBookGenreParams) error
	DeleteGenre(ctx context.Context, name string) error
	DeletePublisher(ctx context.Context, name string) error
	DeleteReadingHistory(ctx context.Context, arg DeleteReadingHistoryParams) error
	DeleteUser(ctx context.Context, id int64) error
	GetAllAuthors(ctx context.Context, arg GetAllAuthorsParams) ([]Author, error)
	GetAllGenres(ctx context.Context, arg GetAllGenresParams) ([]Genre, error)
	GetAllPublishers(ctx context.Context, arg GetAllPublishersParams) ([]Publisher, error)
	GetAllUsers(ctx context.Context, arg GetAllUsersParams) ([]User, error)
	GetAuthorByName(ctx context.Context, name string) (Author, error)
	GetBooksByAuthor(ctx context.Context, authorName string) ([]GetBooksByAuthorRow, error)
	GetBooksByID(ctx context.Context, id int64) (GetBooksByIDRow, error)
	GetBooksByISBN(ctx context.Context, isbn sql.NullString) ([]GetBooksByISBNRow, error)
	GetBooksByPublisher(ctx context.Context, publisherName string) ([]GetBooksByPublisherRow, error)
	GetBooksByTitle(ctx context.Context, title sql.NullString) ([]GetBooksByTitleRow, error)
	GetGenreByName(ctx context.Context, name string) (Genre, error)
	GetGenresByBookID(ctx context.Context, bookID int64) ([]string, error)
	GetPublisherByName(ctx context.Context, name string) (Publisher, error)
	GetReadingHistoryByUserAndBook(ctx context.Context, arg GetReadingHistoryByUserAndBookParams) (GetReadingHistoryByUserAndBookRow, error)
	GetReadingHistoryByUserAndStatus(ctx context.Context, arg GetReadingHistoryByUserAndStatusParams) ([]GetReadingHistoryByUserAndStatusRow, error)
	GetReadingHistoryByUserID(ctx context.Context, arg GetReadingHistoryByUserIDParams) ([]GetReadingHistoryByUserIDRow, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByID(ctx context.Context, id int64) (User, error)
	UpdateBook(ctx context.Context, arg UpdateBookParams) (Book, error)
	UpdateReadingHistory(ctx context.Context, arg UpdateReadingHistoryParams) (ReadingHistory, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
