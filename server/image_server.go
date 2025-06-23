package server

import (
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
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

	file, err := fileHeader.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to close file"})
			return
		}
	}(file)

	data, err := io.ReadAll(file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
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
