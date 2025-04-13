package usecase

import (
	"github.com/google/uuid"
	"path/filepath"
	"readly/env"
	"readly/repository"
)

type UploadImgUseCase interface {
	Upload(req UploadRequest) (*UploadImgResponse, error)
}

type UploadImgUseCaseImpl struct {
	config  env.Config
	imgRepo repository.ImageRepository
}

func NewUploadImgUseCase(
	config env.Config,
	imgRepo repository.ImageRepository,
) UploadImgUseCase {
	return &UploadImgUseCaseImpl{
		config:  config,
		imgRepo: imgRepo,
	}
}

type UploadRequest struct {
	data []byte
	ext  string
}

type UploadImgResponse struct {
	Path string
}

func (u *UploadImgUseCaseImpl) Upload(req UploadRequest) (*UploadImgResponse, error) {
	dst := filepath.Join(env.ProjectRoot(), ".storage/cover_img")
	fileName := uuid.NewString() + req.ext
	saveReq := repository.SaveRequest{
		Dst:      dst,
		FileName: fileName,
		Data:     req.data,
	}
	err := u.imgRepo.Save(saveReq)
	if err != nil {
		return nil, err
	}
	return &UploadImgResponse{
		Path: filepath.Join(dst, fileName),
	}, nil
}
