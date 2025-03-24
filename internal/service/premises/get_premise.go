package premises

import (
	"context"

	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/sys"
)

func (s *srv) GetPremisesByCode(ctx context.Context, code int64) (*model.Premises, error) {
	premise, exist, err := s.premiseRepository.GetPremisesByCode(ctx, code)
	if err != nil {
		return &model.Premises{}, err
	}
	if !exist {
		return &model.Premises{}, sys.PremiseNotFoundError
	}

	return premise, nil
}
