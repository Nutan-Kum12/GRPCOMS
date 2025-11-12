package main

import (
	"log"
	"net/http"

	handler "github.com/Nutan-Kum12/GRPCOMS.git/services/orders/handler/orders"
	"github.com/Nutan-Kum12/GRPCOMS.git/services/orders/services"
)

type httpServer struct {
	addr string
}

func NewHTTPServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}
func (s *httpServer) Run() error {
	// Implement HTTP server logic here
	router := http.NewServeMux()
	orderService := services.NewOrdersService()
	orderHandler := handler.NewOrderhttpHandler(orderService)
	orderHandler.RegisterRoute(router)
	log.Println("server starting on ", s.addr)
	return http.ListenAndServe(s.addr, router)
}
