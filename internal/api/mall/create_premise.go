package mall

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/sys"
)

// CreatePremise создание помещения
// @Summary Создать помещение
// @Description Создаёт новое помещение
// @Tags create
// @Accept  json
// @Produce  json
// @Param body body model.Premises true "Информация о помещении"
// @Success 200
// @Failure 400 {object} sys.ErrorResponse
// @Failure 401 {object} sys.ErrorResponse
// @Failure 500 {object} sys.ErrorResponse
// @Router /mall/ [post]
func (a *API) CreatePremise(c *gin.Context) {
	var req model.Premises

	if err := c.ShouldBindJSON(&req); err != nil {
		sys.HandleError(c, sys.InvalidRequestError)
		return
	}

	err := a.premiseService.CreatePremise(c.Request.Context(), req)
	if err != nil {
		sys.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
