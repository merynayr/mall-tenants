package access

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/sys"
	"github.com/merynayr/mall-tenants/internal/utils/jwt"
)

const (
	authHeader = "Authorization"
	authPrefix = "Bearer "
)

// Check проверяет, имеет ли пользователь доступ к эндпоинту
func (s *srv) Check(ctx *gin.Context, endpointAddress string) (*model.User, error) {

	authHeader := ctx.GetHeader(authHeader)
	if authHeader == "" {
		return nil, errors.New(sys.ErrAuthHeaderNotProvided)
	}

	if !strings.HasPrefix(authHeader, authPrefix) {
		return nil, errors.New(sys.ErrInvalidAuthHeaderFormat)
	}

	accessToken := strings.TrimPrefix(authHeader, authPrefix)

	claims, err := jwt.VerifyToken(accessToken, s.authConfig.AccessTokenSecretKey())
	if err != nil {
		return nil, errors.New(sys.ErrInvalidAccessToken)
	}

	user, err := s.userService.GetUserByEmail(ctx.Request.Context(), claims.Email)
	if err != nil {
		return nil, err
	}

	// Супер админ имеет доступ ко всем эндпоинтам
	if claims.Role == 2 {
		return user, nil
	}

	// Моедратор имеет доступ ко всем эндпоинтам
	if claims.Role == 1 {
		return user, nil
	}

	// смотрим, есть ли доступ у пользователя
	if _, ok := s.userAccesses[endpointAddress]; !ok {
		return nil, sys.AccessDeniedError
	}

	return user, nil
}
