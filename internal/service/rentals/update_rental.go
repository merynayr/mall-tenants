package rentals

import (
	"context"

	"github.com/merynayr/mall-tenants/internal/model"
)

func (s *srv) UpdateRental(ctx context.Context, rental model.Rental) error {
	err := s.rentalRepository.UpdateRental(ctx, &rental)
	if err != nil {
		return err
	}
	return nil
}
