gRPC Client Request
       ↓
[gRPC Handler] CreateOrder(ctx, *CreateOrderRequest) (*CreateOrderResponse, error)
       ↓ (converts & delegates)
[Service Layer] CreateOrder(ctx, *Order) error
       ↓ (business logic)
[Database/External Services]
       ↓
[Service Layer] returns error or nil
       ↓ (converts response)
[gRPC Handler] returns *CreateOrderResponse
       ↓
gRPC Client Response