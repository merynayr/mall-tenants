package rental

import (
	"github.com/gin-gonic/gin"
	"github.com/merynayr/mall-tenants/internal/service"
)

// API rental структура
type API struct {
	rentalService service.RentalService
}

// NewAPI возвращает новый объект имплементации API-слоя mall
func NewAPI(rentalService service.RentalService) *API {
	return &API{
		rentalService: rentalService,
	}
}

// RegisterRoutes регистрирует маршруты
func (api *API) RegisterRoutes(router *gin.Engine) {
	rentalGroup := router.Group("/rental")
	{
		rentalGroup.GET("/:id", api.GetRentalByID)
		rentalGroup.PATCH("/:id", api.UpdateRental)
		rentalGroup.POST("/", api.CreateRental)
	}
}
