package user

import (
	"github.com/merynayr/mall-tenants/internal/repository"
	"github.com/merynayr/mall-tenants/internal/service"
)

// Структура сервисного слоя с объектами репо слоя
// и транзакционного менеджера
type userService struct {
	userRepository repository.UserRepository
}

// NewService возвращает объект сервисного слоя
func NewService(userRepository repository.UserRepository) service.UserService {
	return &userService{
		userRepository: userRepository,
	}
}
