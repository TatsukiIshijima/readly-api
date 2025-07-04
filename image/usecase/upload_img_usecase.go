package usecase

import (
	"errors"
	"github.com/google/uuid"
	"path/filepath"
	"readly/configs"
	"readly/image/repository"
)

type UploadImgUseCase interface {
	Upload(req UploadRequest) (*UploadImgResponse, error)
}

type UploadImgUseCaseImpl struct {
	config  configs.Config
	imgRepo repository.ImageRepository
}

func NewUploadImgUseCase(
	config configs.Config,
	imgRepo repository.ImageRepository,
) UploadImgUseCase {
	return &UploadImgUseCaseImpl{
		config:  config,
		imgRepo: imgRepo,
	}
}

type UploadRequest struct {
	Data []byte
	Ext  string
}

// Validate validates the UploadRequest fields
func (r *UploadRequest) Validate() error {
	if len(r.Data) == 0 {
		return errors.New("file data cannot be empty")
	}
	if r.Ext == "" {
		return errors.New("file extension cannot be empty")
	}
	return nil
}

type UploadImgResponse struct {
	Path string
}

func (u *UploadImgUseCaseImpl) Upload(req UploadRequest) (*UploadImgResponse, error) {
	// Validate the UploadRequest fields
	if err := req.Validate(); err != nil {
		return nil, newError(BadRequest, InvalidRequestError, err.Error())
	}

	dst := filepath.Join(configs.ProjectRoot(), ".storage/cover_img")
	fileName := uuid.NewString() + req.Ext
	saveReq := repository.SaveRequest{
		Dst:      dst,
		FileName: fileName,
		Data:     req.Data,
	}

	// Validate the SaveRequest fields
	if err := saveReq.Validate(); err != nil {
		return nil, newError(BadRequest, InvalidRequestError, err.Error())
	}

	err := u.imgRepo.Save(saveReq)
	if err != nil {
		return nil, newError(Internal, InternalServerError, err.Error())
	}
	return &UploadImgResponse{
		Path: filepath.Join(dst, fileName),
	}, nil
}
