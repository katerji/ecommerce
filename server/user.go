package server

import (
	"context"
	usergrpc "github.com/katerji/ecommerce/proto_out/user"
	"github.com/katerji/ecommerce/service/user"
)

type UserServer struct {
	service *user.Service
	usergrpc.UnimplementedUserServiceServer
}

func NewUserServer() UserServer {
	s := user.New()
	return UserServer{
		service: s,
	}
}

func (s UserServer) Login(_ context.Context, _ *usergrpc.LoginRequest) (*usergrpc.LoginResponse, error) {
	return nil, nil
}
func (s UserServer) Signup(_ context.Context, _ *usergrpc.SignupRequest) (*usergrpc.SignupResponse, error) {
	return nil, nil
}
func (s UserServer) Logout(_ context.Context, _ *usergrpc.LogoutRequest) (*usergrpc.LogoutResponse, error) {
	return nil, nil
}
func (s UserServer) GetAddresses(_ context.Context, _ *usergrpc.GetAddressesRequest) (*usergrpc.GetAddressesResponse, error) {
	return nil, nil
}
func (s UserServer) CreateAddresses(_ context.Context, _ *usergrpc.CreateAddressRequest) (*usergrpc.CreateAddressResponse, error) {
	return nil, nil
}
func (s UserServer) UpdateAddresses(_ context.Context, _ *usergrpc.UpdateAddressRequest) (*usergrpc.UpdateAddressResponse, error) {
	return nil, nil
}
func (s UserServer) DeleteAddresses(_ context.Context, _ *usergrpc.DeleteAddressRequest) (*usergrpc.DeleteAddressResponse, error) {
	return nil, nil
}
