package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merynayr/mall-tenants/internal/config"
	"github.com/merynayr/mall-tenants/internal/service"
)

// API auth структура
type API struct {
	authService service.AuthService
	authConfig  config.AuthConfig
}

// NewAPI возвращает новый объект имплементации API-слоя auth
func NewAPI(authService service.AuthService, authConfig config.AuthConfig) *API {
	return &API{
		authService: authService,
		authConfig:  authConfig,
	}
}

// RegisterRoutes регистрирует маршруты
func (api *API) RegisterRoutes(router *gin.Engine) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", api.Register)
		authGroup.POST("/login", api.Login)
		authGroup.POST("/refresh", api.GetAccessToken)
		authGroup.POST("/refresh-token", api.GetRefreshToken)
	}
}

// setCookies устанавливают токены в куки
func (api *API) setCookies(c *gin.Context, refreshToken, accessToken string) {
	if accessToken != "" {
		cookie := &http.Cookie{
			Name:     "access_token",
			Value:    accessToken,
			Path:     "/",
			MaxAge:   int(api.authConfig.AccessTokenExp() * 2),
			HttpOnly: true,
			Secure:   false,
		}
		http.SetCookie(c.Writer, cookie)
	}
	if refreshToken != "" {
		cookie := &http.Cookie{
			Name:     "refresh_token",
			Value:    refreshToken,
			Path:     "/",
			MaxAge:   int(api.authConfig.RefreshTokenExp() * 2),
			HttpOnly: true,
			Secure:   false,
		}
		http.SetCookie(c.Writer, cookie)
	}
}
