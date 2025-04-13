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
	dst      string
	fileName string
	data     []byte
}

func (r *ImageRepositoryImpl) Save(req SaveRequest) error {
	savePath := filepath.Join(req.dst, req.fileName)

	if err := os.MkdirAll(filepath.Dir(req.dst), os.ModePerm); err != nil {
		return err
	}

	if err := os.WriteFile(savePath, req.data, os.ModePerm); err != nil {
		return err
	}
	return nil
}
