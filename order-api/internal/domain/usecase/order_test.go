package usecase_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/cucumber/godog"
	"github.com/fiap-postech-soat1-group21-stage4/order-api/order-api/internal/domain/entity"
	"github.com/fiap-postech-soat1-group21-stage4/order-api/order-api/internal/domain/port"
	"github.com/fiap-postech-soat1-group21-stage4/order-api/order-api/internal/domain/port/mocks"
	"github.com/fiap-postech-soat1-group21-stage4/order-api/order-api/internal/domain/usecase"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	defaultOrderID  = uuid.MustParse("b4dacf92-7000-4523-9fab-166212acc92d")
	ctxDefaultOrder = context.Background()

	givenOrder = &entity.Order{
		ID:         defaultOrderID,
		Status:     entity.Pending,
		CustomerID: uuid.MustParse("6ba7b810-9dad-11d1-80b4-00c04fd430c8"),
	}

	wantedOrder = &entity.Order{
		ID:         defaultOrderID,
		Status:     entity.Pending,
		CustomerID: uuid.MustParse("6ba7b810-9dad-11d1-80b4-00c04fd430c8"),
	}

	orderItems = []*entity.OrderItems{
		{
			ID:        uuid.MustParse("6ba7b811-9dad-11d1-80b4-00c04fd430c8"),
			OrderID:   defaultOrderID,
			ProductID: uuid.MustParse("550e8400-e29b-41d4-a716-446655440000"),
			Quantity:  2,
		},
		{
			ID:        uuid.MustParse("6ba7b812-9dad-11d1-80b4-00c04fd430c8"),
			OrderID:   defaultOrderID,
			ProductID: uuid.MustParse("6ba7b810-9dad-11d1-80b4-00c04fd430c8"),
			Quantity:  3,
		},
	}
)

func InitializeScenario(s *godog.ScenarioContext) {
	orderTest := &orderTest{}

	s.Given(`^an existing customer with ID "([^"]*)"$`, orderTest.anExistingCustomerWithID)
	s.Given(`^the following items are added to the order:$`, orderTest.theFollowingItemsAreAddedToTheOrder)
	s.When(`^the order is created$`, orderTest.theOrderIsCreated)
	s.Step(`^the order should have the following items:$`, orderTest.theOrderShouldHaveTheFollowingItems)
}

type orderTest struct {
	service     port.OrderUseCase
	repo        *mocks.OrderRepository
	err         error
	result      *entity.Order
	resultItems []*entity.OrderItems
}

func (o *orderTest) anExistingCustomerWithID(customerID string) error {
	givenOrder.CustomerID, _ = uuid.Parse(customerID)
	return nil
}

func (o *orderTest) theFollowingItemsAreAddedToTheOrder(table *godog.Table) error {
	for _, row := range table.Rows {
		productID := row.Cells[0].Value
		quantity, _ := strconv.Atoi(row.Cells[1].Value)

		item := &entity.OrderItems{
			OrderID:   defaultOrderID,
			ProductID: uuid.MustParse(productID),
			Quantity:  quantity,
		}

		orderItems = append(orderItems, item)
	}
	return nil
}

func (o *orderTest) theOrderIsCreated() error {
	o.repo = &mocks.OrderRepository{}
	o.service = usecase.NewOrderUseCase(o.repo)

	o.repo.On("CreateOrder", ctxDefaultOrder, givenOrder).Return(wantedOrder, nil).Once()
	o.repo.On("CreateOrderItems", ctxDefaultOrder, orderItems).Return(orderItems, nil).Once()
	o.repo.On("GetOrderItems", ctxDefaultOrder, defaultOrderID).Return(orderItems, nil).Once()

	o.result, o.err = o.service.CreateOrder(ctxDefaultOrder, givenOrder)
	o.resultItems, _ = o.service.GetOrderItems(ctxDefaultOrder, defaultOrderID)
	return nil
}

func (o *orderTest) theOrderShouldHaveTheFollowingItems(table *godog.Table) error {
	var expectedItems []*entity.OrderItems
	for _, row := range table.Rows[1:] {
		productID := row.Cells[0].Value
		quantity, _ := strconv.Atoi(row.Cells[1].Value)

		item := &entity.OrderItems{
			OrderID:   defaultOrderID,
			ProductID: uuid.MustParse(productID),
			Quantity:  quantity,
		}

		expectedItems = append(expectedItems, item)
	}

	assertExpectedAndActual(assert.EqualValues, expectedItems, o.resultItems, "Expected items do not match actual items")
	return nil
}

// assertExpectedAndActual is a helper function to allow the step function to call
// assertion functions where you want to compare an expected and an actual value.
func assertExpectedAndActual(a expectedAndActualAssertion, expected, actual interface{}, msgAndArgs ...interface{}) error {
	var t asserter
	a(&t, expected, actual, msgAndArgs...)
	return t.err
}

type expectedAndActualAssertion func(t assert.TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool

// asserter is used to be able to retrieve the error reported by the called assertion
type asserter struct {
	err error
}

// Errorf is used by the called assertion to report an error
func (a *asserter) Errorf(format string, args ...interface{}) {
	a.err = fmt.Errorf(format, args...)
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		Name:                 "order",
		ScenarioInitializer:  InitializeScenario,
		TestSuiteInitializer: InitializeTestSuite,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"../../../features/order.feature"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {}
