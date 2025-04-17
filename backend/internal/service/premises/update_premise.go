package premises

import (
	"context"

	"github.com/merynayr/mall-tenants/internal/model"
)

func (s *srv) UpdatePremise(ctx context.Context, premise model.Premises) error {
	err := s.premiseRepository.UpdatePremise(ctx, &premise)
	if err != nil {
		return err
	}
	return nil
}
