package server

import (
	"context"
	"fmt"
	"github.com/katerji/ecommerce/proto_out/generated"
	"github.com/katerji/ecommerce/service"
	"github.com/katerji/ecommerce/service/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type userMicroservice struct{}

func (u userMicroservice) StartGRPCServer() {
	s := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor))
	generated.RegisterUserServiceServer(s, newUserGRPCServer())

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

type userGRPCServer struct {
	service *user.Service
	generated.UnimplementedUserServiceServer
}

func newUserGRPCServer() userGRPCServer {
	return userGRPCServer{
		service: service.GetServiceContainerInstance().UserService,
	}
}

func (s userGRPCServer) Login(_ context.Context, request *generated.LoginRequest) (*generated.LoginResponse, error) {
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

func (s userGRPCServer) Signup(_ context.Context, request *generated.SignupRequest) (*generated.SignupResponse, error) {
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

func (s userGRPCServer) Logout(_ context.Context, _ *generated.LogoutRequest) (*generated.LogoutResponse, error) {
	return nil, nil
}

func (s userGRPCServer) GetAddresses(ctx context.Context, _ *generated.GetAddressesRequest) (*generated.GetAddressesResponse, error) {
	u := GetUser(ctx)
	addressesMap, err := s.service.GetAddresses(u.ID)
	if err != nil {
		return &generated.GetAddressesResponse{
			ResponseStatus: &generated.ResponseStatus{
				Success: false,
				Message: "something went wrong",
			},
		}, nil
	}
	addresses := []*generated.Address{}
	for _, a := range addressesMap {
		addresses = append(addresses, addressToProto(&a))
	}

	return &generated.GetAddressesResponse{
		Addresses: addresses,
		ResponseStatus: &generated.ResponseStatus{
			Success: true,
		},
	}, nil
}

func (s userGRPCServer) CreateAddresses(ctx context.Context, request *generated.CreateAddressRequest) (*generated.CreateAddressResponse, error) {
	u := GetUser(ctx)
	address := &user.Address{
		UserID:       u.ID,
		AddressLine1: request.Address.AddressLine_1,
		AddressLine2: request.Address.AddressLine_2,
		City:         request.Address.City,
		State:        request.Address.State,
		ZipCode:      request.Address.ZipCode,
		Country:      request.Address.Country,
	}
	if address.AddressLine1 == "" || address.City == "" || address.Country == "" {
		return &generated.CreateAddressResponse{
			ResponseStatus: &generated.ResponseStatus{
				Success: false,
				Message: "missing params",
			},
		}, nil
	}
	insertedAddress, err := s.service.CreateAddress(address)
	if err != nil {
		return &generated.CreateAddressResponse{
			ResponseStatus: &generated.ResponseStatus{
				Success: false,
				Message: "something went wrong",
			},
		}, nil
	}
	return &generated.CreateAddressResponse{
		Address: addressToProto(insertedAddress),
		ResponseStatus: &generated.ResponseStatus{
			Success: true,
		},
	}, nil
}
func (s userGRPCServer) UpdateAddresses(ctx context.Context, request *generated.UpdateAddressRequest) (*generated.UpdateAddressResponse, error) {
	u := GetUser(ctx)
	address := &user.Address{
		ID:           int(request.Address.Id),
		UserID:       u.ID,
		AddressLine1: request.Address.AddressLine_1,
		AddressLine2: request.Address.AddressLine_2,
		City:         request.Address.City,
		State:        request.Address.State,
		ZipCode:      request.Address.ZipCode,
		Country:      request.Address.Country,
	}
	if address.AddressLine1 == "" || address.City == "" || address.Country == "" {
		return &generated.UpdateAddressResponse{
			ResponseStatus: &generated.ResponseStatus{
				Success: false,
				Message: "missing params",
			},
		}, nil
	}
	err := s.service.UpdateAddress(address)
	if err != nil {
		return &generated.UpdateAddressResponse{
			ResponseStatus: &generated.ResponseStatus{
				Success: false,
				Message: "something went wrong",
			},
		}, nil
	}
	return &generated.UpdateAddressResponse{
		Address: addressToProto(address),
		ResponseStatus: &generated.ResponseStatus{
			Success: true,
		},
	}, nil
}
func (s userGRPCServer) DeleteAddresses(ctx context.Context, request *generated.DeleteAddressRequest) (*generated.DeleteAddressResponse, error) {
	if request.AddressId == 0 {
		return &generated.DeleteAddressResponse{
			ResponseStatus: &generated.ResponseStatus{
				Success: false,
				Message: "missing address ID",
			},
		}, nil
	}
	if err := s.service.DeleteAddress(int(request.AddressId)); err != nil {
		return &generated.DeleteAddressResponse{
			ResponseStatus: &generated.ResponseStatus{
				Success: false,
				Message: "something went wrong",
			},
		}, nil
	}
	return &generated.DeleteAddressResponse{
		ResponseStatus: &generated.ResponseStatus{
			Success: true,
		},
	}, nil
}

func userToProto(user *user.User) *generated.User {
	return &generated.User{
		Id:          int64(user.ID),
		Email:       user.Email,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
	}
}

func addressToProto(address *user.Address) *generated.Address {
	return &generated.Address{
		Id:            int64(address.ID),
		UserId:        int64(address.UserID),
		AddressLine_1: address.AddressLine1,
		AddressLine_2: address.AddressLine2,
		Country:       address.Country,
		City:          address.City,
		State:         address.State,
		ZipCode:       address.ZipCode,
	}
}
