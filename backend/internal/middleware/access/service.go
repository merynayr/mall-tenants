package access

import (
	"github.com/gin-gonic/gin"
	"github.com/merynayr/mall-tenants/internal/config"
	"github.com/merynayr/mall-tenants/internal/logger"
	"github.com/merynayr/mall-tenants/internal/service"
	"github.com/merynayr/mall-tenants/internal/sys"
)

// Middleware структура для проверки доступа
type Middleware struct {
	accessService service.AccessService
	authConfig    config.AuthConfig
}

// NewAccessMiddleware возвращает новый объект middleware слоя access
func NewAccessMiddleware(accessService service.AccessService, authConfig config.AuthConfig) *Middleware {
	return &Middleware{
		accessService: accessService,
		authConfig:    authConfig,
	}
}

// Check проверяет доступ к ресурсу
func (m *Middleware) Check() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		path := c.FullPath()
		endpoint := method + ":" + path

		user, err := m.accessService.Check(c, endpoint)
		if err != nil {
			logger.Debug(err.Error())
			sys.HandleError(c, err)
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
