package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"readly/middleware/image"
	"readly/pb/readly/v1"
	"readly/usecase"
)

type ImageServer interface {
	Upload(ctx *gin.Context)
}

type ImageServerImpl struct {
	uploadImgUseCase usecase.UploadImgUseCase
}

func NewImageServer(uploadImgUseCase usecase.UploadImgUseCase) ImageServer {
	return &ImageServerImpl{
		uploadImgUseCase: uploadImgUseCase,
	}
}

func (s *ImageServerImpl) Upload(ctx *gin.Context) {
	// TODO: エラーハンドルをcontroller/error.goに合わせる
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	// Get validated image data from context
	validatedData, exists := ctx.Get(image.ValidatedImageDataKey)
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Validated image data not found in context"})
		return
	}

	data, ok := validatedData.([]byte)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid image data type in context"})
		return
	}

	req := usecase.UploadRequest{
		Data: data,
		Ext:  filepath.Ext(fileHeader.Filename),
	}
	res, err := s.uploadImgUseCase.Upload(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
		return
	}

	response := &pb.UploadImageResponse{
		Path: res.Path,
	}
	ctx.JSON(http.StatusOK, response)
}
