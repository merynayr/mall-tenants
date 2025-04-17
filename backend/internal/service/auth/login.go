package auth

import (
	"context"

	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/sys"
	"github.com/merynayr/mall-tenants/internal/utils/hash"
	"github.com/merynayr/mall-tenants/internal/utils/jwt"
)

// Login валидирует данные пользователя, и если все ок, возвращает token-ы
func (s *srv) Login(ctx context.Context, email string, password string) (*model.AuthResponse, error) {
	var userInfo *model.AuthRequest
	user, exist, err := s.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, sys.UserNotFoundError
	}

	userInfo = &model.AuthRequest{
		Email:    user.Email,
		Password: user.Password,
	}

	err = hash.CompareHashAndPass(password, userInfo.Password)
	if err != nil {
		return nil, sys.InvalidPasswordError
	}

	refreshToken, err := jwt.GenerateToken(userInfo, s.authCfg.RefreshTokenSecretKey(), s.authCfg.RefreshTokenExp())
	if err != nil {
		return nil, err
	}
	accessToken, err := jwt.GenerateToken(userInfo, s.authCfg.AccessTokenSecretKey(), s.authCfg.AccessTokenExp())
	if err != nil {
		return nil, err
	}

	return &model.AuthResponse{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}, nil
}
