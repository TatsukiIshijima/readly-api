package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"readly/domain"
	"readly/testdata"
	"strings"
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

func TestRegisterBook(t *testing.T) {
	router := server.router
	url := "/books"

	testCases := []struct {
		name string
		req  RegisterRequest
		code int
	}{
		{
			name: "empty request",
			req:  RegisterRequest{},
			code: http.StatusBadRequest,
		},
		{
			name: "invalid request (title is empty)",
			req: RegisterRequest{
				UserID: 1,
				Title:  "",
			},
			code: http.StatusBadRequest,
		},
		{
			name: "invalid request (too many genres)",
			req: RegisterRequest{
				UserID: 1,
				Title:  "Title",
				Genres: []string{"genre1", "genre2", "genre3", "genre4", "genre5", "genre6"},
			},
			code: http.StatusBadRequest,
		},
		{
			name: "invalid request (description is too long)",
			req: RegisterRequest{
				UserID:      1,
				Title:       "Title",
				Description: strings.Repeat("a", 201),
			},
			code: http.StatusBadRequest,
		},
		{
			name: "invalid request (cover image url is invalid)",
			req: RegisterRequest{
				UserID:        1,
				Title:         "Title",
				CoverImageURL: "invalid",
			},
			code: http.StatusBadRequest,
		},
		{
			name: "invalid request (url is invalid)",
			req: RegisterRequest{
				UserID: 1,
				Title:  "Title",
				URL:    "invalid",
			},
			code: http.StatusBadRequest,
		},
		{
			name: "invalid request (author name is too long)",
			req: RegisterRequest{
				UserID:     1,
				Title:      "Title",
				AuthorName: strings.Repeat("a", 51),
			},
			code: http.StatusBadRequest,
		},
		{
			name: "invalid request (publisher name is too long)",
			req: RegisterRequest{
				UserID:        1,
				Title:         "Title",
				PublisherName: strings.Repeat("a", 51),
			},
			code: http.StatusBadRequest,
		},
		{
			name: "invalid request (isbn is invalid)",
			req: RegisterRequest{
				UserID: 1,
				Title:  "Title",
				ISBN:   "invalid",
			},
			code: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()

			rb, err := json.Marshal(tc.req)
			require.NoError(t, err)
			req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(rb))
			require.NoError(t, err)

			router.ServeHTTP(recorder, req)

			if recorder.Code != tc.code {
				t.Fail()
			}
		})
	}
}
