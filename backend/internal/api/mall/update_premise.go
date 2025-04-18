package mall

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/sys"
)

// UpdatePremise обновляет данные о помещении
// @Summary Обновить помещение
// @Description Обновляет данные о помещении
// @Tags premises
// @Accept  json
// @Produce  json
// @Param body body model.Premises true "Информация о помещении"
// @Success 200
// @Failure 400 {object} sys.ErrorResponse
// @Failure 401 {object} sys.ErrorResponse
// @Failure 404 {object} sys.ErrorResponse "Помещение не найдено"
// @Failure 500 {object} sys.ErrorResponse
// @Router /premise/ [patch]
func (a *API) UpdatePremise(c *gin.Context) {
	var req model.Premises

	if err := c.ShouldBindJSON(&req); err != nil {
		sys.HandleError(c, sys.InvalidRequestError)
		return
	}

	err := a.premiseService.UpdatePremise(c.Request.Context(), req)
	if err != nil {
		sys.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
