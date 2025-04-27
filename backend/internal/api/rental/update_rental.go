package rental

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/sys"
)

// UpdateRental обновляет существующую аренду
// @Summary Обновить аренду
// @Description Обновляет данные аренды по коду
// @Tags rental
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path int64 true "Код аренды"
// @Param rental body model.Rental true "Обновленные данные аренды"
// @Success 200 {object} model.Rental
// @Failure 400 {object} sys.ErrorResponse
// @Failure 401 {object} sys.ErrorResponse
// @Failure 404 {object} sys.ErrorResponse "Аренда не найдена"
// @Failure 500 {object} sys.ErrorResponse
// @Router /rental/{id}  [patch]
func (a *API) UpdateRental(c *gin.Context) {
	codeParam := c.Param("id")

	id, err := strconv.ParseInt(codeParam, 10, 64)
	if err != nil {
		sys.HandleError(c, sys.InvalidRequestError)
		return
	}

	var rental model.Rental
	if err := c.ShouldBindJSON(&rental); err != nil {
		sys.HandleError(c, sys.InvalidRequestError)
		return
	}

	rental.RentalID = id

	err = a.rentalService.UpdateRental(c.Request.Context(), rental)
	if err != nil {
		sys.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, rental)
}
