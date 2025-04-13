package usecase

import (
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"readly/env"
	"testing"
)

func TestUploadImage(t *testing.T) {
	uploadImgUseCase := newTestUploadImgUseCase(t)

	testCases := []struct {
		name  string
		setup func(t *testing.T) UploadRequest
		check func(t *testing.T, res *UploadImgResponse, err error)
	}{
		{
			name: "Upload image success",
			setup: func(t *testing.T) UploadRequest {
				inputImg := filepath.Join(env.ProjectRoot(), "testdata/sample_150.png")
				data, err := os.ReadFile(inputImg)
				require.NoError(t, err)
				return UploadRequest{
					Data: data,
					Ext:  filepath.Ext(inputImg),
				}
			},
			check: func(t *testing.T, res *UploadImgResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.Equal(t, filepath.Join(env.ProjectRoot(), ".storage/cover_img"), filepath.Dir(res.Path))
				require.Equal(t, ".png", filepath.Ext(res.Path))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := tc.setup(t)
			res, err := uploadImgUseCase.Upload(req)
			tc.check(t, res, err)

			// Clean up the uploaded image
			if res != nil {
				err := os.Remove(res.Path)
				require.NoError(t, err)
			}
		})
	}
}
