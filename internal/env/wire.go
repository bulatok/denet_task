//go:build wireinject
// +build wireinject

package env

import (
	"context"

	"github.com/bulatok/denet_task/internal/api"
	"github.com/bulatok/denet_task/internal/config"
	"github.com/bulatok/denet_task/internal/usecase"
	"github.com/bulatok/denet_task/pkg/logger"

	"github.com/google/wire"
)

type Wire struct {
	Lst  api.Listener
	boot *usecase.BootStrap
}

func (w *Wire) Flush(ctx context.Context) error {
	if err := w.Lst.ShutDown(ctx); err != nil {
		return err
	}
	if err := w.boot.Store.Close(); err != nil {
		return err
	}
	return nil
}
func provideWire(lst api.Listener, conf *config.Config, boot *usecase.BootStrap) *Wire {
	logger.Init(conf)
	return &Wire{
		Lst:  lst,
		boot: boot,
	}
}

func InitWire(configPath string) (*Wire, error) {
	wire.Build(
		config.NewConfig,
		usecase.NewBootStrap,
		api.NewListener,
		provideWire,
	)
	return &Wire{}, nil
}
