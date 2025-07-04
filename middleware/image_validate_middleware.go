package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"path/filepath"
	"strings"
)

const (
	// Maximum file size (5MB)
	maxFileSize = 5 * 1024 * 1024

	// Supported image MIME types
	imageMimeTypeJPEG = "image/jpeg"
	imageMimeTypePNG  = "image/png"

	// Supported image extensions
	imageExtJPEG  = ".jpg"
	imageExtJPEG2 = ".jpeg"
	imageExtPNG   = ".png"

	// Context key for validated image data
	ValidatedImageDataKey = "validatedImageData"
)

// errorResponse creates a standard error response format
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// ValidateImageUpload is a middleware that validates image uploads
func ValidateImageUpload() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get file from request
		fileHeader, err := ctx.FormFile("file")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(errors.New("file is required")))
			return
		}

		// Check file size
		if fileHeader.Size > maxFileSize {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(errors.New("file size exceeds the limit of 5MB")))
			return
		}

		// Open file
		file, err := fileHeader.Open()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(errors.New("failed to open file")))
			return
		}
		defer file.Close()

		// Read file data
		data, err := io.ReadAll(file)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(errors.New("failed to read file")))
			return
		}

		// Get file extension
		ext := strings.ToLower(filepath.Ext(fileHeader.Filename))

		// Check if extension is valid
		if !isValidImageExtension(ext) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(errors.New("invalid image file extension")))
			return
		}

		// Detect MIME type from file content
		mimeType := http.DetectContentType(data)

		// Check if file is an image
		if !isImageMimeType(mimeType) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(errors.New("file is not a valid image")))
			return
		}

		// Check if extension matches content type
		if !isExtensionMatchingMimeType(ext, mimeType) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(errors.New("file extension does not match image content")))
			return
		}

		// Store validated image data in context for next handler
		ctx.Set(ValidatedImageDataKey, data)

		// Continue to next middleware/handler
		ctx.Next()
	}
}

// isValidImageExtension checks if the file extension is a valid image extension
func isValidImageExtension(ext string) bool {
	validExtensions := []string{imageExtJPEG, imageExtJPEG2, imageExtPNG}
	for _, validExt := range validExtensions {
		if ext == validExt {
			return true
		}
	}
	return false
}

// isImageMimeType checks if the MIME type is an image MIME type
func isImageMimeType(mimeType string) bool {
	validMimeTypes := []string{imageMimeTypeJPEG, imageMimeTypePNG}
	for _, validMime := range validMimeTypes {
		if strings.HasPrefix(mimeType, validMime) {
			return true
		}
	}
	return false
}

// isExtensionMatchingMimeType checks if the file extension matches the MIME type
func isExtensionMatchingMimeType(ext string, mimeType string) bool {
	switch mimeType {
	case imageMimeTypeJPEG:
		return ext == imageExtJPEG || ext == imageExtJPEG2
	case imageMimeTypePNG:
		return ext == imageExtPNG
	default:
		return false
	}
}
