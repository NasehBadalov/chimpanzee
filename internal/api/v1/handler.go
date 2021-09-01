package v1

import (
	"chimpanzee/internal/service"
)

// Handler handles v1 routes
type Handler struct {
	service service.IService
}

// NewHandler constructs Handler
func NewHandler(service service.IService) *Handler {
	return &Handler{service: service}
}
