package rental

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/merynayr/mall-tenants/internal/service"
	"github.com/merynayr/mall-tenants/internal/sys"
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
		rentalGroup.GET("/:id", api.GetRentalsByID)
		rentalGroup.PATCH("/:id", api.UpdateRental)
		rentalGroup.POST("/", api.CreateRental)
		rentalGroup.GET("/", api.GetAgreements)
	}
}

// GetAgreements godoc
// @Summary      Получить список договоров аренды
// @Description  Возвращает список договоров аренды с пагинацией
// @Tags         Договора
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Param        limit  query     int  false  "Максимальное количество договоров" default(20)
// @Param        offset query     int  false  "Смещение для пагинации" default(0)
// @Success      200    {array}   model.Rental
// @Failure      400    {object}  sys.ErrorResponse
// @Failure      500    {object}  sys.ErrorResponse
// @Router       /rental/ [get]
func (api *API) GetAgreements(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "20")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.ParseUint(limitStr, 10, 64)
	if err != nil {
		sys.HandleError(c, sys.Wrap(err, sys.InvalidRequestError))
		return
	}

	offset, err := strconv.ParseUint(offsetStr, 10, 64)
	if err != nil {
		sys.HandleError(c, sys.Wrap(err, sys.InvalidRequestError))
		return
	}

	agreements, err := api.rentalService.GetAgreements(c.Request.Context(), limit, offset)
	if err != nil {
		sys.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, agreements)
}
