package auth

import (
	"auth/internal/converter"
	desc "auth/pkg/user_v1"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"strings"
)

func (s *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Received user CreateRequest %+v", req.GetUser())

	errors := validateCreateUserRequest(req)
	if len(errors) > 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			strings.Join(errors, ","))
	}
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

func validateCreateUserRequest(req *desc.CreateRequest) []string {
	user := req.GetUser()
	var errors = make([]string, 0)

	if user == nil {
		errors = append(errors, "User creation failed")
		return errors
	}
	if user.GetName() == "" {
		errors = append(errors, "User name cannot be empty")
	}
	if user.GetEmail() == "" {
		errors = append(errors, "User email cannot be empty")
	}
	if user.GetPassword() == "" {
		errors = append(errors, "User password cannot be empty")
	}

	password := user.GetPassword()
	confirmPass := user.GetPasswordConfirm()

	if password != confirmPass {
		errors = append(errors, "User passwords do not match")
	}

	return errors
}
