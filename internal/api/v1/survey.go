package v1

import (
	"chimpanzee/internal/api/v1/dto"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handler) getSurveys(c echo.Context) error {
	surveys, err := h.service.GetSurveys(c.Request().Context())
	if err != nil {
		return c.String(http.StatusInternalServerError, "")
	}
	return c.JSON(http.StatusOK, dto.NewSurveyList(surveys))
}

func (h *Handler) getSurvey(c echo.Context) error {
	pathID := c.Param("id")
	id, err := strconv.ParseUint(pathID, 10, 64)

	survey, err := h.service.GetSurvey(c.Request().Context(), uint(id))
	if err != nil {
		return c.String(http.StatusInternalServerError, "")
	}
	return c.JSON(http.StatusOK, dto.NewSurvey(survey))
}