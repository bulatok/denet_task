package handlers

import (
	"net/http"

	"github.com/bulatok/denet_task/pkg/logger"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// UploadFile godoc
// @Summary Загружает файл
// @Tags files
// @Description Загружает файл, который передается в body запроса как multipart/form-data
// @ID file-upload
// @Accept mpfd
// @Success 200
// @Failure 400
// @Router /files/upload [post]
func (h *Handlers) UploadFile(w http.ResponseWriter, r *http.Request) {
	logger.Debug("request", zap.String("endpoint", r.RequestURI))
	file, err := parseFileMultiPart(r)
	if err != nil {
		apiErrorResponse(w, err)
		return
	}

	if err := h.manager.Files.Upload(file); err != nil {
		apiErrorResponse(w, err)
		return
	}

	apiStatusCodeResponse(w, http.StatusOK)
}

// DownloadFile godoc
// @Summary Скачивает файл по имени
// @Tags files
// @Description Возвращает сам файл и статус 200
// @ID file-download
// @Produce mpfd
// @Param fileName path string true "имя файла"
// @Success 200
// @Failure 400
// @Router /files/download/{fileName} [get]
func (h *Handlers) DownloadFile(w http.ResponseWriter, r *http.Request) {
	logger.Debug("request", zap.String("endpoint", r.RequestURI))
	fileName, ok := mux.Vars(r)["fileName"]
	if !ok {
		apiStatusCodeResponse(w, http.StatusBadRequest)
		return
	}

	myFile, err := h.manager.Files.Download(fileName)
	if err != nil {
		apiErrorResponse(w, err)
		return
	}

	apiResponseEncodeFile(w, myFile)
}

// AllFiles godoc
// @Summary Получить информацию по всех хранящимся файлам
// @Tags files
// @Description Получить информацию по всех хранящимся файлам
// @ID file-all-getter
// @Success 200 {object} []models.FileOut
// @Failure 400
// @Router /files/all [get]
func (h *Handlers) AllFiles(w http.ResponseWriter, r *http.Request) {
	logger.Debug("request", zap.String("endpoint", r.RequestURI))
	filesJSON, err := h.manager.Files.All()
	if err != nil {
		apiErrorResponse(w, err)
		return
	}

	apiResponseJSON(w, filesJSON, http.StatusOK)
}
