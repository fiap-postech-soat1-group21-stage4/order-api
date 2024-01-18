package entity

import (
	"time"

	"github.com/google/uuid"
)

// Order domain table model
type Order struct {
	ID         uuid.UUID     `gorm:"type:uuid;default:gen_random_uuid()"`
	Status     OrderStatus   `gorm:"not null"`
	OrderItems []*OrderItems `gorm:"foreignKey:OrderID"`
	CustomerID uuid.UUID
	CreatedAt  time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"not null;autoUpdateTime"`
}

type OrderItems struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	OrderID   uuid.UUID `gorm:"not null"`
	ProductID uuid.UUID `gorm:"not null"`
	Quantity  int       `gorm:"not_null"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime"`
}

type Item struct {
	ProductID uuid.UUID `gorm:"not null"`
	Quantity  int       `gorm:"not_null"`
}

// OrderResponseList summary list
type OrderResponseList struct {
	Result []*Order
	Count  int64
}

type PaymentStatus string

const (
	Approved = "approved"
	Rejected = "rejected"
)

func (s PaymentStatus) String() string {
	switch s {
	case Approved:
		return Approved
	case Rejected:
		return Rejected
	}
	return "unknown"
}

type OrderStatus string

const (
	Pending   = "pending"
	Received  = "received"
	Preparing = "preparing"
	Ready     = "ready"
	Finished  = "finished"
	Cancelled = "cancelled"
)

func (s OrderStatus) String() string {
	switch s {
	case Pending:
		return Pending
	case Received:
		return Received
	case Preparing:
		return Preparing
	case Ready:
		return Ready
	case Finished:
		return Finished
	}
	return "unknown"
}
