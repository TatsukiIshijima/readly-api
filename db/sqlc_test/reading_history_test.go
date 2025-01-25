package sqlc_test

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"readly/db/sqlc"
	"readly/testdata"
	"sort"
	"strings"
	"testing"
	"time"
)

func createRandomReadingHistory(t *testing.T, user db.User, genresLen int, status db.ReadingStatus) (db.Book, []db.Genre, db.ReadingHistory) {
	b := createBook(
		t,
		testdata.RandomString(6),
		"",
		"",
		testdata.RandomString(13),
	)
	genres := make([]db.Genre, genresLen)
	for i := 0; i < genresLen; i++ {
		g := createRandomGenre(t)
		createRandomBookGenre(t, b, g)
		genres[i] = g
	}
	y, m, d := time.Now().Date()
	startDate := sql.NullTime{
		Time:  time.Time{},
		Valid: false,
	}
	endDate := sql.NullTime{
		Time:  time.Time{},
		Valid: false,
	}
	switch status {
	case db.ReadingStatusUnread:
	case db.ReadingStatusReading:
		startDate = sql.NullTime{
			Time:  time.Date(y, m, d, 0, 0, 0, 0, time.UTC),
			Valid: true,
		}
	case db.ReadingStatusDone:
		endDate = sql.NullTime{
			Time:  time.Date(y, m, d, 0, 0, 0, 0, time.UTC),
			Valid: true,
		}
	default:
		t.Fatalf("invalid status: %v", status)
	}
	arg := db.CreateReadingHistoryParams{
		BookID:    b.ID,
		UserID:    user.ID,
		Status:    status,
		StartDate: startDate,
		EndDate:   endDate,
	}
	rh, err := querier.CreateReadingHistory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, rh)
	require.NotZero(t, rh.BookID)
	require.NotZero(t, rh.UserID)
	require.Equal(t, user.ID, rh.UserID)
	require.Equal(t, b.ID, rh.BookID)
	require.Equal(t, arg.Status, rh.Status)
	compareNullTimes(t, arg.StartDate, rh.StartDate)
	compareNullTimes(t, arg.EndDate, rh.EndDate)
	require.NotZero(t, rh.CreatedAt)
	require.NotZero(t, rh.UpdatedAt)
	return b, genres, rh
}

func compareNullTimes(t *testing.T, t1 sql.NullTime, t2 sql.NullTime) {
	if t1.Valid != t2.Valid {
		require.Fail(t, "Valid field is different")
	}
	if t1.Valid == false && t2.Valid == false {
		require.Equal(t, t1.Valid, t2.Valid)
	} else if t1.Valid == true && t2.Valid == true {
		lt1 := t1.Time.UTC()
		lt2 := t2.Time.UTC()
		require.Equal(t, lt1, lt2)
	}
}

func TestCreateReadingHistory(t *testing.T) {
	user := createRandomUser(t)
	createRandomReadingHistory(t, user, 0, db.ReadingStatusUnread)
}

func TestGetReadingHistoryByUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)
	b1, _, rh1 := createRandomReadingHistory(t, user1, 0, db.ReadingStatusUnread)
	b2, g2, rh2 := createRandomReadingHistory(t, user1, 2, db.ReadingStatusUnread)

	args1 := db.GetReadingHistoryByUserParams{
		UserID: user1.ID,
		Limit:  5,
		Offset: 0,
	}
	args2 := db.GetReadingHistoryByUserParams{
		UserID: user2.ID,
		Limit:  5,
		Offset: 0,
	}
	result1, err := querier.GetReadingHistoryByUser(context.Background(), args1)
	require.NoError(t, err)
	require.NotEmpty(t, result1)
	require.Len(t, result1, 2)
	result2, err := querier.GetReadingHistoryByUser(context.Background(), args2)
	require.NoError(t, err)
	require.Empty(t, result2)
	require.Len(t, result2, 0)

	require.Equal(t, result1[0].ID, sql.NullInt64{Int64: b1.ID, Valid: true})
	require.Equal(t, result1[0].Title.String, b1.Title)
	require.Empty(t, result1[0].Genres)
	require.Equal(t, result1[0].Description, b1.Description)
	require.Equal(t, result1[0].CoverImageUrl, b1.CoverImageUrl)
	require.Equal(t, result1[0].Url, b1.Url)
	require.Equal(t, result1[0].AuthorName, b1.AuthorName)
	require.Equal(t, result1[0].PublisherName, b1.PublisherName)
	require.Equal(t, result1[0].PublishedDate, b1.PublishedDate)
	require.Equal(t, result1[0].Isbn, b1.Isbn)
	require.Equal(t, result1[0].Status, rh1.Status)
	require.Equal(t, result1[0].StartDate, rh1.StartDate)
	require.Equal(t, result1[0].EndDate, rh1.EndDate)

	require.Equal(t, result1[1].ID, sql.NullInt64{Int64: b2.ID, Valid: true})
	require.Equal(t, result1[1].Title.String, b2.Title)
	g2Names := make([]string, len(g2))
	for i := 0; i < len(g2); i++ {
		g2Names[i] = g2[i].Name
	}
	sort.Strings(g2Names)
	require.Equal(t, g2Names, strings.Split(string(result1[1].Genres), ", "))
	require.Equal(t, result1[1].Description, b2.Description)
	require.Equal(t, result1[1].CoverImageUrl, b2.CoverImageUrl)
	require.Equal(t, result1[1].Url, b2.Url)
	require.Equal(t, result1[1].AuthorName, b2.AuthorName)
	require.Equal(t, result1[1].PublisherName, b2.PublisherName)
	require.Equal(t, result1[1].PublishedDate, b2.PublishedDate)
	require.Equal(t, result1[1].Isbn, b2.Isbn)
	require.Equal(t, result1[1].Status, rh2.Status)
	require.Equal(t, result1[1].StartDate, rh2.StartDate)
	require.Equal(t, result1[1].EndDate, rh2.EndDate)
}

