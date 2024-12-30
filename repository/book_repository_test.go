package repository

import (
	"context"
	"github.com/stretchr/testify/require"
	"readly/test"
	"testing"
	"time"
)

func TestRegisterBook(t *testing.T) {
	store := NewStore(test.DB)
	repo := NewBookRepository(store)

	var err error
	user, err := test.CreateRandomUser()
	require.NoError(t, err)

	args := RegisterBookParams{
		UserID:        user.ID,
		Title:         test.RandomString(6),
		Genres:        []string{"fiction"},
		Description:   test.RandomString(12),
		CoverImageURL: "https://example.com",
		URL:           "https://example.com",
		AuthorName:    test.RandomString(6),
		PublisherName: test.RandomString(6),
		PublishDate:   time.Now(),
		ISBN:          test.RandomString(13),
	}

	err = repo.Register(context.Background(), args)
	require.NoError(t, err)
}
