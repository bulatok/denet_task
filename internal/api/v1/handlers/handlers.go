package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/bulatok/denet_task/internal/manager"

	"github.com/bulatok/denet_task/internal/models"
	"github.com/bulatok/denet_task/pkg/logger"

	"github.com/bulatok/denet_task/internal/usecase"
)

type Handlers struct {
	boot    *usecase.BootStrap
	manager *manager.Manager
}

func NewHandlers(boot *usecase.BootStrap) *Handlers {
	return &Handlers{
		boot:    boot,
		manager: manager.NewManager(boot),
	}
}

func apiErrorResponse(w http.ResponseWriter, err error) {
	logger.Debug("error response", zap.String("error", err.Error()))
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	customError, ok := err.(*models.ApiError)
	if !ok {
		logger.Debug("apiErrorResponse: could not parse error")
		return
	}

	w.WriteHeader(customError.StatusCode)
	if err = json.NewEncoder(w).Encode(customError); err != nil {
		logger.Error(err.Error())
	}
}

func apiStatusCodeResponse(w http.ResponseWriter, statusCode int) {
	logger.Debug("status code response", zap.Int("statusCode", statusCode))
	w.WriteHeader(statusCode)
}

func apiResponseJSON(w http.ResponseWriter, res []byte, statusCode int) {
	logger.Debug("json response")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	w.Write(res)
}

func apiResponseEncodeFile(w http.ResponseWriter, myFile *models.File) {
	logger.Debug("file response", zap.String("fileName", myFile.Name))

	w.Header().Set("Content-Type", fmt.Sprintf("%s; charset=utf-8", myFile.ContentType))
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(myFile.Bytes()); err != nil {
		logger.Error(err.Error())
	}
	logger.Debug("success file download")
}
