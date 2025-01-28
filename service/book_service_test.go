package service

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"readly/entity"
	"testing"
)

func TestRegister(t *testing.T) {
	server, err := NewTestServer()
	require.NoError(t, err)
	router := server.router

	testCases := []struct {
		name string
		req  RegisterBookRequest
		code int
		res  entity.Book
	}{
		{
			name: "invalid request by missing required fields",
			req:  RegisterBookRequest{},
			code: http.StatusBadRequest,
			res:  entity.Book{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			url := "/books"
			body, err := json.Marshal(tc.req)
			println("request: " + string(body))
			require.NoError(t, err)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			router.ServeHTTP(recorder, req)

			if recorder.Code != tc.code {
				t.Fail()
			} else {
				switch recorder.Code {
				case http.StatusOK:
					var res entity.Book
					err := json.Unmarshal(recorder.Body.Bytes(), &res)
					require.NoError(t, err)
				}
			}
		})
	}
}

func TestDelete(t *testing.T) {
	server, err := NewTestServer()
	require.NoError(t, err)
	router := server.router

	testCases := []struct {
		name string
		req  DeleteBookRequest
		code int
	}{
		{
			name: "invalid request by missing required fields",
			req:  DeleteBookRequest{},
			code: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			url := "/books"
			body, err := json.Marshal(tc.req)
			println("request: " + string(body))
			require.NoError(t, err)
			req := httptest.NewRequest(http.MethodDelete, url, bytes.NewBuffer(body))
			router.ServeHTTP(recorder, req)

			if recorder.Code != tc.code {
				t.Fail()
			}
		})
	}
}
