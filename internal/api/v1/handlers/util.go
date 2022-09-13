package handlers

import (
	"mime/multipart"
	"net/http"

	"github.com/bulatok/denet_task/pkg/logger"
	"go.uber.org/zap"

	"github.com/bulatok/denet_task/pkg/types"

	"github.com/bulatok/denet_task/internal/models"
)

const (
	// 1 GB for example
	maxFileSize = types.MB * 1024
)

func parseFileMultiPart(r *http.Request) (*multipart.FileHeader, error) {
	if err := r.ParseMultipartForm(maxFileSize); err != nil {
		return nil, models.ErrInvalidMultipartForm
	}

	form := r.MultipartForm
	if len(form.File) != 1 { // only 1 file
		return nil, models.ErrTooManyFiles
	}

	files := form.File
	fileName := ""
	for name, _ := range files { // we checked that files has at least 1 file
		fileName = name
	}

	return files[fileName][0], nil
}

func (h *Handlers) Ping(w http.ResponseWriter, r *http.Request) {
	logger.Debug("request", zap.String("endpoint", r.RequestURI))

	apiResponseJSON(w, []byte(`{"message": "pong"}`), http.StatusOK)
}
