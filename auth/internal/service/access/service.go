package access

import (
	"auth/internal/repository"
	"auth/internal/service"
)

type serv struct {
	authRepository repository.AuthRepository
}

func NewService(authRepository repository.AuthRepository) service.AccessService {
	return &serv{authRepository: authRepository}
}
