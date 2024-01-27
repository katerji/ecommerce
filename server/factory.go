package server

import (
	"github.com/katerji/ecommerce/envs"
)

const serverUser = "user"
const serverCart = "cart"

type IServer interface {
	StartGRPCServer()
}

type Factory struct{}

func (f Factory) GetServer() IServer {
	switch envs.GetInstance().Server() {
	case serverUser:
		return userMicroservice{}
	case serverCart:
		return cartMicroservice{}
	}
	return userMicroservice{}
}
