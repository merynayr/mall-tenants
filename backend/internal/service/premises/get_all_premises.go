package premises

import (
	"context"

	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/sys"
)

func (s *srv) GetAllPremises(ctx context.Context) ([]model.Premises, error) {
	premises, err := s.premiseRepository.GetAllPremises(ctx)
	if err != nil {
		return nil, err
	}
	if len(premises) == 0 {
		return nil, sys.PremiseNotFoundError
	}

	return premises, nil
}
