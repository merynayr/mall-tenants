package mall

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/merynayr/mall-tenants/internal/sys"
)

// GetPremisesByCode получает данные о помещении
// @Summary Получить данные о помещении
// @Description Получает данные о помещении по коду
// @Tags premises
// @Accept  json
// @Produce  json
// @Param code path int64 true "Код помещения"
// @Success 200 {object} model.Premises
// @Failure 400 {object} sys.ErrorResponse
// @Failure 401 {object} sys.ErrorResponse
// @Failure 404 {object} sys.ErrorResponse "Помещение не найдено"
// @Failure 500 {object} sys.ErrorResponse
// @Router /premise/{code} [get]
func (a *API) GetPremisesByCode(c *gin.Context) {
	codeParam := c.Param("code")

	code, err := strconv.ParseInt(codeParam, 10, 64)
	if err != nil {
		sys.HandleError(c, sys.InvalidRequestError)
		return
	}
	premise, err := a.premiseService.GetPremisesByCode(c.Request.Context(), code)
	if err != nil {
		sys.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, premise)
}
