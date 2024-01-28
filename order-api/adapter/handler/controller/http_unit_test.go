package controller_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fiap-postech-soat1-group21-stage4/order-api/order-api/adapter/handler/controller"
	"github.com/fiap-postech-soat1-group21-stage4/order-api/order-api/internal/domain/port/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	t.Run("when body is invalid; should return response 404", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(`{>}`)))
		w := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(w)

		usecaseMock := mocks.NewOrderUseCase(t)
		handler := controller.NewHandler(usecaseMock)

		engine.POST("/order/", handler.CreateOrder)
		engine.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}
