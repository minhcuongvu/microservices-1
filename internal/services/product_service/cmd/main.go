package main

import (
	"github.com/minhcuongvu/microservices-1/internal/pkg/http"
	echoserver "github.com/minhcuongvu/microservices-1/internal/pkg/http/echo/server"
	"github.com/minhcuongvu/microservices-1/internal/pkg/logger"
	"go.uber.org/fx"

	"github.com/minhcuongvu/microservices-1/internal/services/product_service/config"
	"github.com/minhcuongvu/microservices-1/internal/services/product_service/server"
)

func main() {
	fx.New(
		fx.Options(
			fx.Provide(
				config.InitConfig,
				logger.InitLogger,
				http.NewContext,
				echoserver.NewEchoServer,
			),
			fx.Invoke(server.RunServers),
		),
	).Run()
}
