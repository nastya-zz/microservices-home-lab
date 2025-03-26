package auth

import (
	"auth/internal/converter"
	desc "auth/pkg/user_v1"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Received user CreateRequest %+v", req.GetUser())

	//todo validation
	preparedUser := converter.ToCreateUserFromDesc(req.GetUser())

	id, err := s.authService.Create(ctx, preparedUser)

	log.Printf("Created user with id %d", id)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"User creation failed")

	}

	return &desc.CreateResponse{Id: id}, nil
}
