package user

import (
	"auth/internal/model"
	"context"
)

func (s *serv) Update(ctx context.Context, updateUser *model.UpdateUser) error {
	if err := s.authRepository.Update(ctx, updateUser); err != nil {
		return err
	}

	return nil
}
