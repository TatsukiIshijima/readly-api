package repository

import (
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
)

func TestSaveRequestValidate(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "image_repository_test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	testCases := []struct {
		name          string
		request       SaveRequest
		expectedError string
	}{
		{
			name: "Valid request",
			request: SaveRequest{
				Dst:      tempDir,
				FileName: "test.png",
				Data:     []byte("test data"),
			},
			expectedError: "",
		},
		{
			name: "Empty destination",
			request: SaveRequest{
				Dst:      "",
				FileName: "test.png",
				Data:     []byte("test data"),
			},
			expectedError: "destination directory cannot be empty",
		},
		{
			name: "Empty file name",
			request: SaveRequest{
				Dst:      tempDir,
				FileName: "",
				Data:     []byte("test data"),
			},
			expectedError: "file name cannot be empty",
		},
		{
			name: "Empty data",
			request: SaveRequest{
				Dst:      tempDir,
				FileName: "test.png",
				Data:     []byte{},
			},
			expectedError: "file data cannot be empty",
		},
		{
			name: "Path traversal attack - absolute path",
			request: SaveRequest{
				Dst:      tempDir,
				FileName: "/etc/passwd",
				Data:     []byte("test data"),
			},
			expectedError: "file name cannot be an absolute path",
		},
		{
			name: "Path traversal attack - parent directory",
			request: SaveRequest{
				Dst:      tempDir,
				FileName: "../passwd",
				Data:     []byte("test data"),
			},
			expectedError: "path traversal attack detected",
		},
		{
			name: "Path traversal attack - multiple parent directories",
			request: SaveRequest{
				Dst:      tempDir,
				FileName: "../../etc/passwd",
				Data:     []byte("test data"),
			},
			expectedError: "path traversal attack detected",
		},
		{
			name: "Path traversal attack - normalized path",
			request: SaveRequest{
				Dst:      tempDir,
				FileName: "test/../../etc/passwd",
				Data:     []byte("test data"),
			},
			expectedError: "file name contains invalid path sequences",
		},
		{
			name: "Path traversal attack - using current directory",
			request: SaveRequest{
				Dst:      tempDir,
				FileName: "./test/../../../etc/passwd",
				Data:     []byte("test data"),
			},
			expectedError: "file name contains invalid path sequences",
		},
		{
			name: "Filename with directory component",
			request: SaveRequest{
				Dst:      tempDir,
				FileName: "subdir/test.png",
				Data:     []byte("test data"),
			},
			expectedError: "file name cannot contain directory components",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.request.Validate()
			if tc.expectedError == "" {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectedError)
			}
		})
	}
}

func TestSaveRequestValidateWithSubdirectory(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "image_repository_test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create a subdirectory
	subDir := filepath.Join(tempDir, "subdir")
	err = os.Mkdir(subDir, 0755)
	require.NoError(t, err)

	testCases := []struct {
		name          string
		request       SaveRequest
		expectedError string
	}{
		{
			name: "Filename with directory component in subdirectory test",
			request: SaveRequest{
				Dst:      tempDir,
				FileName: "subdir/test.png",
				Data:     []byte("test data"),
			},
			expectedError: "file name cannot contain directory components",
		},
		{
			name: "Path traversal attack - escaping from subdirectory",
			request: SaveRequest{
				Dst:      subDir,
				FileName: "../outside.png",
				Data:     []byte("test data"),
			},
			expectedError: "path traversal attack detected",
		},
		{
			name: "Path traversal attack - escaping from subdirectory with multiple levels",
			request: SaveRequest{
				Dst:      subDir,
				FileName: "../../outside.png",
				Data:     []byte("test data"),
			},
			expectedError: "path traversal attack detected",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.request.Validate()
			if tc.expectedError == "" {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectedError)
			}
		})
	}
}
