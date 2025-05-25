package rentals

import (
	"context"

	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/sys"
)

func (s *srv) GetRentalByID(ctx context.Context, code int64) ([]model.Rental, error) {
	rentals, err := s.rentalRepository.GetRentalsByID(ctx, code)
	if err != nil {
		return nil, err
	}
	if len(rentals) == 0 {
		return nil, sys.RentalsNotFoundError
	}

	return rentals, nil
}
