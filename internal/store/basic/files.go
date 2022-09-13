package basic

import (
	"sync"

	"github.com/bulatok/denet_task/pkg/types"

	"github.com/bulatok/denet_task/pkg/logger"
	"go.uber.org/zap"

	"github.com/bulatok/denet_task/internal/models"
)

type files struct {
	mu   *sync.Mutex
	data map[string]interface{}
}

func newFiles() *files {
	return &files{
		mu:   new(sync.Mutex),
		data: map[string]interface{}{},
	}
}

func (f *files) Get(name string) (*models.File, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	file, exist := f.data[name]
	if !exist {
		return nil, models.ErrNotFound
	}

	myFile, ok := file.(*models.File)
	if !ok {
		return nil, models.ErrInternalServer
	}

	return myFile, nil
}

func (f *files) Set(file *models.File) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	if file.Name == "" {
		return models.ErrEmptyFileName
	}

	f.data[file.Name] = file
	return nil
}

func (f *files) All() ([]models.FileOut, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	var allFiles []models.FileOut
	for _, val := range (*f).data {
		myFile, ok := val.(*models.File)
		if !ok {
			logger.Debug("could not cast file", zap.String("where", "basic.All()"))
			continue
		}

		file := models.FileOut{
			Name:                myFile.Name,
			EncodedSizeBytes:    types.MB * myFile.GetCurrentSizeMB(),
			EncodedSizeMB:       myFile.GetCurrentSizeMB(),
			OriginalSizeMB:      myFile.GetUnderLineSizeMB(),
			OriginalSizeInBytes: myFile.GetUnderLineSizeBytes(),
		}
		allFiles = append(allFiles, file)
	}
	return allFiles, nil
}
