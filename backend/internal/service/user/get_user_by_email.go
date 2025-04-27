package user

import (
	"context"

	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/sys"
)

func (s *userService) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user, exists, err := s.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, sys.UserNotFoundError
	}

	return user, nil
}
