package main

import (
	"context"
	"flag"

	"github.com/bulatok/denet_task/internal/env"
	"github.com/bulatok/denet_task/pkg/logger"
)

var (
	configPath string
)

// @title       Swagger denet_task API
// @version     0.1
// @description Api for uploading/downloading files
// @contact.url https://bulatok.github.io/

// @host     localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in                         header
// @name                       Authorization
func init() {
	flag.StringVar(&configPath, "config", "config.yml", "select the configs path")
}

func main() {
	wire, err := env.InitWire(configPath)
	if err != nil {
		logger.Fatal(err)
	}

	// clearing the buffer
	defer func() {
		logger.Fatal(wire.Flush(context.Background()))
	}()

	if err := wire.Lst.Listen(); err != nil {
		logger.Fatal(err)
	}
}
