package contract

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/service"
	"github.com/merynayr/mall-tenants/internal/sys"
)

// API структура для работы с договорами
type API struct {
	contractService service.ContractService
}

// NewAPI возвращает новый объект API с зависимостями
func NewAPI(contractService service.ContractService) *API {
	return &API{
		contractService: contractService,
	}
}

// RegisterRoutes регистрирует все маршруты API для работы с платежами
func (API *API) RegisterRoutes(router *gin.Engine) {
	contractGroup := router.Group("/contracts")
	{
		contractGroup.POST("/:contract_id/sign", API.SignContract)
		contractGroup.GET("/:contract_id", API.DownloadContract)
		contractGroup.DELETE("/:contract_id", API.DeleteContract)
	}
}

// DownloadContract отдает файл договора по contract_id
// @Summary Скачать договор аренды
// @Description Возвращает PDF-файл договора, связанного с арендуемой записью
// @Tags contracts
// @Produce application/pdf
// @Security BearerAuth
// @Param contract_id path int64 true "ID аренды"
// @Success 200 {file} file
// @Failure 400 {object} sys.ErrorResponse
// @Failure 404 {object} sys.ErrorResponse "Договор не найден"
// @Failure 500 {object} sys.ErrorResponse
// @Router /contracts/{contract_id} [get]
func (API *API) DownloadContract(c *gin.Context) {
	contractIDParam := c.Param("contract_id")
	contractID, err := strconv.ParseInt(contractIDParam, 10, 64)
	if err != nil {
		sys.HandleError(c, sys.Wrap(err, sys.InvalidIDFormatError))
		return
	}

	contract, err := API.contractService.GetByContractID(c.Request.Context(), contractID)
	if err != nil {
		sys.HandleError(c, err)
		return
	}

	file, err := os.Open(contract.FilePath)
	if err != nil {
		sys.HandleError(c, err)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			sys.HandleError(c, err)
			return
		}
	}()

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=contract-%d.pdf", contract.ContractID))
	c.Header("Content-Type", "application/pdf")
	c.File(contract.FilePath)
}

// SignContract подписывает контракт и сохраняет аренду
// @Summary Подписать контракт
// @Description Подписывает контракт по его идентификатору и сохраняет связанную аренду
// @Tags contracts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param contract_id path int64 true "ID контракта"
// @Param rental body model.RentalWithContract true "Данные аренды, связанные с контрактом"
// @Success 200 {object} map[string]string "Сообщение об успешной подписи контракта"
// @Failure 400 {object} sys.ErrorResponse "Некорректный запрос"
// @Failure 401 {object} sys.ErrorResponse "Неавторизован"
// @Failure 404 {object} sys.ErrorResponse "Контракт не найден"
// @Failure 500 {object} sys.ErrorResponse "Внутренняя ошибка сервера"
// @Router /contracts/{contract_id}/sign [post]
func (API *API) SignContract(c *gin.Context) {
	contractIDParam := c.Param("contract_id")
	contractID, err := strconv.ParseInt(contractIDParam, 10, 64)
	if err != nil {
		sys.HandleError(c, sys.Wrap(err, sys.InvalidIDFormatError))
		return
	}

	var rent model.RentalWithContract
	if err := c.ShouldBindJSON(&rent); err != nil {
		sys.HandleError(c, sys.Wrap(err, sys.InvalidRequestError))
		return
	}

	err = API.contractService.SignContract(c.Request.Context(), contractID, rent)
	if err != nil {
		sys.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "contract signed"})
}

// DeleteContract удаляет существующий контракт.
// @Summary Удалить контракт
// @Description Удаляет контракт по идентификатору, если он существует.
// @Tags contracts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID контракта"
// @Success 200 {string} string "ok"
// @Failure 400 {object} sys.ErrorResponse "Некорректный идентификатор"
// @Failure 401 {object} sys.ErrorResponse "Неавторизован"
// @Failure 404 {object} sys.ErrorResponse "Контракт не найден"
// @Failure 500 {object} sys.ErrorResponse "Внутренняя ошибка сервера"
// @Router /contracts/{contract_id} [delete]
func (API *API) DeleteContract(c *gin.Context) {
	codeParam := c.Param("contract_id")

	id, err := strconv.ParseInt(codeParam, 10, 64)
	if err != nil {
		sys.HandleError(c, sys.Wrap(err, sys.InvalidRequestError))
		return
	}

	err = API.contractService.DeleteContract(c.Request.Context(), id)
	if err != nil {
		sys.HandleError(c, err)
		return
	}

	c.Status(http.StatusOK)
}
