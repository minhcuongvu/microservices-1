package main

import (
	"github.com/minhcuongvu/microservices-1/internal/pkg/logger"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Options(
			fx.Provide(
				logger.InitLogger,
			),
			fx.Invoke(server.RunServers),
		),
	).Run()
}
