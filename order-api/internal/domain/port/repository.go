package port

import (
	"context"

	"github.com/fiap-postech-soat1-group21-stage4/order-api/order-api/internal/domain/entity"
	"github.com/google/uuid"
)

// OrderRepository is the interface for order database
type OrderRepository interface {
	List(context.Context) (*entity.OrderResponseList, error)
	GetOrderItems(ctx context.Context, orderID uuid.UUID) ([]*entity.OrderItems, error)
	UpdateOrderStatus(ctx context.Context, orderID uuid.UUID, status string) error
	CreateOrder(context.Context, *entity.Order) (*entity.Order, error)
	CreateOrderItems(ctx context.Context, order []*entity.OrderItems) ([]*entity.OrderItems, error)
}
