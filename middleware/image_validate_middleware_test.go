package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"image"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

// createTestFile creates a file with the specified size and content type
func createTestFile(t *testing.T, size int64, contentType, path string) {
	f, err := os.Create(path)
	require.NoError(t, err)
	defer f.Close()

	// Write header based on content type
	if contentType == "image/png" {
		// Create a small valid PNG header
		img := image.NewRGBA(image.Rect(0, 0, 1, 1))
		err = png.Encode(f, img)
		require.NoError(t, err)
	} else if contentType == "text/plain" {
		_, err = f.WriteString("This is a text file\n")
		require.NoError(t, err)
	}

	// Get current file size
	info, err := f.Stat()
	require.NoError(t, err)
	currentSize := info.Size()

	// Fill the rest with zeros to reach the desired size
	if currentSize < size {
		zeros := make([]byte, size-currentSize)
		_, err = f.Write(zeros)
		require.NoError(t, err)
	}
}

// setupTestFiles creates all the test files needed for the tests
func setupTestFiles(t *testing.T) {
	// Create test files directory if it doesn't exist
	testDir := filepath.Join("testdata")
	if _, err := os.Stat(testDir); os.IsNotExist(err) {
		err = os.MkdirAll(testDir, 0755)
		require.NoError(t, err)
	}

	// 5MB PNG file with .png extension
	createTestFile(t, 5*1024*1024, "image/png", filepath.Join(testDir, "5mb_valid.png"))

	// 6MB PNG file with .png extension
	createTestFile(t, 6*1024*1024, "image/png", filepath.Join(testDir, "6mb_invalid.png"))

	// 5MB PNG file with .jpg extension
	createTestFile(t, 5*1024*1024, "image/png", filepath.Join(testDir, "5mb_mismatch.jpg"))

	// Text file
	createTestFile(t, 1024, "text/plain", filepath.Join(testDir, "invalid.txt"))
}

// cleanupTestFiles removes all the test files created for the tests
func cleanupTestFiles(t *testing.T) {
	files := []string{
		filepath.Join("testdata", "5mb_valid.png"),
		filepath.Join("testdata", "6mb_invalid.png"),
		filepath.Join("testdata", "5mb_mismatch.jpg"),
		filepath.Join("testdata", "invalid.txt"),
	}

	for _, file := range files {
		err := os.Remove(file)
		if err != nil && !os.IsNotExist(err) {
			t.Logf("Failed to remove test file %s: %v", file, err)
		}
	}
}

// createRequestWithFile creates an HTTP request with a file upload
func createRequestWithFile(t *testing.T, method, url, fieldName, filePath string) (*http.Request, *httptest.ResponseRecorder) {
	// Create a buffer to store the request body
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Open the file
	file, err := os.Open(filePath)
	require.NoError(t, err)
	defer file.Close()

	// Create a form file field
	part, err := writer.CreateFormFile(fieldName, filepath.Base(filePath))
	require.NoError(t, err)

	// Copy the file content to the form field
	_, err = io.Copy(part, file)
	require.NoError(t, err)

	// Close the writer
	err = writer.Close()
	require.NoError(t, err)

	// Create the request
	req, err := http.NewRequest(method, url, body)
	require.NoError(t, err)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Create a response recorder
	recorder := httptest.NewRecorder()

	return req, recorder
}

func TestValidateImageUpload(t *testing.T) {
	// Set up test files
	setupTestFiles(t)
	defer cleanupTestFiles(t)

	// Set up test cases
	testCases := []struct {
		name     string
		filePath string
		wantCode int
	}{
		{
			name:     "Valid 5MB PNG file",
			filePath: filepath.Join("testdata", "5mb_valid.png"),
			wantCode: http.StatusOK,
		},
		{
			name:     "Invalid 6MB PNG file",
			filePath: filepath.Join("testdata", "6mb_invalid.png"),
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "Invalid PNG file with .jpg extension",
			filePath: filepath.Join("testdata", "5mb_mismatch.jpg"),
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "Invalid .txt file",
			filePath: filepath.Join("testdata", "invalid.txt"),
			wantCode: http.StatusBadRequest,
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Set up Gin router
			gin.SetMode(gin.TestMode)
			router := gin.Default()

			// Add middleware and test handler
			router.POST("/test", ValidateImageUpload(), func(ctx *gin.Context) {
				ctx.Status(http.StatusOK)
			})

			// Create request with file
			req, recorder := createRequestWithFile(t, http.MethodPost, "/test", "file", tc.filePath)

			// Serve the request
			router.ServeHTTP(recorder, req)

			// Check response
			require.Equal(t, tc.wantCode, recorder.Code)
		})
	}
}
