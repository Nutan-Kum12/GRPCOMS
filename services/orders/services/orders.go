package services

import (
	"context"

	"github.com/Nutan-Kum12/GRPCOMS.git/services/common/genproto/orders"
)

var ordersDb = make([]orders.Order, 0)

type OrdersService struct {
	// Service fields and methods
}

func NewOrdersService() *OrdersService {
	return &OrdersService{}
}
func (s *OrdersService) CreateOrder(ctx context.Context, order *orders.Order) error {
	// Business logic to create an order
	ordersDb = append(ordersDb, *order)
	return nil
}
func (s *OrdersService) GetOrder(ctx context.Context) []*orders.Order {
	// Convert []orders.Order to []*orders.Order
	result := make([]*orders.Order, len(ordersDb))
	for i := range ordersDb {
		result[i] = &ordersDb[i]
	}
	return result
}
