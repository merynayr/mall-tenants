package access

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/merynayr/mall-tenants/internal/logger"
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
		return nil, sys.AuthHeaderNotProvidedError
	}

	if !strings.HasPrefix(authHeader, authPrefix) {
		return nil, sys.InvalidAuthHeaderFormatError
	}

	accessToken := strings.TrimPrefix(authHeader, authPrefix)

	claims, err := jwt.VerifyToken(accessToken, s.authConfig.AccessTokenSecretKey())
	if err != nil {
		logger.Debug(err.Error())
		return nil, err
	}

	user, err := s.userService.GetUserByEmail(ctx.Request.Context(), claims.Email)
	if err != nil {
		return nil, err
	}

	roleAccessMap, ok := s.userAccesses[model.UserRole(claims.Role)]
	if !ok {
		return nil, sys.AccessDeniedError
	}

	if _, allowed := roleAccessMap[endpointAddress]; !allowed {
		return nil, sys.AccessDeniedError
	}

	return user, nil
}
