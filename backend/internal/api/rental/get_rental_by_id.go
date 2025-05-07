package rental

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/merynayr/mall-tenants/internal/sys"
)

// GetRentalByID получает данные о аренде
// @Summary Получить данные о аренде
// @Description Получает данные о аренде по коду
// @Tags rental
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path int64 true "Код аренды"
// @Success 200 {object} model.Rental
// @Failure 400 {object} sys.ErrorResponse
// @Failure 401 {object} sys.ErrorResponse
// @Failure 404 {object} sys.ErrorResponse "Аренда не найдена"
// @Failure 500 {object} sys.ErrorResponse
// @Router /rental/{id} [get]
func (a *API) GetRentalByID(c *gin.Context) {
	codeParam := c.Param("id")

	id, err := strconv.ParseInt(codeParam, 10, 64)
	if err != nil {
		sys.HandleError(c, sys.InvalidRequestError)
		return
	}

	rental, err := a.rentalService.GetRentalByID(c.Request.Context(), id)
	if err != nil {
		sys.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, rental)
}
