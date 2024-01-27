package service

import (
	"github.com/katerji/ecommerce/service/cart"
	"github.com/katerji/ecommerce/service/user"
)

type Container struct {
	UserService *user.Service
	CartService *cart.Service
}

func (s *Container) init() {
	s.UserService = user.New()
	s.CartService = cart.New()
}

var serviceContainerInstance *Container

func GetServiceContainerInstance() *Container {
	if serviceContainerInstance == nil {
		serviceContainerInstance = &Container{}
		serviceContainerInstance.init()
	}
	return serviceContainerInstance
}
