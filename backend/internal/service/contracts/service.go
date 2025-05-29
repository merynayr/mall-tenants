package contracts

import (
	"context"

	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/repository"
	"github.com/merynayr/mall-tenants/internal/service"
)

type srv struct {
	contractRepository repository.ContractRepository
}

func NewService(r repository.ContractRepository) service.ContractService {
	return &srv{contractRepository: r}
}

func (s *srv) CreateContract(ctx context.Context, c model.Contract) error {
	return s.contractRepository.CreateContract(ctx, &c)
}

func (s *srv) GetAllContracts(ctx context.Context) ([]model.Contract, error) {
	return s.contractRepository.GetAll(ctx)
}

func (s *srv) GetByContractID(ctx context.Context, contractID int64) (*model.Contracts, error) {
	contract, err := s.contractRepository.GetByContractID(ctx, contractID)
	return contract, err
}
