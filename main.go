package main

import (
	"fmt"
	"github.com/katerji/ecommerce/proto_out/generated"
	"github.com/katerji/ecommerce/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()

	generated.RegisterUserServiceServer(s, server.NewUserServer())

	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	reflection.Register(s)

	fmt.Printf("Server is listening on %s...\n", ":9999")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