func TestGetReadingHistoryByUserAndStatus(t *testing.T) {
	user := createRandomUser(t)
	_, _, _ = createRandomReadingHistory(t, user, 0, db.ReadingStatusUnread)
	b2, _, rh2 := createRandomReadingHistory(t, user, 0, db.ReadingStatusReading)
	b3, g3, rh3 := createRandomReadingHistory(t, user, 1, db.ReadingStatusReading)

	args := db.GetReadingHistoryByUserAndStatusParams{
		UserID: user.ID,
		Status: db.ReadingStatusReading,
		Limit:  5,
		Offset: 0,
	}
	result, err := querier.GetReadingHistoryByUserAndStatus(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Len(t, result, 2)
	require.Equal(t, result[0].ID, sql.NullInt64{Int64: b2.ID, Valid: true})
	require.Equal(t, result[0].Title.String, b2.Title)
	require.Empty(t, result[0].Genres)
	require.Equal(t, result[0].Description, b2.Description)
	require.Equal(t, result[0].CoverImageUrl, b2.CoverImageUrl)
	require.Equal(t, result[0].Url, b2.Url)
	require.Equal(t, result[0].AuthorName, b2.AuthorName)
	require.Equal(t, result[0].PublisherName, b2.PublisherName)
	require.Equal(t, result[0].PublishedDate, b2.PublishedDate)
	require.Equal(t, result[0].Isbn, b2.Isbn)
	require.Equal(t, result[0].Status, db.ReadingStatusReading)
	require.Equal(t, result[0].StartDate, rh2.StartDate)
	require.Equal(t, result[0].EndDate, rh2.EndDate)

	require.Equal(t, result[1].ID, sql.NullInt64{Int64: b3.ID, Valid: true})
	require.Equal(t, result[1].Title.String, b3.Title)
	gNames := make([]string, len(g3))
	for i := 0; i < len(g3); i++ {
		gNames[i] = g3[i].Name
	}
	sort.Strings(gNames)
	require.Equal(t, gNames, strings.Split(string(result[1].Genres), ", "))
	require.Equal(t, result[1].Description, b3.Description)
	require.Equal(t, result[1].CoverImageUrl, b3.CoverImageUrl)
	require.Equal(t, result[1].Url, b3.Url)
	require.Equal(t, result[1].AuthorName, b3.AuthorName)
	require.Equal(t, result[1].PublisherName, b3.PublisherName)
	require.Equal(t, result[1].PublishedDate, b3.PublishedDate)
	require.Equal(t, result[1].Isbn, b3.Isbn)
	require.Equal(t, result[1].Status, db.ReadingStatusReading)
	require.Equal(t, result[1].StartDate, rh3.StartDate)
	require.Equal(t, result[1].EndDate, rh3.EndDate)
}

func TestUpdateReadingHistory(t *testing.T) {
	user := createRandomUser(t)
	_, _, readingHistory1 := createRandomReadingHistory(t, user, 0, db.ReadingStatusUnread)

	arg := db.UpdateReadingHistoryParams{
		UserID: user.ID,
		BookID: readingHistory1.BookID,
		Status: db.ReadingStatusReading,
		StartDate: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		EndDate: sql.NullTime{
			Time:  time.Time{},
			Valid: false,
		},
	}
	readingHistory2, err := querier.UpdateReadingHistory(context.Background(), arg)
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
	_, _, readingHistory1 := createRandomReadingHistory(t, user, 0, db.ReadingStatusUnread)

	args1 := db.DeleteReadingHistoryParams{
		UserID: user.ID,
		BookID: readingHistory1.BookID,
	}
	rowsAffected, err := querier.DeleteReadingHistory(context.Background(), args1)
	require.NoError(t, err)
	require.Equal(t, int64(1), rowsAffected)

	args2 := db.GetReadingHistoryByUserAndBookParams{
		UserID: user.ID,
		BookID: readingHistory1.BookID,
	}

	readingHistory2, err := querier.GetReadingHistoryByUserAndBook(context.Background(), args2)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, readingHistory2)
}
