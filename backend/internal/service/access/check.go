package access

import (
	"github.com/gin-gonic/gin"
	"github.com/merynayr/mall-tenants/internal/logger"
	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/sys"
	"github.com/merynayr/mall-tenants/internal/utils/jwt"
)

const (
	authCookieName = "access_token"
)

// Check проверяет, имеет ли пользователь доступ к эндпоинту
func (s *srv) Check(ctx *gin.Context, endpointAddress string) (string, error) {
	accessToken, err := ctx.Cookie(authCookieName)
	if err != nil {
		return "", sys.AuthHeaderNotProvidedError
	}

	claims, err := jwt.VerifyToken(accessToken, s.authConfig.AccessTokenSecretKey())
	if err != nil {
		logger.Debug(err.Error())
		return "", err
	}

	roleAccessMap, ok := s.userAccesses[model.UserRole(claims.Role)]
	if !ok {
		return "", sys.AccessDeniedError
	}

	if _, allowed := roleAccessMap[endpointAddress]; !allowed {
		return "", sys.AccessDeniedError
	}

	return claims.Email, nil
}
