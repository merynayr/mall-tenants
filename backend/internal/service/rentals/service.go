package rentals

import (
	"github.com/merynayr/mall-tenants/internal/client/db"
	"github.com/merynayr/mall-tenants/internal/repository"
	"github.com/merynayr/mall-tenants/internal/service"
)

type srv struct {
	rentalRepository  repository.RentalRepository
	premiseRepository repository.PremiseRepository
	txManager         db.TxManager
}

// NewService возвращает новый объект сервисного слоя mall-tenants
func NewService(
	rentalRepo repository.RentalRepository,
	premiseRepo repository.PremiseRepository,
	txManager db.TxManager,
) service.RentalService {
	return &srv{
		rentalRepository:  rentalRepo,
		premiseRepository: premiseRepo,
		txManager:         txManager,
	}
}
