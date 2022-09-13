package api

import (
	"context"

	v1 "github.com/bulatok/denet_task/internal/api/v1"
	"github.com/bulatok/denet_task/internal/config"
	"github.com/bulatok/denet_task/internal/usecase"
)

type Listener interface {
	Listen() error
	ShutDown(context.Context) error
}

func NewListener(conf *config.Config, boot *usecase.BootStrap) Listener {
	return v1.NewServer(conf.HTTP, boot)
}
