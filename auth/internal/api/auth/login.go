package auth

import (
	desc "auth/pkg/auth_v1"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Login(ctx context.Context, req *desc.LoginRequest) (*desc.LoginResponse, error) {

	email := req.GetLogin()
	password := req.GetPassword()

	token, err := i.authService.Login(ctx, email, password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "method Login")
	}

	return &desc.LoginResponse{
		RefreshToken: token,
	}, nil
}
