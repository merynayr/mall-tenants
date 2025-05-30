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
	rentalRepository   repository.RentalRepository
	premiseRepository  repository.PremiseRepository
	paymentRepository  repository.PaymentRepository
	contractRepository repository.ContractRepository
	txManager          db.TxManager
}

// NewService возвращает новый объект сервисного слоя mall-tenants
func NewService(
	rentalRepo repository.RentalRepository,
	premiseRepo repository.PremiseRepository,
	paymentRepository repository.PaymentRepository,
	contractRepository repository.ContractRepository,
	txManager db.TxManager,
) service.RentalService {
	return &srv{
		rentalRepository:   rentalRepo,
		premiseRepository:  premiseRepo,
		paymentRepository:  paymentRepository,
		contractRepository: contractRepository,
		txManager:          txManager,
	}
}

func (s *srv) GetAgreements(ctx context.Context, filter model.RentFilter) ([]model.Agreements, error) {
	var rentals []model.Agreements
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error

		rentals, errTx = s.rentalRepository.GetAgreements(ctx, filter)
		if errTx != nil {
			return errTx
		}
		if len(rentals) == 0 {
			return sys.RentalsNotFoundError
		}

		for i := range rentals {
			contracts, err := s.contractRepository.GetByRentalID(ctx, rentals[i].RentalID)
			if err != nil {
				return err
			}

			rentals[i].Contracts = contracts
		}

		return nil
	})

	return rentals, err
}
