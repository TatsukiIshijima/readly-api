package repository

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	sqlc "readly/db/sqlc"
	"testing"
	"time"
)

func createReadingHistory(t *testing.T, uid int64, bid int64) *CreateReadingHistoryResponse {
	req := CreateReadingHistoryRequest{
		UserID: uid,
		BookID: bid,
		Status: Unread,
	}
	rh, err := readingHistoryRepo.Create(context.Background(), req)
	require.NoError(t, err)
	return rh
}

func TestCreate(t *testing.T) {
	u := createRandomUser(t)
	b := createRandomBook(t)
	rh := createReadingHistory(t, u.ID, b.ID)
	param := sqlc.GetReadingHistoryByUserAndBookParams{
		UserID: u.ID,
		BookID: b.ID,
	}
	_, err := querier.GetReadingHistoryByUserAndBook(context.Background(), param)
	require.NoError(t, err)

	require.Equal(t, rh.BookID, b.ID)
	require.Equal(t, rh.Status, Unread)
	require.Nil(t, rh.StartDate)
	require.Nil(t, rh.EndDate)
}

func TestDelete(t *testing.T) {
	u := createRandomUser(t)
	b := createRandomBook(t)
	_ = createReadingHistory(t, u.ID, b.ID)
	req := DeleteReadingHistoryRequest{
		UserID: u.ID,
		BookID: b.ID,
	}
	err := readingHistoryRepo.Delete(context.Background(), req)
	require.NoError(t, err)

	param := sqlc.GetReadingHistoryByUserAndBookParams{
		UserID: u.ID,
		BookID: b.ID,
	}
	_, err = querier.GetReadingHistoryByUserAndBook(context.Background(), param)
	require.ErrorIs(t, err, sql.ErrNoRows)
}

func TestGetByUser(t *testing.T) {
	u := createRandomUser(t)
	b1 := createRandomBook(t)
	b2 := createRandomBook(t)
	rh1 := createReadingHistory(t, u.ID, b1.ID)
	rh2 := createReadingHistory(t, u.ID, b2.ID)
	crhs := []CreateReadingHistoryResponse{*rh1, *rh2}

	req := GetReadingHistoryByUserRequest{
		UserID: u.ID,
		Limit:  5,
		Offset: 0,
	}
	grhs, err := readingHistoryRepo.GetByUser(context.Background(), req)
	require.NoError(t, err)
	require.Len(t, grhs, 2)
	for i, rh := range grhs {
		require.Equal(t, crhs[i].BookID, rh.BookID)
		require.Equal(t, crhs[i].Status, rh.Status)
		require.Nil(t, rh.StartDate)
		require.Nil(t, rh.EndDate)
	}
}

func TestGetByUserAndBook(t *testing.T) {
	u := createRandomUser(t)
	b := createRandomBook(t)
	rh := createReadingHistory(t, u.ID, b.ID)

	req := GetReadingHistoryByUserAndBookRequest{
		UserID: u.ID,
		BookID: b.ID,
	}
	grh, err := readingHistoryRepo.GetByUserAndBook(context.Background(), req)
	require.NoError(t, err)
	require.Equal(t, rh.BookID, grh.BookID)
	require.Equal(t, rh.Status, grh.Status)
	require.Nil(t, grh.StartDate)
	require.Nil(t, grh.EndDate)
}

func TestGetByUserAndStatus(t *testing.T) {
	u := createRandomUser(t)
	b1 := createRandomBook(t)
	b2 := createRandomBook(t)
	rh1 := createReadingHistory(t, u.ID, b1.ID)
	_ = createReadingHistory(t, u.ID, b2.ID)

	now := time.Now().UTC()
	updateReq := UpdateReadingHistoryRequest{
		UserID:    u.ID,
		BookID:    b1.ID,
		Status:    Reading,
		StartDate: &now,
		EndDate:   nil,
	}
	_, err := readingHistoryRepo.Update(context.Background(), updateReq)

	getReq := GetReadingHistoryByUserAndStatusRequest{
		UserID: u.ID,
		Status: Reading,
		Limit:  5,
		Offset: 0,
	}
	grhs, err := readingHistoryRepo.GetByUserAndStatus(context.Background(), getReq)
	require.NoError(t, err)
	require.Len(t, grhs, 1)
	require.Equal(t, rh1.BookID, grhs[0].BookID)
	require.Equal(t, Reading, grhs[0].Status)
	require.Equal(t, now.Year(), grhs[0].StartDate.Year())
	require.Equal(t, now.Month(), grhs[0].StartDate.Month())
	require.Equal(t, now.Day(), grhs[0].StartDate.Day())
	require.Nil(t, grhs[0].EndDate)
}

func TestUpdate(t *testing.T) {
	u := createRandomUser(t)
	b := createRandomBook(t)
	rh := createReadingHistory(t, u.ID, b.ID)

	now := time.Now().UTC()
	updateReq := UpdateReadingHistoryRequest{
		UserID:    u.ID,
		BookID:    b.ID,
		Status:    Reading,
		StartDate: &now,
		EndDate:   nil,
	}
	urh, err := readingHistoryRepo.Update(context.Background(), updateReq)
	require.NoError(t, err)

	require.Equal(t, rh.BookID, urh.BookID)
	require.Equal(t, Reading, urh.Status)
	require.Equal(t, now.Year(), urh.StartDate.Year())
	require.Equal(t, now.Month(), urh.StartDate.Month())
	require.Equal(t, now.Day(), urh.StartDate.Day())
	require.Nil(t, urh.EndDate)
}
