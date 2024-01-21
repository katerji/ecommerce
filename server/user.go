package server

import (
	"context"
	"fmt"
	"github.com/katerji/ecommerce/proto_out/generated"
	"github.com/katerji/ecommerce/service"
	"github.com/katerji/ecommerce/service/user"
)

type UserServer struct {
	service *user.Service
	generated.UnimplementedUserServiceServer
}

func NewUserServer() UserServer {
	return UserServer{
		service: service.GetServiceContainerInstance().UserServer,
	}
}

func (s UserServer) Login(_ context.Context, request *generated.LoginRequest) (*generated.LoginResponse, error) {
	if request.Email == "" && request.PhoneNumber == "" {
		return &generated.LoginResponse{
			Success: false,
			Message: "email or phone required",
		}, nil
	}
	if request.Email != "" {
		result, err := s.service.LoginWithEmail(request.Email, request.Password)
		if err != nil {
			return &generated.LoginResponse{
				Success: false,
				Message: err.Error(),
			}, err
		}
		return &generated.LoginResponse{
			Success:          true,
			User:             userToProto(result.User),
			AccessToken:      result.JWTPair.AccessToken,
			ExpiresAt:        result.JWTPair.ExpiresAt,
			RefreshToken:     result.JWTPair.RefreshToken,
			RefreshExpiresAt: result.JWTPair.RefreshExpiresAt,
		}, nil
	}
	result, err := s.service.LoginWithPhoneNumber(request.PhoneNumber, request.Password)
	if err != nil {
		return &generated.LoginResponse{
			Success: false,
			Message: err.Error(),
		}, err
	}
	return &generated.LoginResponse{
		Success:          true,
		User:             userToProto(result.User),
		AccessToken:      result.JWTPair.AccessToken,
		ExpiresAt:        result.JWTPair.ExpiresAt,
		RefreshToken:     result.JWTPair.RefreshToken,
		RefreshExpiresAt: result.JWTPair.RefreshExpiresAt,
	}, nil
}

func (s UserServer) Signup(_ context.Context, request *generated.SignupRequest) (*generated.SignupResponse, error) {
	isOneSet := request.Email != "" || request.PhoneNumber != ""
	if !isOneSet || request.Name == "" || request.Password == "" {
		return &generated.SignupResponse{
			Message: "missing param",
			Success: false,
		}, nil
	}
	user := &user.User{
		Email:       request.Email,
		Name:        request.Name,
		PhoneNumber: request.PhoneNumber,
	}
	result, err := s.service.Signup(user, request.Password)
	if err != nil {
		return &generated.SignupResponse{
			Message: err.Error(),
			Success: false,
		}, nil
	}
	return &generated.SignupResponse{
		Success:          true,
		User:             userToProto(result.User),
		AccessToken:      result.JWTPair.AccessToken,
		ExpiresAt:        result.JWTPair.ExpiresAt,
		RefreshToken:     result.JWTPair.RefreshToken,
		RefreshExpiresAt: result.JWTPair.RefreshExpiresAt,
	}, nil
}

func (s UserServer) Logout(_ context.Context, _ *generated.LogoutRequest) (*generated.LogoutResponse, error) {
	return nil, nil
}

func (s UserServer) GetAddresses(ctx context.Context, _ *generated.GetAddressesRequest) (*generated.GetAddressesResponse, error) {
	user := GetUser(ctx)
	fmt.Println(user)
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

func userToProto(user *user.User) *generated.User {
	return &generated.User{
		Id:          int64(user.ID),
		Email:       user.Email,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
	}
}
