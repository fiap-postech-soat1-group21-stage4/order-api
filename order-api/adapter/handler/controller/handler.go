package controller

import (
	"github.com/fiap-postech-soat1-group21-stage4/order-api/order-api/internal/domain/port"
	"github.com/gin-gonic/gin"
)

// Handler provides order funcionalities
type Handler struct {
	useCase port.OrderUseCase
}

// NewHandler is the order handler builder
func NewHandler(u port.OrderUseCase) *Handler {
	return &Handler{
		useCase: u,
	}
}

// RegisterRoutes register routes
func (h *Handler) RegisterRoutes(routes *gin.RouterGroup) {
	orderRoute := routes.Group("/orders")
	orderRoute.GET("", h.List)
	orderRoute.POST("", h.CreateOrder)
}
