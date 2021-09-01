package v1

import "github.com/labstack/echo/v4"

// RegisterRoutes adds routes for handlers of "/api/v1/*"
func (h *Handler) RegisterRoutes(echo *echo.Group) {
	v1 := echo.Group("/api/v1")

	v1.GET("/surveys", h.getSurveys)
	v1.GET("/surveys/:id", h.getSurvey)
	v1.GET("/reports", h.getReports)
	v1.POST("/reports", h.postReport)
}
