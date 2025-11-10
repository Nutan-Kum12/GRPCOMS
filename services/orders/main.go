package main

func main() {
	gRPCServer := NewgRPCServer(":50051")
	if err := gRPCServer.Run(); err != nil {
		panic(err)
	}
}
