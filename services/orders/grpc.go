package orders

type gRPCServer struct {
	addr string
}

func NewgRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}
func (s *gRPCServer) Start() error {
	// Implementation for starting the gRPC server
	return nil
}
