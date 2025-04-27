package access

import (
	"github.com/merynayr/mall-tenants/internal/config"
	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/service"
)

type srv struct {
	userService  service.UserService
	userAccesses map[model.UserRole]map[string]struct{}
	authConfig   config.AuthConfig
}

// NewService возвращает новый объект сервисного слоя access
func NewService(
	userService service.UserService,
	userAccesses map[model.UserRole]map[string]struct{},
	authConfig config.AuthConfig,
) service.AccessService {
	return &srv{
		userService:  userService,
		userAccesses: userAccesses,
		authConfig:   authConfig,
	}
}
