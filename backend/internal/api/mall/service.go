package mall

import (
	"github.com/gin-gonic/gin"
	"github.com/merynayr/mall-tenants/internal/service"
)

// API premise структура
type API struct {
	premiseService service.PremiseService
}

// NewAPI возвращает новый объект имплементации API-слоя mall
func NewAPI(premiseService service.PremiseService) *API {
	return &API{
		premiseService: premiseService,
	}
}

// RegisterRoutes регистрирует маршруты
func (api *API) RegisterRoutes(router *gin.Engine) {
	mallGroup := router.Group("/mall")
	{
		mallGroup.GET("/:code", api.GetPremisesByCode)
		mallGroup.PATCH("/", api.UpdatePremise)
		mallGroup.POST("/", api.CreatePremise)
	}
}
