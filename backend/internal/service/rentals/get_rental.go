package rentals

import (
	"context"

	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/sys"
)

func (s *srv) GetRentalByID(ctx context.Context, code int64) ([]model.RentalWithContract, error) {
	var rents []model.RentalWithContract
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		rentals, errTx := s.rentalRepository.GetRentalsByID(ctx, code)
		if errTx != nil {
			return errTx
		}
		if len(rentals) == 0 {
			return sys.RentalsNotFoundError
		}

		for _, r := range rentals {
			contracts, err := s.contractRepository.GetByRentalID(ctx, r.RentalID)
			if err != nil {
				return err
			}
			if len(contracts) == 0 {
				continue
			}

			rents = append(rents, model.RentalWithContract{
				RentalID:  r.RentalID,
				SpaceID:   r.SpaceID,
				ClientID:  r.ClientID,
				StartDate: r.StartDate,
				EndDate:   r.EndDate,
				CreatedAt: r.CreatedAt,
				Contracts: contracts,
			})
		}

		return nil
	})

	return rents, err
}
