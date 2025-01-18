package repository

import (
	"context"
	"github.com/stretchr/testify/require"
	sqlc "readly/db/sqlc"
	"testing"
)

func createReadingHistory(t *testing.T) (int64, int64, *CreateReadingHistoryResponse) {
	u := createRandomUser(t)
	b := createRandomBook(t)
	req := CreateReadingHistoryRequest{
		UserID: u.ID,
		BookID: b.ID,
		Status: Unread,
	}
	rh, err := readingHistoryRepo.Create(context.Background(), req)
	require.NoError(t, err)
	return u.ID, b.ID, rh
}

func TestCreate(t *testing.T) {
	uid, bid, rh := createReadingHistory(t)
	param := sqlc.GetReadingHistoryByUserAndBookParams{
		UserID: uid,
		BookID: bid,
	}
	_, err := querier.GetReadingHistoryByUserAndBook(context.Background(), param)
	require.NoError(t, err)

	require.Equal(t, rh.BookID, bid)
	require.Equal(t, rh.Status, Unread)
	require.Nil(t, rh.StartDate)
	require.Nil(t, rh.EndDate)
}

func TestDelete(t *testing.T) {

}

func TestGetByUser(t *testing.T) {

}

func TestGetByUserAndBook(t *testing.T) {

}

func TestGetByUserAndStatus(t *testing.T) {

}

func TestUpdate(t *testing.T) {

}
