package usecase

import (
	"github.com/bulatok/denet_task/internal/config"
	"github.com/bulatok/denet_task/internal/store"
)

type BootStrap struct {
	Store store.Store
}

func NewBootStrap(conf *config.Config) (*BootStrap, error) {
	st, err := store.NewStore(conf)
	if err != nil {
		return nil, err
	}
	return &BootStrap{Store: st}, nil
}
