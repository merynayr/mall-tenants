package rentals

import (
	"context"

	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/sys"
)

func (s *srv) GetRentalByID(ctx context.Context, code int64) (*model.Rental, error) {
	rental, exist, err := s.rentalRepository.GetRentalByID(ctx, code)
	if err != nil {
		return &model.Rental{}, err
	}
	if !exist {
		return &model.Rental{}, sys.NotFoundError
	}

	return rental, nil
}
