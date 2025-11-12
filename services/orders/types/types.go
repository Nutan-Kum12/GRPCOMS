package types

import (
	"context"

	"github.com/Nutan-Kum12/GRPCOMS.git/services/common/genproto/orders"
)

type OrderService interface {
	// Define methods for the OrderService
	CreateOrder(context.Context, *orders.Order) error
	GetOrder(context.Context) []*orders.Order
}
