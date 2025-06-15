package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"math/rand/v2"
	"readly/testdata"
	"testing"
	"time"
)

func createGenreIfNeed(t *testing.T) Genre {
	genres := testdata.GetGenres()
	i := rand.IntN(len(genres))
	genre := genres[i]

	g, err := querier.GetGenreByName(context.Background(), genre)
	if err == nil {
		return g
	}
	g, err = querier.CreateGenre(context.Background(), genre)
	require.NoError(t, err)
	require.NotEmpty(t, genre)
	return g
}

func TestGetGenreByName(t *testing.T) {
	genre1 := createGenreIfNeed(t)
	genre2, err := querier.GetGenreByName(context.Background(), genre1.Name)
	require.NoError(t, err)
	require.NotEmpty(t, genre2)
	require.Equal(t, genre1.Name, genre2.Name)
	require.WithinDuration(t, genre1.CreatedAt, genre2.CreatedAt, time.Second)
}

func TestDeleteGenre(t *testing.T) {
	genre1, err := querier.CreateGenre(context.Background(), "テストジャンル")
	require.NoError(t, err)
	err = querier.DeleteGenre(context.Background(), genre1.Name)
	require.NoError(t, err)

	genre2, err := querier.GetGenreByName(context.Background(), genre1.Name)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, genre2)
}

func TestGetAllGenre(t *testing.T) {
	genres, err := querier.GetAllGenres(context.Background())
	require.NoError(t, err)

	for i := 0; i < 4; i++ {
		createGenreIfNeed(t)
	}

	allGenres, err := querier.GetAllGenres(context.Background())
	require.NoError(t, err)
	require.Len(t, allGenres, len(genres))

	for _, genre := range genres {
		require.NotEmpty(t, genre)
	}
}
