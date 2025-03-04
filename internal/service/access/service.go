package access

import (
	"github.com/merynayr/mall-tenants/internal/config"
	"github.com/merynayr/mall-tenants/internal/service"
)

type srv struct {
	userService  service.UserService
	userAccesses map[string]struct{}
	authConfig   config.AuthConfig
}

// NewService возвращает новый объект сервисного слоя access
func NewService(userAccesses map[string]struct{}, authConfig config.AuthConfig) service.AccessService {
	return &srv{
		userAccesses: userAccesses,
		authConfig:   authConfig,
	}
}
