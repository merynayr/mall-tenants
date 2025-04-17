package auth

import (
	"github.com/merynayr/mall-tenants/internal/config"
	"github.com/merynayr/mall-tenants/internal/repository"
	"github.com/merynayr/mall-tenants/internal/service"
)

type srv struct {
	userRepository repository.UserRepository
	authCfg        config.AuthConfig
}

// NewService возвращает новый объект сервисного слоя mall-tenants
func NewService(userRepo repository.UserRepository, authCfg config.AuthConfig) service.AuthService {
	return &srv{
		userRepository: userRepo,
		authCfg:        authCfg,
	}
}
