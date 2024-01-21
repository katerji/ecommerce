package main

import (
	"github.com/katerji/ecommerce/proto_out/user"
	"github.com/katerji/ecommerce/server"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	user.RegisterUserServiceServer(s, server.NewUserServer())
}
