package repository

import (
	"errors"
	"os"
	"path/filepath"
)

type ImageRepository interface {
	Save(req SaveRequest) error
}

type ImageRepositoryImpl struct{}

func NewImageRepository() ImageRepository {
	return &ImageRepositoryImpl{}
}

type SaveRequest struct {
	Dst      string
	FileName string
	Data     []byte
}

// Validate validates the SaveRequest fields
func (r *SaveRequest) Validate() error {
	if r.Dst == "" {
		return errors.New("destination directory cannot be empty")
	}
	if r.FileName == "" {
		return errors.New("file name cannot be empty")
	}
	if len(r.Data) == 0 {
		return errors.New("file data cannot be empty")
	}
	return nil
}

func (r *ImageRepositoryImpl) Save(req SaveRequest) error {
	savePath := filepath.Join(req.Dst, req.FileName)

	if err := os.MkdirAll(req.Dst, 0755); err != nil {
		return err
	}

	if err := os.WriteFile(savePath, req.Data, 0644); err != nil {
		return err
	}
	return nil
}
