package service

import (
	"context"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/merynayr/mall-tenants/internal/model"
)

// UserService интерфейс сервисного слоя user
type UserService interface {
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetClients(ctx context.Context, limit, offset uint64) ([]model.Client, error)
	CreateClient(ctx context.Context, req model.RegisterRequest) error
}

// AuthService интерфейс сервисного слоя auth
type AuthService interface {
	Register(ctx context.Context, req model.RegisterRequest) (*model.AuthResponse, error)
	Login(ctx context.Context, email string, password string) (*model.AuthResponse, error)
	GetRefreshToken(ctx context.Context, oldRefreshToken string) (string, error)
	GetAccessToken(ctx context.Context, refreshToken string) (string, error)
}

// AccessService интерфейс сервисного слоя access
type AccessService interface {
	Check(ctx *gin.Context, endpointAddress string) (string, error)
}

// PremiseService интерфейс сервисного слоя access
type PremiseService interface {
	GetAllPremises(ctx context.Context) ([]model.Premises, error)
	GetPremisesByCode(ctx context.Context, code int64) (*model.Premises, error)
	CreatePremise(ctx context.Context, premise model.Premises) error
	UpdatePremise(ctx context.Context, premise model.Premises) error
}

// RentalService - интерфейс репо слоя для аренд
type RentalService interface {
	GetRentalByID(ctx context.Context, code int64) (*model.Rental, error)
	CreateRental(ctx context.Context, rental model.Rental) error
	UpdateRental(ctx context.Context, rental model.Rental) error
	GetAgreements(ctx context.Context, limit, offset uint64) ([]model.Agreements, error)
}

// PaymentService - интерфейс репо слоя для платежей
type PaymentService interface {
	// CreatePayment(ctx context.Context, payment model.Payment) error
	// GetPaymentByID(ctx context.Context, id int64) (*model.Payment, error)
	// GetLastPaymentByRentalID(ctx context.Context, rentalID int64) (*model.Payment, error)
	// GetOverduePayments(ctx context.Context, rentalID int64) ([]model.Payment, error)
	CreatePayment(ctx context.Context, p model.Payment) (int64, error)
	MarkAsPaid(ctx context.Context, id int64) error
	MarkPaymentsAsPaid(ctx context.Context, paymentIDs []int64) error
	GetByID(ctx context.Context, id int64) (model.Payment, error)
	ListPayments(ctx context.Context, filter model.PaymentFilter) ([]model.Payment, error)
}

// FloorPlanService интерфейс сервисного слоя access
type FloorPlanService interface {
	SaveFloorPlan(ctx context.Context, floor int64, reader io.Reader) error
	GetFloorPlanContent(ctx context.Context, floor int64) ([]byte, error)
	AddPolygon(ctx context.Context, poly *model.PremisePolygon) error
	GetPolygons(ctx context.Context, floor int64) ([]*model.Polygons, error)
	DeletPolygon(ctx context.Context, premiseCode int64) error
}
