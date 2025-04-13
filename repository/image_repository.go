package repository

import (
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

func (r *ImageRepositoryImpl) Save(req SaveRequest) error {
	savePath := filepath.Join(req.Dst, req.FileName)

	if err := os.MkdirAll(req.Dst, os.ModePerm); err != nil {
		return err
	}

	if err := os.WriteFile(savePath, req.Data, os.ModePerm); err != nil {
		return err
	}
	return nil
}
