package repository

import (
	"context"
	"fmt"

	"github.com/fiap-postech-soat1-group21-stage4/order-api/order-api/internal/domain/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Order struct type
type Order struct {
	db *gorm.DB
}

// NewOrderRepository instantiates order repository
func NewOrderRepository(db *gorm.DB) *Order {
	return &Order{db}
}

// List retrives all orders.
func (o *Order) List(ctx context.Context) (*entity.OrderResponseList, error) {
	dbFn := o.db.WithContext(ctx)

	var count int64
	var order []*entity.Order

	result := dbFn.Table("order").Find(&order).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	return &entity.OrderResponseList{
		Result: order,
		Count:  count,
	}, nil
}

func (o *Order) GetOrderItems(ctx context.Context, orderID uuid.UUID) ([]*entity.OrderItems, error) {
	dbFn := o.db.WithContext(ctx)

	var orderItems []*entity.OrderItems
	result := dbFn.Table("order_item").Where("order_id = ?", orderID).Find(&orderItems)

	fmt.Sprintf("%+v", orderItems)
	if result.Error != nil {
		return nil, result.Error
	}

	return orderItems, nil
}

func (o *Order) CreateOrder(ctx context.Context, order *entity.Order) (*entity.Order, error) {
	dbFn := o.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true})

	if result := dbFn.Table("order").Create(order); result.Error != nil {
		return nil, result.Error
	}

	return order, nil
}

func (o *Order) CreateOrderItems(ctx context.Context, orderItems []*entity.OrderItems) ([]*entity.OrderItems, error) {
	dbFn := o.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true})

	if result := dbFn.Table("order_item").Create(orderItems); result.Error != nil {
		return nil, result.Error
	}

	return orderItems, nil
}

func (o *Order) UpdateOrderStatus(ctx context.Context, orderID uuid.UUID, status string) error {
	dbFn := o.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true})

	if result := dbFn.Table("order").Where("id = ?", orderID).Update("status", status); result.Error != nil {
		return result.Error
	}

	return nil
}
