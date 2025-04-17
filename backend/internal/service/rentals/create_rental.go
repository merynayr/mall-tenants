package rentals

import (
	"context"

	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/sys"
)

func (s *srv) CreateRental(ctx context.Context, rental model.Rental) error {
	premise, exist, err := s.premiseRepository.GetPremisesByCode(ctx, rental.SpaceID)
	if err != nil {
		return err
	}
	if !exist {
		return sys.PremiseNotFoundError
	}

	if premise.Status == string(model.Occupied) {
		return sys.PremiseOccupiedError
	} else if premise.Status == string(model.Maintenance) {
		return sys.PremiseUnderRepairError
	}

	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error

		errTx = s.rentalRepository.CreateRental(ctx, &rental)
		if errTx != nil {
			return errTx
		}

		premise := model.Premises{
			Code:   rental.SpaceID,
			Status: string(model.Occupied),
		}

		errTx = s.premiseRepository.UpdatePremise(ctx, &premise)
		if errTx != nil {
			return errTx
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
