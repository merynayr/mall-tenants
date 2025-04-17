package repository

import (
	"context"
	"time"

	"github.com/merynayr/mall-tenants/internal/model"
)

// UserRepository - интерфейс репо слоя user
type UserRepository interface {
	CreateUser(ctx context.Context, req *model.RegisterRequest) (int64, error)
	UpdateUser(ctx context.Context, user *model.UserUpdate) error
	GetUserByEmail(ctx context.Context, email string) (*model.User, bool, error)
	IsEmailExist(ctx context.Context, email string) (bool, error)
}

// PremiseRepository - интерфейс репо слоя для помещений
type PremiseRepository interface {
	GetPremisesByCode(ctx context.Context, code int64) (*model.Premises, bool, error)
	CreatePremise(ctx context.Context, premise *model.Premises) error
	UpdatePremise(ctx context.Context, premise *model.Premises) error
}

// RentalRepository - интерфейс репо слоя для аренд
type RentalRepository interface {
	GetRentalByID(ctx context.Context, id int64) (*model.Rental, bool, error)
	CreateRental(ctx context.Context, rental *model.Rental) error
	UpdateRental(ctx context.Context, rental *model.Rental) error
}

// PaymentRepository - интерфейс репо слоя для платежей
type PaymentRepository interface {
	CreatePayment(ctx context.Context, payment *model.Payment) error
	GetPaymentByID(ctx context.Context, id int64) (*model.Payment, bool, error)
	GetLastPayment(ctx context.Context, rentID int64) (*model.Payment, bool, error)
	GetOverduePayments(ctx context.Context, rentID int64) ([]model.Payment, error)
	GetPaymentByRentIDAndPeriod(ctx context.Context, rentID int64, startDate, endDate time.Time) (bool, error)
}
