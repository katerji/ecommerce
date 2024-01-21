package service

import (
	"github.com/katerji/ecommerce/service/user"
)

type Container struct {
	UserServer *user.Service
}

func (s *Container) init() {
	userService := &user.Service{}
	userService.InitService()
	s.UserServer = userService
}

var serviceContainerInstance *Container

func GetServiceContainerInstance() *Container {
	if serviceContainerInstance == nil {
		serviceContainerInstance = &Container{}
		serviceContainerInstance.init()
	}
	return serviceContainerInstance
}
