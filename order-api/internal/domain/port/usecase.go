package port

import (
	"context"

	"github.com/fiap-postech-soat1-group21-stage4/order-api/order-api/internal/domain/entity"
	"github.com/google/uuid"
)

// OrderUseCase is the interface for order repository
type OrderUseCase interface {
	CreateOrder(ctx context.Context, order *entity.Order) (*entity.Order, error)
	GetOrderItems(ctx context.Context, orderID uuid.UUID) ([]*entity.OrderItems, error)
	UpdateOrderStatus(ctx context.Context, orderID uuid.UUID, status string) error
	CreateOrderItems(ctx context.Context, orderItems []*entity.OrderItems) ([]*entity.OrderItems, error)
	List(context.Context) (*entity.OrderResponseList, error)
}
