package rentals

import (
	"context"

	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/sys"
)

func (s *srv) CreateRental(ctx context.Context, rental model.Rental) error {
	flag, err := s.rentalRepository.CheckRentalOverlap(ctx, rental.SpaceID, rental.StartDate, rental.EndDate)
	if err != nil {
		return err
	}
	if !flag {
		err = s.rentalRepository.CreateRental(ctx, &rental)
		if err != nil {
			return err
		}
	} else {
		return sys.ExistError
	}
	return nil
}
