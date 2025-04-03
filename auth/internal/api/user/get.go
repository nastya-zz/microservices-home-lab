package user

import (
	"auth/internal/converter"
	desc "auth/pkg/user_v1"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Received user GetRequest %+v", req.GetId())

	if req.GetId() == 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"user id is required")
	}

	user, err := s.authService.Get(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"User get failed")
	}

	log.Printf("id: %d, username: %s, created_at: %v, updated_at: %v\n", user.ID, user.Name, user.CreatedAt, user.UpdatedAt)

	return &desc.GetResponse{
		User: converter.ToUserFromService(user),
	}, nil
}
