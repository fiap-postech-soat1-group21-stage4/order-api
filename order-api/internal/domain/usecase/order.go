package usecase

import (
	"context"

	"github.com/fiap-postech-soat1-group21/order-api/order-api/internal/domain/entity"
	"github.com/fiap-postech-soat1-group21/order-api/order-api/internal/domain/port"
	"github.com/google/uuid"
)

type useCaseOrder struct {
	repository port.OrderRepository
}

// NewOrderUseCase is responsible for all use cases for orders
func NewOrderUseCase(orderRepo port.OrderRepository) port.OrderUseCase {
	return &useCaseOrder{
		repository: orderRepo,
	}
}

// List retrieves all orders
func (u *useCaseOrder) List(ctx context.Context) (*entity.OrderResponseList, error) {
	res, err := u.repository.List(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *useCaseOrder) GetOrderItems(ctx context.Context, orderID uuid.UUID) ([]*entity.OrderItems, error) {
	res, err := u.repository.GetOrderItems(ctx, orderID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *useCaseOrder) CreateOrder(ctx context.Context, input *entity.Order) (*entity.Order, error) {
	res, err := u.repository.CreateOrder(ctx, input)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *useCaseOrder) CreateOrderItems(ctx context.Context, input []*entity.OrderItems) ([]*entity.OrderItems, error) {
	res, err := u.repository.CreateOrderItems(ctx, input)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *useCaseOrder) UpdateOrderStatus(ctx context.Context, orderID uuid.UUID, status string) error {
	err := u.repository.UpdateOrderStatus(ctx, orderID, status)
	if err != nil {
		return err
	}

	return nil
}
