package controller

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"net/http"
	"readly/entity"
	"testing"
)

func TestRegister(t *testing.T) {
	bc, _ := setupControllers()

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
			url := "/readly/books"
			body, err := json.Marshal(tc.req)
			println("request: " + string(body))
			require.NoError(t, err)

			ctx, recorder := setupTestContext("POST", url, body)
			bc.Register(ctx)

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
	bc, _ := setupControllers()

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
			url := "/readly/books"
			body, err := json.Marshal(tc.req)
			println("request: " + string(body))
			require.NoError(t, err)

			ctx, recorder := setupTestContext("DELETE", url, body)
			bc.Delete(ctx)

			if recorder.Code != tc.code {
				t.Fail()
			}
		})
	}
}
