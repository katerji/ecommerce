package server

import (
	"context"
	"github.com/katerji/ecommerce/proto_out/generated"
	"github.com/katerji/ecommerce/service/user"
)

type UserServer struct {
	service *user.Service
	generated.UnimplementedUserServiceServer
}

func NewUserServer() UserServer {
	s := user.New()
	return UserServer{
		service: s,
	}
}

func (s UserServer) Login(_ context.Context, _ *generated.LoginRequest) (*generated.LoginResponse, error) {
	return nil, nil
}
func (s UserServer) Signup(_ context.Context, _ *generated.SignupRequest) (*generated.SignupResponse, error) {
	return nil, nil
}
func (s UserServer) Logout(_ context.Context, _ *generated.LogoutRequest) (*generated.LogoutResponse, error) {
	return nil, nil
}
func (s UserServer) GetAddresses(_ context.Context, _ *generated.GetAddressesRequest) (*generated.GetAddressesResponse, error) {
	return nil, nil
}
func (s UserServer) CreateAddresses(_ context.Context, _ *generated.CreateAddressRequest) (*generated.CreateAddressResponse, error) {
	return nil, nil
}
func (s UserServer) UpdateAddresses(_ context.Context, _ *generated.UpdateAddressRequest) (*generated.UpdateAddressResponse, error) {
	return nil, nil
}
func (s UserServer) DeleteAddresses(_ context.Context, _ *generated.DeleteAddressRequest) (*generated.DeleteAddressResponse, error) {
	return nil, nil
}
