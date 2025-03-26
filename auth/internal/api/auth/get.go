package auth

import (
	"auth/internal/converter"
	desc "auth/pkg/user_v1"
	"context"
	"log"
)

func (s *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Received user GetRequest %+v", req.GetId())

	user, err := s.authService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("id: %d, username: %s, created_at: %v, updated_at: %v\n", user.ID, user.Name, user.CreatedAt, user.UpdatedAt)

	return &desc.GetResponse{
		User: converter.ToUserFromService(user),
	}, nil
}
