package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomReadingHistory(t *testing.T, user User) ReadingHistory {
	book := createRandomBook(t)
	arg := CreateReadingHistoryParams{
		BookID: book.ID,
		UserID: user.ID,
		Status: ReadingStatusUnread,
		StartDate: sql.NullTime{
			Time:  time.Time{},
			Valid: false,
		},
		EndDate: sql.NullTime{
			Time:  time.Time{},
			Valid: false,
		},
	}
	readingHistory, err := testQueries.CreateReadingHistory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, readingHistory)
	require.NotZero(t, readingHistory.BookID)
	require.NotZero(t, readingHistory.UserID)
	require.Equal(t, user.ID, readingHistory.UserID)
	require.Equal(t, book.ID, readingHistory.BookID)
	require.Equal(t, arg.Status, readingHistory.Status)
	require.Equal(t, arg.StartDate, readingHistory.StartDate)
	require.Equal(t, arg.EndDate, readingHistory.EndDate)
	require.NotZero(t, readingHistory.CreatedAt)
	require.NotZero(t, readingHistory.UpdatedAt)
	return readingHistory
}

func TestCreateReadingHistory(t *testing.T) {
	user := createRandomUser(t)
	createRandomReadingHistory(t, user)
}

func TestGetReadingHistoryByUserAndBook(t *testing.T) {
	user := createRandomUser(t)
	readingHistory1 := createRandomReadingHistory(t, user)

	args := GetReadingHistoryByUserAndBookParams{
		UserID: user.ID,
		BookID: readingHistory1.BookID,
	}
	readingHistory2, err := testQueries.GetReadingHistoryByUserAndBook(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, readingHistory2)
	require.Equal(t, readingHistory1.BookID, readingHistory2.BookID)
	require.Equal(t, readingHistory1.UserID, readingHistory2.UserID)
	require.Equal(t, readingHistory1.Status, readingHistory2.Status)
	require.Equal(t, readingHistory1.StartDate, readingHistory2.StartDate)
	require.Equal(t, readingHistory1.EndDate, readingHistory2.EndDate)
	require.WithinDuration(t, readingHistory1.CreatedAt, readingHistory2.CreatedAt, time.Second)
	require.WithinDuration(t, readingHistory1.UpdatedAt, readingHistory2.UpdatedAt, time.Second)
}

func TestGetReadingHistoryByUserAndStatus(t *testing.T) {
	user := createRandomUser(t)
	_ = createRandomReadingHistory(t, user)
	_ = createRandomReadingHistory(t, user)

	args := GetReadingHistoryByUserAndStatusParams{
		UserID: user.ID,
		Status: ReadingStatusUnread,
		Limit:  5,
		Offset: 0,
	}
	readingHistories, err := testQueries.GetReadingHistoryByUserAndStatus(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, readingHistories)
	require.Len(t, readingHistories, 2)
}

func TestGetReadingHistoryByUserID(t *testing.T) {
	user := createRandomUser(t)
	_ = createRandomReadingHistory(t, user)
	_ = createRandomReadingHistory(t, user)

	args := GetReadingHistoryByUserIDParams{
		UserID: user.ID,
		Limit:  5,
		Offset: 0,
	}
	readingHistories, err := testQueries.GetReadingHistoryByUserID(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, readingHistories)
	require.Len(t, readingHistories, 2)
}

func TestUpdateReadingHistory(t *testing.T) {
	user := createRandomUser(t)
	readingHistory1 := createRandomReadingHistory(t, user)

	arg := UpdateReadingHistoryParams{
		UserID: user.ID,
		BookID: readingHistory1.BookID,
		Status: ReadingStatusReading,
		StartDate: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		EndDate: sql.NullTime{
			Time:  time.Time{},
			Valid: false,
		},
	}
	readingHistory2, err := testQueries.UpdateReadingHistory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, readingHistory2)
	require.Equal(t, readingHistory1.UserID, readingHistory2.UserID)
	require.Equal(t, readingHistory1.BookID, readingHistory2.BookID)
	require.Equal(t, arg.Status, readingHistory2.Status)
	require.Equal(t, arg.StartDate.Time.Year(), readingHistory2.StartDate.Time.Year())
	require.Equal(t, arg.StartDate.Time.Month(), readingHistory2.StartDate.Time.Month())
	require.Equal(t, arg.StartDate.Time.Day(), readingHistory2.StartDate.Time.Day())
	require.Equal(t, readingHistory1.EndDate, readingHistory2.EndDate)
	require.Equal(t, readingHistory1.CreatedAt, readingHistory2.CreatedAt)
	require.WithinDuration(t, readingHistory1.UpdatedAt, readingHistory2.UpdatedAt, time.Second)
}

func TestDeleteReadingHistory(t *testing.T) {
	user := createRandomUser(t)
	readingHistory1 := createRandomReadingHistory(t, user)

	args1 := DeleteReadingHistoryParams{
		UserID: user.ID,
		BookID: readingHistory1.BookID,
	}
	err := testQueries.DeleteReadingHistory(context.Background(), args1)
	require.NoError(t, err)

	args2 := GetReadingHistoryByUserAndBookParams{
		UserID: user.ID,
		BookID: readingHistory1.BookID,
	}

	readingHistory2, err := testQueries.GetReadingHistoryByUserAndBook(context.Background(), args2)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, readingHistory2)
}
