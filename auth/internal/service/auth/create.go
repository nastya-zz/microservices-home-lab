package auth

import (
	"auth/internal/model"
	"context"
)

func (s *serv) Create(ctx context.Context, user *model.CreateUser) (int64, error) {
	id, err := s.authRepository.Create(ctx, user)
	if err != nil {
		return 0, err
	}

	return id, nil
}
