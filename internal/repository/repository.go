package repository

import (
	"context"

	"github.com/merynayr/mall-tenants/internal/model"
)

// UserRepository - интерфейс репо слоя user
type UserRepository interface {
	CreateUser(ctx context.Context, req *model.RegisterRequest) (int64, error)
	UpdateUser(ctx context.Context, user *model.UserUpdate) error
	GetUserByEmail(ctx context.Context, email string) (*model.User, bool, error)
	IsEmailExist(ctx context.Context, email string) (bool, error)
}
