package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/merynayr/mall-tenants/internal/config"
	"github.com/merynayr/mall-tenants/internal/middleware/access"
	"github.com/merynayr/mall-tenants/internal/service"
)

// Middleware интерфейс для всех middleware
type Middleware interface {
	Access() *access.Middleware
	TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc
}

// provider структура, реализующая Middleware
type provider struct {
	accessMiddleware *access.Middleware
}

// NewMiddlewareProvider создает новый экземпляр провайдера middleware
func NewMiddlewareProvider(accessService service.AccessService, authConfig config.AuthConfig) Middleware {
	return &provider{
		accessMiddleware: access.NewAccessMiddleware(accessService, authConfig),
	}
}

// Access возвращает middleware для доступа
func (p *provider) Access() *access.Middleware {
	return p.accessMiddleware
}

// TimeoutMiddleware ограничивает выполнение обработчика по времени
func (p *provider) TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		ch := make(chan struct{})

		c.Request = c.Request.WithContext(ctx)

		go func() {
			c.Next()
			close(ch)
		}()

		select {
		case <-ch:
			return
		case <-ctx.Done():
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "request timeout"})
		}
	}
}
