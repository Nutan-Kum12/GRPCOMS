package handler

import (
	"net/http"

	"github.com/Nutan-Kum12/GRPCOMS.git/services/common/genproto/orders"
	"github.com/Nutan-Kum12/GRPCOMS.git/services/common/util"
	"github.com/Nutan-Kum12/GRPCOMS.git/services/orders/types"
)

type OrdershttpHandler struct {
	orderService types.OrderService
}

func NewOrderhttpHandler(orderService types.OrderService) *OrdershttpHandler {
	handler := &OrdershttpHandler{
		orderService: orderService,
	}
	return handler
}
func (h *OrdershttpHandler) RegisterRoute(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}
func (h *OrdershttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest
	err := util.ParseJSON(r, &req) //read the HTTP request body and converts the
	// JSON into your protobuf structs so you can access it using the auto-gen getter methods.
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	order := &orders.Order{
		OrderID:    42,
		CustomerID: req.GetCustomerID(),
		ProductID:  req.GetProductID(),
		Quantity:   req.GetQuantity(),
	}

	err = h.orderService.CreateOrder(r.Context(), order)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	res := &orders.CreateOrderResponse{Status: "success"}
	util.WriteJSON(w, http.StatusOK, res)

}
