package v1

import (
	"context"
	"net/http"

	"github.com/bulatok/denet_task/internal/api/v1/handlers"
	"github.com/bulatok/denet_task/internal/config"
	"github.com/bulatok/denet_task/internal/usecase"
	"github.com/bulatok/denet_task/pkg/logger"

	"github.com/rs/cors"
	"go.uber.org/zap"
)

type ApiServer struct {
	httpSrv *http.Server
}

func (a *ApiServer) Listen() error {
	logger.Info(
		"started listening server on",
		zap.String("addr", a.httpSrv.Addr),
	)
	return a.httpSrv.ListenAndServe()
}

func (a *ApiServer) ShutDown(ctx context.Context) error {
	logger.Info("stopped server")
	return a.httpSrv.Shutdown(ctx)
}

func NewServer(conf *config.HTTP, boot *usecase.BootStrap) *ApiServer {
	h := handlers.NewHandlers(boot)
	router := newRouter(h)

	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowedHeaders: []string{
			"*",
		},
		AllowCredentials: true,
	})

	s := &ApiServer{
		httpSrv: &http.Server{
			Addr:    conf.Port,
			Handler: corsOptions.Handler(router),
		},
	}
	return s
}
