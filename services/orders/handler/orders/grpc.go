package handler

import (
	"context"

	"github.com/Nutan-Kum12/GRPCOMS.git/services/common/genproto/orders"
	"github.com/Nutan-Kum12/GRPCOMS.git/services/orders/types"
	"google.golang.org/grpc"
)

type OrdersGrpcHandler struct {
	orderService types.OrderService
	//service injection
	orders.UnimplementedOrderServiceServer
	// Without this embedding, if your
	// file defines 5 RPC methods but your
	// handler only implements 3 of them, your
	// code wouldn't compile. With UnimplementedOrderServiceServer embedded:
	// The 3 methods you implement will work normally
	// The 2 unimplemented methods will automatically return a gRPC error with status UNIMPLEMENTED

	// Partial Implementation: It allows your OrdersGrpcHandler struct
	// to implement only the RPC methods you actually need, while the
	// rest fall back to returning "method not implemented" errors.
}

func NewOrderGrpcHandler(grpc *grpc.Server, orderService types.OrderService) {
	grpchandler := &OrdersGrpcHandler{
		orderService: orderService,
	}
	orders.RegisterOrderServiceServer(grpc, grpchandler)
	//  It's part of the standard gRPC service registration pattern that
	// connects your business logic to the gRPC transport layer.
}
func (h *OrdersGrpcHandler) GetOrder(ctx context.Context, req *orders.GetOrderRequest) (*orders.GetOrderResponse, error) {
	orderList := h.orderService.GetOrder(ctx)
	res := &orders.GetOrderResponse{
		Orders: orderList,
	}
	return res, nil
}
func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID:    42,
		CustomerID: 2,
		ProductID:  2,
		Quantity:   2,
	}
	err := h.orderService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}
	res := &orders.CreateOrderResponse{
		Status: "success",
	}
	return res, nil
}
