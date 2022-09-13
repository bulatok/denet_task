package store

import (
	"github.com/bulatok/denet_task/internal/config"
	"github.com/bulatok/denet_task/internal/store/basic"
	"github.com/bulatok/denet_task/pkg/storage"
)

type Store interface {
	Close() error
	Files() storage.Files
}

// NewStore config is passed for future scaling
func NewStore(conf *config.Config) (Store, error) {
	return basic.NewBasic()
}
