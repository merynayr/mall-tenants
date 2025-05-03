package auth

import (
	"context"

	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/sys"
	"github.com/merynayr/mall-tenants/internal/utils/jwt"
)

// GetAccessToken принимает refresh token и на его основе создает и возвращает access token
func (s *srv) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	claims, err := jwt.VerifyToken(refreshToken, s.authCfg.RefreshTokenSecretKey())
	if err != nil {
		return "", err
	}

	user, exist, err := s.userRepository.GetUserByEmail(ctx, claims.Email)
	if err != nil {
		return "", err
	}
	if !exist {
		return "", sys.UserNotFoundError
	}

	userInfo := &model.UserClaims{
		Email: user.Email,
		Role:  model.UserRole(user.Role),
	}

	token, err := jwt.GenerateToken(userInfo, s.authCfg.AccessTokenSecretKey(), s.authCfg.AccessTokenExp())
	if err != nil {
		return "", err
	}

	return token, nil
}
