package premise

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merynayr/mall-tenants/internal/sys"
)

// GetAllPremises получает данные  о помещениях
// @Summary Получить список помещений
// @Description Получает список помещений
// @Tags premises
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Success 200 {object} []model.Premises
// @Failure 400 {object} sys.ErrorResponse
// @Failure 401 {object} sys.ErrorResponse
// @Failure 404 {object} sys.ErrorResponse "Помещений не найдено"
// @Failure 500 {object} sys.ErrorResponse
// @Router /premise [get]
func (a *API) GetAllPremises(c *gin.Context) {
	premises, err := a.premiseService.GetAllPremises(c.Request.Context())
	if err != nil {
		sys.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, premises)
}
