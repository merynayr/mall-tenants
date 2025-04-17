package auth

import (
	"context"

	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/sys"
	"github.com/merynayr/mall-tenants/internal/utils/jwt"
)

// Login валидирует данные пользователя, и если все ок, возвращает token-ы
func (s *srv) Register(ctx context.Context, req model.RegisterRequest) (*model.AuthResponse, error) {
	if req.Password != req.ConfirmPassword {
		return nil, sys.PasswordsDoNotMatchError
	}

	var userInfo *model.AuthRequest
	_, exist, err := s.userRepository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, sys.UserExistError
	}

	_, err = s.userRepository.CreateUser(ctx, &req)
	if err != nil {
		return nil, err
	}
	userInfo = &model.AuthRequest{
		Email:    req.Email,
		Password: req.Password,
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
