package manage

import (
	o "github.com/fiap-postech-soat1-group21-stage4/order-api/order-api/adapter/handler/controller"
	"github.com/fiap-postech-soat1-group21-stage4/order-api/order-api/internal/domain/port"
	"github.com/gin-gonic/gin"
)

type apps interface {
	RegisterRoutes(routes *gin.RouterGroup)
}

type Manage struct {
	order apps
}

type UseCases struct {
	Order port.OrderUseCase
}

func New(uc *UseCases) *Manage {

	orderHandler := o.NewHandler(uc.Order)

	return &Manage{
		order: orderHandler,
	}
}

func (m *Manage) RegisterRoutes(group *gin.RouterGroup) {
	m.order.RegisterRoutes(group)
}
