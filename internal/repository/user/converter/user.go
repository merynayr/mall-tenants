package converter

import (
	"github.com/merynayr/mall-tenants/internal/model"
	modelRepo "github.com/merynayr/mall-tenants/internal/repository/user/model"
)

// ToUserFromRepo конвертирует модель пользователя репо слоя в
// модель сервисного слоя
func ToUserFromRepo(user *modelRepo.User) *model.User {
	if user == nil {
		return nil
	}

	return &model.User{
		UserID:   user.ID,
		Email:    user.Username,
		Password: user.Password,
		Role:     user.Role,
	}
}
