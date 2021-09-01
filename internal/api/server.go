package api

import (
	"chimpanzee/internal/config"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	gommonLog "github.com/labstack/gommon/log"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// NewServer creates new echo server object and registers start and end of lifecycle of app
// to start echo on start and gracefully shut it down on exit
func NewServer(lc fx.Lifecycle, cfg *config.Config) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	// avoid any native logging of echo, because we use custom library for logging
	e.HideBanner = true              // don't log the banner on startup
	e.HidePort = true                // hide log about port server started on
	e.Logger.SetLevel(gommonLog.OFF) // disable echo#Logger

	// Add hook to start server whenever app starts
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			zap.S().Infof("server running on port: %d", cfg.Service.Port)
			go func() {
				err := e.Start(fmt.Sprintf(":%d", cfg.Service.Port))
				if err != nil {
					zap.S().Fatalf("Error while starting server: %s", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			zap.S().Info("Server is shutting down")
			return e.Shutdown(ctx)
		},
	})

	return e
}
