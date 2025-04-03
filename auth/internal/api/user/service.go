package user

import (
	"auth/internal/service"
	desc "auth/pkg/user_v1"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	authService service.UserService
}

func NewImplementation(authService service.UserService) *Implementation {
	return &Implementation{
		authService: authService,
	}
}
