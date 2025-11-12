package main

func main() {
	httpServer := NewHTTPServer(":8081")
	go httpServer.Run()
	gRPCServer := NewgRPCServer(":50051")
	if err := gRPCServer.Run(); err != nil {
		panic(err)
	}
}
