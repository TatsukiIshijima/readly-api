package repository

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
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

// Validate validates the SaveRequest fields and prevents path traversal attacks
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

	// Prevent path traversal attacks
	if filepath.IsAbs(r.FileName) {
		return errors.New("file name cannot be an absolute path")
	}

	// Check if the filename contains path traversal sequences
	if r.FileName != filepath.Clean(r.FileName) {
		return errors.New("file name contains invalid path sequences")
	}

	// Ensure the cleaned path doesn't try to go outside the destination directory
	cleanedPath := filepath.Join(r.Dst, r.FileName)
	destinationPath, err := filepath.Abs(r.Dst)
	if err != nil {
		return err
	}

	cleanedAbsPath, err := filepath.Abs(cleanedPath)
	if err != nil {
		return err
	}

	// Check if the final path is still within the destination directory
	relPath, err := filepath.Rel(destinationPath, cleanedAbsPath)
	if err != nil {
		return err
	}

	// If the relative path starts with "..", it means the path is outside the destination directory
	if strings.HasPrefix(relPath, ".."+string(filepath.Separator)) || relPath == ".." {
		return errors.New("path traversal attack detected")
	}

	// Check if the filename contains directory components
	if filepath.Base(r.FileName) != r.FileName {
		return errors.New("file name cannot contain directory components")
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
