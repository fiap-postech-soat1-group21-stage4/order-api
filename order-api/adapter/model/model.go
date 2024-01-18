package model

import (
	"time"

	"github.com/fiap-postech-soat1-group21/order-api/order-api/internal/domain/entity"
	"github.com/google/uuid"
)

// OrderResponseDTO is the struct responsible to marshal to json body
type OrderResponseDTO struct {
	ID         uuid.UUID            `json:"id"`
	Status     entity.OrderStatus   `json:"status"`
	CustomerID uuid.UUID            `json:"customer_id"`
	Items      []*entity.OrderItems `json:"items"`
	CreatedAt  time.Time            `json:"created_at"`
	UpdatedAt  time.Time            `json:"updated_at"`
}

type OrderRequestDTO struct {
	CustomerID uuid.UUID  `json:"customer_id"`
	Items      []ItemsDTO `json:"items"`
}

type ItemsDTO struct {
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
}

// OrderResponseList is the struct responsible to marshal to json the response
type OrderResponseList struct {
	Result []*OrderResponseDTO `json:"result"`
	Count  int64               `json:"count"`
}
