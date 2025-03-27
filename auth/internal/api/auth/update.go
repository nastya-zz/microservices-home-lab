package auth

import (
	"auth/internal/converter"
	desc "auth/pkg/user_v1"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"strings"
)

func (s *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("Received user UpdateRequest %+v", req.GetUserInfo())

	errors := validateUserUpdateRequest(req)
	if len(errors) > 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			strings.Join(errors, ","))
	}

	err := s.authService.Update(ctx, converter.ToUpdateUserFromDesc(req.GetUserInfo()))
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			"Id does not exists")
	}

	return &emptypb.Empty{}, nil
}

func validateUserUpdateRequest(req *desc.UpdateRequest) []string {
	user := req.GetUserInfo()
	var errors = make([]string, 0)

	if user == nil {
		errors = append(errors, "User creation failed")
		return errors
	}
	if user.GetId() == 0 {
		errors = append(errors, "User id is empty")
	}

	return errors
}
