package handlers

import (
	"encoding/json"
	"mime/multipart"
	"net/http"

	"github.com/bulatok/denet_task/internal/models"
	"github.com/bulatok/denet_task/pkg/logger"
	"github.com/bulatok/denet_task/pkg/types"

	"go.uber.org/zap"
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

type pongResponse struct {
	Message string `json:"message"`
}

// Ping godoc
// @Tags ping
// @Description Пингует сервер
// @ID ping-server
// @Success 200 {object} pongResponse
// @Router /ping [get]
func (h *Handlers) Ping(w http.ResponseWriter, r *http.Request) {
	logger.Debug("request", zap.String("endpoint", r.RequestURI))

	resp := &pongResponse{Message: "pong"}
	d, _ := json.Marshal(resp)
	apiResponseJSON(w, d, http.StatusOK)
}
