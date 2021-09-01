package main

import (
	"chimpanzee/internal/api"
	"chimpanzee/internal/api/v1"
	"chimpanzee/internal/config"
	"chimpanzee/internal/log"
	"chimpanzee/internal/repository/postgres"
	"chimpanzee/internal/service"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	// init logger
	err := log.InitLogger()
	if err != nil {
		panic(err)
	}

	fx.New(
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: zap.L()}
		}),
		fx.Provide(
			// Provide logger
			zap.L,
			// Provide configuration
			config.NewConfig,
			// Provide repository
			postgres.ProvideRepository,
			// Provide service
			service.ProvideService,
			// Create new server and register to lifecycle
			api.NewServer,
			// Add all handlers here
			func(cfg *config.Config, svc service.IService) []api.Handler {
				return []api.Handler{
					v1.NewHandler(svc),  // v1
				}
			},
		),
		fx.Invoke(
			// Register all routes to http server
			api.RegisterRoutes,
		),
	).Run()
}
