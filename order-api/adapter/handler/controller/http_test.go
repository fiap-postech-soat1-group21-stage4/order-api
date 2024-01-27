package controller_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/cucumber/godog"
	"github.com/fiap-postech-soat1-group21-stage4/order-api/order-api/adapter/handler/controller"
	"github.com/fiap-postech-soat1-group21-stage4/order-api/order-api/adapter/model"
	"github.com/fiap-postech-soat1-group21-stage4/order-api/order-api/internal/domain/entity"
	mocks "github.com/fiap-postech-soat1-group21-stage4/order-api/order-api/internal/domain/port/mocks"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	orderModelInput = &model.OrderRequestDTO{
		CustomerID: uuid.MustParse("6ba7b810-9dad-11d1-80b4-00c04fd430c8"),
		Items: []model.ItemsDTO{
			{ProductID: uuid.MustParse("550e8400-e29b-41d4-a716-446655440000"), Quantity: 2},
			{ProductID: uuid.MustParse("6ba7b810-9dad-11d1-80b4-00c04fd430c8"), Quantity: 1},
		},
	}

	orderEntityOutput = &entity.Order{
		ID:         uuid.MustParse("123e4567-e89b-12d3-a456-426614174001"),
		CustomerID: orderModelInput.CustomerID,
		Status:     entity.Pending,
	}

	orderItemsOutput = []*entity.OrderItems{
		{OrderID: orderEntityOutput.ID, ProductID: orderModelInput.Items[0].ProductID, Quantity: orderModelInput.Items[0].Quantity},
		{OrderID: orderEntityOutput.ID, ProductID: orderModelInput.Items[1].ProductID, Quantity: orderModelInput.Items[1].Quantity},
	}
)

type handlerContext struct {
	handler     *controller.Handler
	w           *httptest.ResponseRecorder
	req         *http.Request
	err         error
	body        []byte
	usecaseMock *mocks.OrderUseCase
}

func (h *handlerContext) theFollowingOrderDetails(table *godog.Table) error {
	orderModelInput = &model.OrderRequestDTO{
		CustomerID: uuid.MustParse(table.Rows[1].Cells[0].Value),
		Items: []model.ItemsDTO{
			{ProductID: uuid.MustParse(table.Rows[1].Cells[1].Value), Quantity: atoi(table.Rows[1].Cells[2].Value)},
			{ProductID: uuid.MustParse(table.Rows[2].Cells[1].Value), Quantity: atoi(table.Rows[2].Cells[2].Value)},
		},
	}
	return nil
}

func (h *handlerContext) aRequestIsMadeToCreateTheOrder() error {
	jsonBytes, err := json.Marshal(orderModelInput)
	if err != nil {
		return err
	}

	h.req = httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBuffer(jsonBytes))
	h.w = httptest.NewRecorder()

	ctxGin, _ := gin.CreateTestContext(h.w)
	ctxGin.Request = h.req

	h.usecaseMock = &mocks.OrderUseCase{}
	h.usecaseMock.On("CreateOrder", ctxGin, mock.AnythingOfType("*entity.Order")).Return(orderEntityOutput, h.err).Once()
	h.usecaseMock.On("CreateOrderItems", ctxGin, mock.AnythingOfType("[]*entity.OrderItems")).Return(orderItemsOutput, h.err).Once()

	h.handler = controller.NewHandler(h.usecaseMock)
	h.handler.CreateOrder(ctxGin)

	res := h.w.Result()
	defer res.Body.Close()
	h.body = h.w.Body.Bytes()

	return nil
}

func (h *handlerContext) theResponseShouldHaveStatusCode(statusCode int) error {
	return assertExpectedAndActual(assert.Equal, statusCode, h.w.Code, "status code")
}

func (h *handlerContext) theResponseBodyShouldMatchTheExpectedOrderDetails() error {
	wantGot, err := json.Marshal(model.OrderResponseDTO{
		ID:         orderEntityOutput.ID,
		Status:     orderEntityOutput.Status,
		CustomerID: orderEntityOutput.CustomerID,
		CreatedAt:  orderEntityOutput.CreatedAt,
		UpdatedAt:  orderEntityOutput.UpdatedAt,
		Items:      orderItemsOutput,
	})
	if err != nil {
		return err
	}

	return assertExpectedAndActual(assert.Equal, string(wantGot), strings.TrimSuffix(h.w.Body.String(), "\n"), "response body")
}

func atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}

func assertExpectedAndActual(a expectedAndActualAssertion, expected, actual interface{}, msgAndArgs ...interface{}) error {
	var t asserter
	a(&t, expected, actual, msgAndArgs...)
	return t.err
}

type expectedAndActualAssertion func(t assert.TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool

type asserter struct {
	err error
}

func (a *asserter) Errorf(format string, args ...interface{}) {
	a.err = fmt.Errorf(format, args...)
}

// InitializeScenario initializes the context for godog scenarios
func InitializeScenario(s *godog.ScenarioContext) {
	h := &handlerContext{}
	s.Given(`^the following order details$`, h.theFollowingOrderDetails)
	s.When(`^a request is made to create the order$`, h.aRequestIsMadeToCreateTheOrder)
	s.Then(`^the response should have status code (\d+)$`, h.theResponseShouldHaveStatusCode)
	s.Step(`^the response body should match the expected order details$`, h.theResponseBodyShouldMatchTheExpectedOrderDetails)

	s.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		h.err = nil
		return ctx, nil
	})

	s.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		return ctx, nil
	})
}

// TestFeatures runs the godog feature tests
func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		Name:                 "http",
		ScenarioInitializer:  InitializeScenario,
		TestSuiteInitializer: InitializeTestSuite,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"../../../features/http.feature"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

// InitializeTestSuite initializes the godog test suite
func InitializeTestSuite(ctx *godog.TestSuiteContext) {}

// TODO: Import necessary packages and define other utility functions if needed.
