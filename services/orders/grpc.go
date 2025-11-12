package main

import (
	"log"
	"net"

	handler "github.com/Nutan-Kum12/GRPCOMS.git/services/orders/handler/orders"
	"github.com/Nutan-Kum12/GRPCOMS.git/services/orders/services"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewgRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}
func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr) //
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	// register our grpc services
	orderService := services.NewOrdersService()
	handler.NewOrderGrpcHandler(grpcServer, orderService)

	log.Println("Starting gRPC server on", s.addr)
	return grpcServer.Serve(lis)
}

// When Run() executes:
// 1. Opens port 8080
// 2. Creates gRPC server
// 3. Registers OrderService
// 4. Starts serving requests
// 5. Waits forever for clients to connect
