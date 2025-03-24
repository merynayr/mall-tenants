package payments

import (
	"context"

	"github.com/merynayr/mall-tenants/internal/client/db"
	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/repository"
	"github.com/merynayr/mall-tenants/internal/service"
	"github.com/merynayr/mall-tenants/internal/sys"
)

type srv struct {
	paymentRepository repository.PaymentRepository
	rentalRepository  repository.RentalRepository
	premiseRepository repository.PremiseRepository
	txManager         db.TxManager
}

// NewService создаёт новый сервис платежей
func NewService(
	paymentRepo repository.PaymentRepository,
	rentalRepo repository.RentalRepository,
	premiseRepo repository.PremiseRepository,
	txManager db.TxManager,
) service.PaymentService {
	return &srv{
		paymentRepository: paymentRepo,
		rentalRepository:  rentalRepo,
		premiseRepository: premiseRepo,
		txManager:         txManager,
	}
}

// CreatePayment создаёт новый платёж
func (s *srv) CreatePayment(ctx context.Context, payment model.Payment) error {
	rental, exists, err := s.rentalRepository.GetRentalByID(ctx, payment.RentID)
	if err != nil {
		return err
	}
	if !exists {
		return sys.RentalNotFoundError
	}

	premise, exist, err := s.premiseRepository.GetPremisesByCode(ctx, rental.SpaceID)
	if err != nil {
		return err
	}
	if !exist {
		return sys.PremiseNotFoundError
	}

	if payment.Amount < premise.RentPerMonth {
		return sys.NotEnoughCoinsError
	}

	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error

		errTx = s.paymentRepository.CreatePayment(ctx, &payment)
		if errTx != nil {
			return errTx
		}

		errTx = s.rentalRepository.UpdateRental(
			ctx,
			&model.Rental{
				RentalID:   rental.RentalID,
				PaidMonths: rental.PaidMonths + payment.Amount/premise.RentPerMonth,
			},
		)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	return err
}

// GetPaymentByID получает платёж по его ID
func (s *srv) GetPaymentByID(ctx context.Context, id int64) (*model.Payment, error) {
	payment, exists, err := s.paymentRepository.GetPaymentByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, sys.PaymentNotFoundError
	}
	return payment, nil
}

// GetLastPaymentByRentalID получает последний платёж по ID аренды
func (s *srv) GetLastPaymentByRentalID(ctx context.Context, rentalID int64) (*model.Payment, error) {
	// Проверяем, существует ли аренда
	_, exists, err := s.rentalRepository.GetRentalByID(ctx, rentalID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, sys.RentalNotFoundError
	}

	// Получаем последний платёж
	payment, found, err := s.paymentRepository.GetLastPayment(ctx, rentalID)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, sys.PaymentNotFoundError
	}

	return payment, nil
}

// GetOverduePayments получает список просроченных платежей
func (s *srv) GetOverduePayments(ctx context.Context, rentalID int64) ([]model.Payment, error) {
	payments, err := s.paymentRepository.GetOverduePayments(ctx, rentalID)
	if err != nil {
		return nil, err
	}
	if len(payments) == 0 {
		return nil, sys.OverduePaymentsNotFoundError
	}
	return payments, nil
}
