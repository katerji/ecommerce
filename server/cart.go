package server

import (
	"fmt"
	"github.com/katerji/ecommerce/proto_out/generated"
	"github.com/katerji/ecommerce/service"
	"github.com/katerji/ecommerce/service/cart"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type cartMicroservice struct{}

func (c cartMicroservice) StartGRPCServer() {
	s := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor))
	generated.RegisterUserServiceServer(s, newCartGRPCServer())

	lis, err := net.Listen("tcp", ":9998")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	reflection.Register(s)

	fmt.Printf("Server is listening on %s...\n", ":9999")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

type cartGRPCServer struct {
	service *cart.Service
	generated.UnimplementedUserServiceServer
}

func newCartGRPCServer() cartGRPCServer {
	return cartGRPCServer{
		service: service.GetServiceContainerInstance().CartService,
	}
}
