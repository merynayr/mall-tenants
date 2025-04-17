package premises

import (
	"github.com/merynayr/mall-tenants/internal/repository"
	"github.com/merynayr/mall-tenants/internal/service"
)

type srv struct {
	premiseRepository repository.PremiseRepository
}

// NewService возвращает новый объект сервисного слоя mall-tenants
func NewService(premiseRepo repository.PremiseRepository) service.PremiseService {
	return &srv{
		premiseRepository: premiseRepo,
	}
}
