package controller

import (
	"fmt"
	"net/http"

	"github.com/fiap-postech-soat1-group21-stage4/order-api/order-api/adapter/model"
	"github.com/fiap-postech-soat1-group21-stage4/order-api/order-api/internal/domain/entity"
	"github.com/gin-gonic/gin"
)

var (
	MaxQuantityPerOrder = 10
)

func (h *Handler) List(ctx *gin.Context) {
	res, err := h.useCase.List(ctx)
	if err != nil {
		return
	}

	responseItems := make([]*model.OrderResponseDTO, 0, len(res.Result))

	for _, item := range res.Result {
		orderItems, err := h.useCase.GetOrderItems(ctx, item.ID)
		if err != nil {
			return
		}

		responseItems = append(responseItems, &model.OrderResponseDTO{
			ID:         item.ID,
			Status:     item.Status,
			Items:      orderItems,
			CustomerID: item.CustomerID,
			CreatedAt:  item.CreatedAt,
			UpdatedAt:  item.UpdatedAt,
		})
		fmt.Printf("%+v", responseItems)
	}

	output := &model.OrderResponseList{
		Result: responseItems,
		Count:  res.Count,
	}

	ctx.JSON(http.StatusOK, output)
}

func (h *Handler) CreateOrder(ctx *gin.Context) {
	var input *model.OrderRequestDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	domainOrder := &entity.Order{
		CustomerID: input.CustomerID,
		Status:     entity.Pending,
	}

	res, err := h.useCase.CreateOrder(ctx, domainOrder)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	orderItems := make([]*entity.OrderItems, 0, MaxQuantityPerOrder)

	for _, item := range input.Items {
		fmt.Print(item)
		orderItems = append(orderItems, &entity.OrderItems{
			OrderID:   res.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	resItems, err := h.useCase.CreateOrderItems(ctx, orderItems)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	output := &model.OrderResponseDTO{
		ID:         res.ID,
		Status:     res.Status,
		CustomerID: res.CustomerID,
		CreatedAt:  res.CreatedAt,
		UpdatedAt:  res.UpdatedAt,
		Items:      resItems,
	}

	ctx.JSON(http.StatusCreated, output)
}
