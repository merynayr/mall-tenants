package premises

import (
	"context"

	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/sys"
)

func (s *srv) CreatePremise(ctx context.Context, premise model.Premises) error {
	_, exist, err := s.premiseRepository.GetPremisesByCode(ctx, premise.Code)
	if err != nil {
		return err
	}
	if exist {
		return sys.PremiseAlreadyExistsError
	}

	err = s.premiseRepository.CreatePremise(ctx, &premise)
	if err != nil {
		return err
	}
	return nil
}
