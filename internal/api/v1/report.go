package v1

import (
	"chimpanzee/internal/api/v1/dto"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

func (h *Handler) postReport(c echo.Context) error {
	var r dto.Report

	if err := c.Bind(&r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Cannot parse request"})
	}

	ra, err := dto.NewReportAnswer(r.Answers)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Cannot parse request"})
	}

	err = h.service.AddSurveyReport(c.Request().Context(), r.SurveyID, ra)
	if err != nil {
		zap.S().With("error", err).Error("Adding survey failed")
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server occurred"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Success"})
}

func (h *Handler) getReports(c echo.Context) error {
	rr, err := h.service.GetReports(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dto.NewReportResults(rr))
}
