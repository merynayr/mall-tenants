package auth

import (
	"context"
	"fmt"

	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/sys"
	"github.com/merynayr/mall-tenants/internal/utils/jwt"
)

// GetRefreshToken принимает старый refresh token и на его основе создает и возвращает новый
func (s *srv) GetRefreshToken(ctx context.Context, oldRefreshToken string) (string, error) {
	claims, err := jwt.VerifyToken(oldRefreshToken, s.authCfg.RefreshTokenSecretKey())
	if err != nil {
		return "", fmt.Errorf(sys.ErrInvalidRefreshToken)
	}

	user, exist, err := s.userRepository.GetUserByEmail(ctx, claims.Email)
	if err != nil {
		return "", err
	}
	if !exist {
		return "", sys.UserNotFoundError
	}

	userInfo := &model.AuthRequest{
		Email:    user.Email,
		Password: user.Password,
	}

	token, err := jwt.GenerateToken(userInfo, s.authCfg.RefreshTokenSecretKey(), s.authCfg.RefreshTokenExp())
	if err != nil {
		return "", err
	}

	return token, nil
}
