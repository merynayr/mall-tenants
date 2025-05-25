package access

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/sys"
	"github.com/merynayr/mall-tenants/internal/utils/jwt"
)

const (
	authCookieName = "access_token"
	authHeader     = "Authorization"
	authPrefix     = "Bearer "
)

// Check проверяет, имеет ли пользователь доступ к эндпоинту
func (s *srv) Check(ctx *gin.Context, endpointAddress string) (string, error) {
	isProtected := false
	for _, accessMap := range s.userAccesses {
		if _, ok := accessMap[endpointAddress]; ok {
			isProtected = true
			break
		}
	}
	if !isProtected {
		// Не защищённый путь — доступ разрешён без проверки токена
		return "", nil
	}

	accessToken, err := ctx.Cookie(authCookieName)
	if err != nil {
		authHeader := ctx.GetHeader(authHeader)
		if authHeader == "" {
			return "", sys.AuthHeaderMissingError
		}

		if !strings.HasPrefix(authHeader, authPrefix) {
			return "", sys.AuthHeaderInvalidFormatError
		}

		accessToken = strings.TrimPrefix(authHeader, authPrefix)
		if len(accessToken) == 0 {
			return "", sys.AuthHeaderInvalidFormatError
		}
	}

	claims, err := jwt.VerifyToken(accessToken, s.authConfig.AccessTokenSecretKey())
	if err != nil {
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
