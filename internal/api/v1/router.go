package v1

import (
	"net/http"

	"github.com/bulatok/denet_task/internal/api/v1/handlers"

	"github.com/gorilla/mux"
)

const (
	POST = http.MethodPost
	GET  = http.MethodGet
)

func newRouter(h *handlers.Handlers) *mux.Router {
	r := mux.NewRouter()
	r.Methods(GET).Path("/ping").HandlerFunc(h.Ping)

	filesRouter := r.PathPrefix("/files").Subrouter()
	filesRouter.Methods(POST).Path("/upload").HandlerFunc(h.UploadFile)
	filesRouter.Methods(GET).Path("/download/{fileName}").HandlerFunc(h.DownloadFile)
	filesRouter.Methods(GET).Path("/all").HandlerFunc(h.AllFiles)

	return r
}
