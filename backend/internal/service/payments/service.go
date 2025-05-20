package payments

import (
	"context"
	"errors"
	"time"

	"github.com/merynayr/mall-tenants/internal/client/db"
	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/repository"
	"github.com/merynayr/mall-tenants/internal/service"
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

func (s *srv) CreatePayment(ctx context.Context, p model.Payment) (int64, error) {
	// можно добавить валидацию, например:
	if p.Amount <= 0 {
		return 0, errors.New("amount must be greater than 0")
	}

	return s.paymentRepository.CreatePayment(ctx, p)
}

func (s *srv) MarkAsPaid(ctx context.Context, id int64) error {
	return s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		paidAt := time.Now().UTC()
		return s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
			err := s.paymentRepository.MarkAsPaid(ctx, id, paidAt)
			if err != nil {
				return err
			}

			payment, err := s.paymentRepository.GetByID(ctx, id)
			if err != nil {
				return err
			}

			rental := &model.Rental{
				RentalID:   payment.RentalID,
				PaidMonths: int64(1),
			}

			err = s.rentalRepository.UpdateRental(ctx, rental)
			if err != nil {
				return err
			}

			return nil
		})
	})
}

func (s *srv) MarkPaymentsAsPaid(ctx context.Context, paymentIDs []int64) error {
	if len(paymentIDs) == 0 {
		return nil
	}

	return s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		err := s.paymentRepository.MarkAsPaidMany(ctx, paymentIDs, time.Now().UTC())
		if err != nil {
			return err
		}

		payment, err := s.paymentRepository.GetByID(ctx, paymentIDs[0])
		if err != nil {
			return err
		}

		rental := &model.Rental{
			RentalID:   payment.RentalID,
			PaidMonths: int64(len(paymentIDs)),
		}

		err = s.rentalRepository.UpdateRental(ctx, rental)
		if err != nil {
			return err
		}

		return nil
	})
}

func (s *srv) GetByID(ctx context.Context, id int64) (model.Payment, error) {
	return s.paymentRepository.GetByID(ctx, id)
}

func (s *srv) ListPayments(ctx context.Context, filter model.PaymentFilter) ([]model.Payment, error) {
	return s.paymentRepository.List(ctx, filter)
}
