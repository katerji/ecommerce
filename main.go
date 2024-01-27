package main

import "github.com/katerji/ecommerce/server"

func main() {
	factory := server.Factory{}
	s := factory.GetServer()
	s.StartGRPCServer()
}
