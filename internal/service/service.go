package service

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/merynayr/mall-tenants/internal/model"
)

// UserService интерфейс сервисного слоя user
type UserService interface {
	GetUserByEmail(ctx context.Context, name string) (*model.User, error)
}

// AuthService интерфейс сервисного слоя auth
type AuthService interface {
	Register(ctx context.Context, req model.RegisterRequest) (*model.AuthResponse, error)
	Login(ctx context.Context, username string, password string) (*model.AuthResponse, error)
	GetRefreshToken(ctx context.Context, oldRefreshToken string) (string, error)
	GetAccessToken(ctx context.Context, refreshToken string) (string, error)
}

// AccessService интерфейс сервисного слоя access
type AccessService interface {
	Check(ctx *gin.Context, endpointAddress string) (*model.User, error)
}

// PremiseService интерфейс сервисного слоя access
type PremiseService interface {
	GetPremisesByCode(ctx context.Context, code int64) (*model.Premises, error)
	CreatePremise(ctx context.Context, premise model.Premises) error
	UpdatePremise(ctx context.Context, premise model.Premises) error
}

// RentalService - интерфейс репо слоя для аренд
type RentalService interface {
	GetRentalByID(ctx context.Context, code int64) (*model.Rental, error)
	CreateRental(ctx context.Context, rental model.Rental) error
	UpdateRental(ctx context.Context, rental model.Rental) error
}
