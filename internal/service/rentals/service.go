package rentals

import (
	"github.com/merynayr/mall-tenants/internal/repository"
	"github.com/merynayr/mall-tenants/internal/service"
)

type srv struct {
	rentalRepository repository.RentalRepository
}

// NewService возвращает новый объект сервисного слоя mall-tenants
func NewService(rentalRepo repository.RentalRepository) service.RentalService {
	return &srv{
		rentalRepository: rentalRepo,
	}
}
