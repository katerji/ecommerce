package server

import (
	"context"
	"github.com/katerji/ecommerce/proto_out/generated"
	"github.com/katerji/ecommerce/service"
	"github.com/katerji/ecommerce/service/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	for _, r := range getAnonymousRoutes() {
		if info.FullMethod == r {
			return handler(ctx, req)
		}
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	accessToken := extractAccessToken(md)
	if accessToken == "" {
		return nil, status.Errorf(codes.Unauthenticated, "access token is missing")
	}

	user, err := authenticateUser(accessToken)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "authentication failed: %v", err)
	}

	ctx = WithUser(ctx, user)

	return handler(ctx, req)
}

func extractAccessToken(md metadata.MD) string {
	values := md.Get("access_token")
	if len(values) > 0 {
		return values[0]
	}
	return ""
}

// authenticateUser performs authentication logic. Replace it with your own logic.
func authenticateUser(accessToken string) (*user.User, error) {
	user, err := service.GetServiceContainerInstance().UserServer.VerifyAccessToken(accessToken)
	if err != nil {
		return nil, err
	}
	return user, err
}

// WithUser adds the authenticated user to the context.
func WithUser(ctx context.Context, user *user.User) context.Context {
	return context.WithValue(ctx, userContextKey, user)
}

// GetUser retrieves the authenticated user from the context.
func GetUser(ctx context.Context) *user.User {
	user, _ := ctx.Value(userContextKey).(*user.User)
	return user
}

// userContextKey is a custom type to use as the key for storing the user in the context.
type userContextKeyType string

const userContextKey userContextKeyType = "user"

func getAnonymousRoutes() []string {
	return []string{
		generated.UserService_Signup_FullMethodName,
		generated.UserService_Login_FullMethodName,
	}
}
