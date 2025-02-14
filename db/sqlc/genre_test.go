package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"readly/testdata"
	"testing"
	"time"
)

func createRandomGenre(t *testing.T) Genre {
	arg := testdata.RandomString(6)
	genre, err := querier.CreateGenre(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, genre)
	return genre
}

func TestCreateGenre(t *testing.T) {
	createRandomGenre(t)
}

func TestGetGenreByName(t *testing.T) {
	genre1 := createRandomGenre(t)
	genre2, err := querier.GetGenreByName(context.Background(), genre1.Name)
	require.NoError(t, err)
	require.NotEmpty(t, genre2)
	require.Equal(t, genre1.Name, genre2.Name)
	require.WithinDuration(t, genre1.CreatedAt, genre2.CreatedAt, time.Second)
}

func TestDeleteGenre(t *testing.T) {
	genre1 := createRandomGenre(t)
	err := querier.DeleteGenre(context.Background(), genre1.Name)
	require.NoError(t, err)

	genre2, err := querier.GetGenreByName(context.Background(), genre1.Name)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, genre2)
}

func TestGetAllGenre(t *testing.T) {
	for i := 0; i < 4; i++ {
		createRandomGenre(t)
	}

	arg := GetAllGenresParams{
		Limit:  2,
		Offset: 0,
	}

	genres, err := querier.GetAllGenres(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, genres, 2)

	for _, genre := range genres {
		require.NotEmpty(t, genre)
	}

	arg = GetAllGenresParams{
		Limit:  2,
		Offset: 2,
	}

	genres, err = querier.GetAllGenres(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, genres, 2)

	for _, genre := range genres {
		require.NotEmpty(t, genre)
	}
}
