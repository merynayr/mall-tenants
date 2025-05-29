package applications

import (
	"context"
	"time"

	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/repository"
	"github.com/merynayr/mall-tenants/internal/service"
	"github.com/merynayr/mall-tenants/internal/sys"
)

type srv struct {
	applicationRepository repository.ApplicationRepository
}

// NewService возвращает новый объект сервисного слоя mall-tenants
func NewService(applicationRepo repository.ApplicationRepository) service.ApplicationService {
	return &srv{
		applicationRepository: applicationRepo,
	}
}

func (s *srv) CreateApplication(ctx context.Context, a *model.Application) error {
	a.CreatedAt = time.Now().UTC()
	a.IsProcessed = false

	err := s.applicationRepository.Create(ctx, a)
	return err
}

func (s *srv) GetApplications(ctx context.Context, filter model.ApplicationFilter) ([]*model.Application, error) {
	apps, err := s.applicationRepository.GetAll(ctx, filter)
	if err != nil {
		return nil, err
	}
	if len(apps) == 0 {
		return nil, sys.ApplicationsNotFoundError
	}
	return apps, nil
}

func (s *srv) UpdateApplicationStatus(ctx context.Context, id int64, isProcessed bool) error {
	err := s.applicationRepository.SetProcessedStatus(ctx, id, isProcessed)
	return err
}
