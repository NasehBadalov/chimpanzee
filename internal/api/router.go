package api

import (
	"chimpanzee/internal/config"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

// RegisterRoutes iterates over handlers and registers them in given echo server instance
func RegisterRoutes(cfg *config.Config, e *echo.Echo, handlers []Handler) {
	e.Use(middleware.Recover())

	// Register all passed routes
	for _, h := range handlers {
		h.RegisterRoutes(e.Group(""))
	}

	// Log all routes if debug mode is enabled
	if cfg.Service.Debug {
		var routeList []string

		// Make an array of routes(formatted strings)
		for _, r := range e.Routes() {
			routeList = append(routeList, fmt.Sprintf("[%s] %s", r.Method, r.Path))
		}
		zap.S().With("routes", routeList).Infof("Routers are intialized")
	}
}
