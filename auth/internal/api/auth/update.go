package auth

import (
	"auth/internal/converter"
	desc "auth/pkg/user_v1"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (s *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("Received user UpdateRequest %+v", req.GetUserInfo())

	err := s.authService.Update(ctx, converter.ToUpdateUserFromDesc(req.GetUserInfo()))
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			"Id does not exists")
	}

	return &emptypb.Empty{}, nil
}
