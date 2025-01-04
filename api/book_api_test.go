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

	testCases := []struct {
		name string
		req  GetBookRequest
		code int
		exp  domain.Book
	}{
		{
			name: "invalid request",
			req:  GetBookRequest{},
			code: http.StatusBadRequest,
			exp:  domain.Book{},
		},
		{
			name: "book found",
			req:  GetBookRequest{ID: 1},
			code: http.StatusOK,
			exp: domain.Book{
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
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/books/%d", tc.req.ID)
			req, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			router.ServeHTTP(recorder, req)

			if recorder.Code != tc.code {
				t.Fail()
			} else {
				if recorder.Code == http.StatusOK {
					var act GetBookResponse
					err = json.Unmarshal(recorder.Body.Bytes(), &act)
					require.NoError(t, err)
					require.Equal(t, act.Book, tc.exp)
				}
			}
		})
	}
}
