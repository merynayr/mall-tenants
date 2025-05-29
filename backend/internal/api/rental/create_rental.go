package rental

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/sys"
)

// CreateRental создает новую аренду с возможной генерацией договора
// @Summary Создать аренду
// @Description Создает новую аренду и генерирует черновик договора. Можно загрузить собственный шаблон .docx или использовать шаблон по умолчанию.
// @Tags rental
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param space_code formData int true "ID помещения"
// @Param client_id formData int true "ID клиента"
// @Param start_date formData string true "Дата начала аренды (в формате RFC3339)"
// @Param end_date formData string true "Дата окончания аренды (в формате RFC3339)"
// @Param template_file formData file false "Файл шаблона договора .docx (необязательно)"
// @Success 201 {object} map[string]interface{} "Аренда и ссылка на предварительный просмотр договора"
// @Failure 400 {object} sys.ErrorResponse
// @Failure 401 {object} sys.ErrorResponse
// @Failure 500 {object} sys.ErrorResponse
// @Router /rental/ [post]
func (api *API) CreateRental(c *gin.Context) {
	data := c.PostForm("data")
	if data == "" {
		sys.HandleError(c, sys.InvalidRequestError)
		return
	}

	var rental model.Rental
	if err := json.Unmarshal([]byte(data), &rental); err != nil {
		sys.HandleError(c, sys.Wrap(err, sys.InvalidRequestError))
		return
	}

	file, err := c.FormFile("template_file")
	if err != nil {
		sys.HandleError(c, sys.Wrap(err, sys.InvalidRequestError))
		return
	}
	rental.TemplateFile = file

	err = api.rentalService.CreateRental(c.Request.Context(), rental)
	if err != nil {
		sys.HandleError(c, err)
		return
	}

	c.Status(http.StatusCreated)
}
