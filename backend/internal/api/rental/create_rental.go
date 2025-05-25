package rental

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/sys"
)

// CreateRental создает новую аренду
// @Summary Создать аренду
// @Description Создает новый договор аренды
// @Tags rental
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param rental body model.Rental true "Данные аренды"
// @Success 201 {object} model.Rental
// @Failure 400 {object} sys.ErrorResponse
// @Failure 401 {object} sys.ErrorResponse
// @Failure 500 {object} sys.ErrorResponse
// @Router /rental/ [post]
func (a *API) CreateRental(c *gin.Context) {
	var rental model.Rental
	if err := c.ShouldBindJSON(&rental); err != nil {
		sys.HandleError(c, sys.Wrap(err, sys.InvalidRequestError))
		return
	}

	err := a.rentalService.CreateRental(c.Request.Context(), rental)
	if err != nil {
		sys.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, rental)
}
