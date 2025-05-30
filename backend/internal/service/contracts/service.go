package contracts

import (
	"context"
	"fmt"
	"time"

	"github.com/merynayr/mall-tenants/internal/client/db"
	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/repository"
	"github.com/merynayr/mall-tenants/internal/service"
	"github.com/merynayr/mall-tenants/internal/sys"
	"github.com/merynayr/mall-tenants/internal/utils/crypto"
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
) service.ContractService {
	return &srv{
		rentalRepository:   rentalRepo,
		premiseRepository:  premiseRepo,
		paymentRepository:  paymentRepository,
		contractRepository: contractRepository,
		txManager:          txManager,
	}
}

func (s *srv) GetByContractID(ctx context.Context, contractID int64) (*model.Contracts, error) {
	contract, err := s.contractRepository.GetByContractID(ctx, contractID)
	return contract, err
}

func (s *srv) SignContract(ctx context.Context, contractID int64, rent model.RentalWithContract) error {
	contract, err := s.contractRepository.GetByContractID(ctx, contractID)
	if err != nil || contract.IsSigned {
		return fmt.Errorf("contract not found or already signed")
	}

	signature, err := crypto.SignFileHash(contract.FilePath, "C:/keys/private_key.pem")
	if err != nil {
		return err
	}

	publicKey, err := crypto.LoadPublicKeyPEM("C:/keys/public_key.pem")
	if err != nil {
		return err
	}

	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		premise, exist, errTx := s.premiseRepository.GetPremisesByCode(ctx, rent.SpaceID)
		if errTx != nil {
			return errTx
		}
		if !exist {
			return sys.PremiseNotFoundError
		}

		if errTx := s.contractRepository.UpdateSignature(ctx, contractID, signature, publicKey); errTx != nil {
			return errTx
		}

		updatePremise := model.Premises{
			Code:   rent.SpaceID,
			Status: string(model.Occupied),
		}

		if errTx := s.premiseRepository.UpdatePremise(ctx, &updatePremise); errTx != nil {
			return errTx
		}

		start := rent.StartDate
		end := rent.EndDate
		payments := make([]model.Payment, 0)

		currentStart := start

		for currentStart.Before(end) {
			nextMonth := currentStart.AddDate(0, 1, 0)

			var currentEnd time.Time
			if nextMonth.After(end) {
				currentEnd = end
			} else {
				currentEnd = nextMonth.AddDate(0, 0, -1)
			}

			payment := model.Payment{
				RentalID:    rent.RentalID,
				PeriodStart: currentStart,
				PeriodEnd:   currentEnd,
				Amount:      int(premise.RentPerMonth),
				IsPaid:      false,
				CreatedAt:   time.Now().UTC(),
			}
			payments = append(payments, payment)

			currentStart = nextMonth
		}

		if errTx := s.paymentRepository.CreatePayment(ctx, payments); errTx != nil {
			return errTx
		}

		return nil
	})
	return err
}

func (s *srv) DeleteContract(ctx context.Context, contractID int64) error {
	err := s.contractRepository.DeleteContract(ctx, contractID)
	return err
}
