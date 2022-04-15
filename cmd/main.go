package main

import (
	// "fmt"

	"bitbucket.org/udevs/example_api_gateway/api"

	"bitbucket.org/udevs/example_api_gateway/config"
	"bitbucket.org/udevs/example_api_gateway/pkg/logger"
	"bitbucket.org/udevs/example_api_gateway/services"
	// "github.com/gomodule/redigo/redis"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "example_api_gateway")

	gprcClients, _ := services.NewGrpcClients(&cfg)

	server := api.New(&api.RouterOptions{
		Log:      log,
		Cfg:      cfg,
		Services: gprcClients,
	})

	server.Run(cfg.HttpPort)
}
