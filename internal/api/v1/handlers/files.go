package handlers

import (
	"net/http"

	"github.com/bulatok/denet_task/pkg/logger"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

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

func (h *Handlers) AllFiles(w http.ResponseWriter, r *http.Request) {
	logger.Debug("request", zap.String("endpoint", r.RequestURI))
	filesJSON, err := h.manager.Files.All()
	if err != nil {
		apiErrorResponse(w, err)
		return
	}

	apiResponseJSON(w, filesJSON, http.StatusOK)
}
