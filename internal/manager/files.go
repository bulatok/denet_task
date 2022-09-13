package manager

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/bulatok/denet_task/internal/models"
	"github.com/bulatok/denet_task/internal/usecase"
	"github.com/bulatok/denet_task/pkg/logger"

	"go.uber.org/zap"
)

type FilesManager interface {
	Download(name string) (*models.File, error)
	Upload(header *multipart.FileHeader) error
	All() ([]byte, error)
}

func NewFilesManager(boot *usecase.BootStrap) FilesManager {
	return &filesManager{
		boot: boot,
	}
}

type filesManager struct {
	boot *usecase.BootStrap
}

// All returns all files as JSON
func (f *filesManager) All() ([]byte, error) {
	allFiles, err := f.boot.Store.Files().All()
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(allFiles)
	if err != nil {
		logger.Info(
			"could not marshal files",
			zap.String("error", err.Error()))
		return nil, models.ErrInternalServer
	}

	return data, nil
}

func (f *filesManager) Download(name string) (*models.File, error) {
	myFile, err := f.boot.Store.Files().Get(name)
	if err != nil {
		return nil, err
	}

	// decoding file
	fileDecoded, err := models.DecodeFile(*myFile)
	if err != nil {
		return nil, err
	}

	return fileDecoded, nil
}

func (f *filesManager) Upload(header *multipart.FileHeader) error {
	myFile := &models.File{Name: header.Filename}
	httpFile, err := header.Open()
	if err != nil {
		return models.ErrCouldNotOpenFile
	}

	data, err := io.ReadAll(httpFile)
	if err != nil {
		return models.ErrCouldNotReadFile
	}

	contentType := http.DetectContentType(data)

	myFile.SetData(data)
	myFile.SetContentType(contentType)
	if err := myFile.Encode(); err != nil {
		return err
	}

	return f.boot.Store.Files().Set(myFile)
}
