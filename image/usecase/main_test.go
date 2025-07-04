package usecase

import (
	"log"
	"os"
	"path/filepath"
	"readly/configs"
	"readly/image/repository"
	"testing"
)

func TestMain(m *testing.M) {
	setupMain()
	os.Exit(m.Run())
}

var config configs.Config

func setupMain() {
	c, err := configs.Load(filepath.Join(configs.ProjectRoot(), "/configs/env"))
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}
	config = c
}

func newTestUploadImgUseCase(t *testing.T) UploadImgUseCase {
	imgRepo := repository.NewImageRepository()
	return NewUploadImgUseCase(config, imgRepo)
}
