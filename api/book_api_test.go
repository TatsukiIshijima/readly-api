package api

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"readly/domain"
	"readly/testdata"
	"testing"
)

func TestGetBook(t *testing.T) {
	router := server.router
	recorder := httptest.NewRecorder()

	args := GetBookRequest{
		ID: 1,
	}

	url := fmt.Sprintf("/books/%d", args.ID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fail()
	}

	var act GetBookResponse
	err = json.Unmarshal(recorder.Body.Bytes(), &act)
	require.NoError(t, err)

	exp := domain.Book{
		ID:            1,
		Title:         "Title",
		Genres:        []string{"Genre1", "Genre2"},
		Description:   "Description",
		CoverImageURL: "https://example.com",
		URL:           "https://example.com",
		AuthorName:    "Author",
		PublisherName: "Publisher",
		PublishDate:   testdata.TimeFrom("1970-01-01 00:00:00"),
		ISBN:          "1234567890123",
	}

	require.Equal(t, act.ID, exp.ID)
	require.Equal(t, act.Title, exp.Title)
	require.Equal(t, act.Genres, exp.Genres)
	require.Equal(t, act.Description, exp.Description)
	require.Equal(t, act.CoverImageURL, exp.CoverImageURL)
	require.Equal(t, act.URL, exp.URL)
	require.Equal(t, act.AuthorName, exp.AuthorName)
	require.Equal(t, act.PublisherName, exp.PublisherName)
	require.Equal(t, act.PublishDate, exp.PublishDate)
	require.Equal(t, act.ISBN, exp.ISBN)
}
