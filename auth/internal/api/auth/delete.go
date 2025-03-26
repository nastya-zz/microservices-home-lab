package auth

import (
	desc "auth/pkg/user_v1"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (s *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Received user UpdateRequest %+v", req.GetId())

	if err := s.authService.Delete(ctx, req.GetId()); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
