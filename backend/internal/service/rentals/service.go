package rentals

import (
	"context"

	"github.com/merynayr/mall-tenants/internal/client/db"
	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/repository"
	"github.com/merynayr/mall-tenants/internal/service"
	"github.com/merynayr/mall-tenants/internal/sys"
)

type srv struct {
	rentalRepository  repository.RentalRepository
	premiseRepository repository.PremiseRepository
	paymentRepository repository.PaymentRepository
	txManager         db.TxManager
}

// NewService возвращает новый объект сервисного слоя mall-tenants
func NewService(
	rentalRepo repository.RentalRepository,
	premiseRepo repository.PremiseRepository,
	paymentRepository repository.PaymentRepository,
	txManager db.TxManager,
) service.RentalService {
	return &srv{
		rentalRepository:  rentalRepo,
		premiseRepository: premiseRepo,
		paymentRepository: paymentRepository,
		txManager:         txManager,
	}
}

func (s *srv) GetAgreements(ctx context.Context, limit, offset uint64) ([]model.Agreements, error) {
	if limit <= 0 {
		limit = 20
	}

	rental, err := s.rentalRepository.GetAgreements(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	if len(rental) == 0 {
		return nil, sys.RentalsNotFoundError
	}

	return rental, nil
}
