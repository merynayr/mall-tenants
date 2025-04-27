package payment

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/service"
	"github.com/merynayr/mall-tenants/internal/sys"
)

// API структура для работы с Payment и Rental сервисами
type API struct {
	paymentService service.PaymentService
}

// NewAPI возвращает новый объект API с зависимостями
func NewAPI(paymentService service.PaymentService) *API {
	return &API{
		paymentService: paymentService,
	}
}

// RegisterRoutes регистрирует все маршруты API для работы с платежами
func (api *API) RegisterRoutes(router *gin.Engine) {
	paymentGroup := router.Group("/payment")
	{
		paymentGroup.POST("/", api.CreatePayment)
		paymentGroup.GET("/:paymentID", api.GetPaymentByID)
		paymentGroup.GET("/rental/:rentalID/last", api.GetLastPaymentByRentalID)
		paymentGroup.GET("/rental/:rentalID/overdue", api.GetOverduePayments)
	}
}

// CreatePayment создаёт новый платеж
// @Summary Создать платеж
// @Description Создает новый платеж для аренды
// @Tags payment
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param payment body model.Payment true "Данные платежа"
// @Success 201 "Payment created successfully"
// @Failure 400 {object} sys.ErrorResponse
// @Failure 401 {object} sys.ErrorResponse
// @Failure 500 {object} sys.ErrorResponse
// @Router /payment/ [post]
func (api *API) CreatePayment(c *gin.Context) {
	var payment model.Payment

	if err := c.ShouldBindJSON(&payment); err != nil {
		sys.HandleError(c, sys.InvalidRequestError)
		return
	}

	err := api.paymentService.CreatePayment(c, payment)
	if err != nil {
		sys.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Payment created successfully"})
}

// GetPaymentByID получает платеж по ID
// @Summary Получить платеж
// @Description Получает информацию о платеже по его ID
// @Tags payment
// @Produce  json
// @Security BearerAuth
// @Param paymentID path int true "ID платежа"
// @Success 200 {object} model.Payment
// @Failure 400 {object} sys.ErrorResponse
// @Failure 401 {object} sys.ErrorResponse
// @Failure 404 {object} sys.ErrorResponse
// @Failure 500 {object} sys.ErrorResponse
// @Router /payment/{paymentID} [get]
func (api *API) GetPaymentByID(c *gin.Context) {
	idParam := c.Param("paymentID")

	paymentID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		sys.HandleError(c, sys.InvalidRequestError)
		return
	}

	payment, err := api.paymentService.GetPaymentByID(c, paymentID)
	if err != nil {
		sys.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, payment)
}

// GetLastPaymentByRentalID получает последний платеж по аренде
// @Summary Получить последний платеж
// @Description Получает последний платеж по rentalID
// @Tags payment
// @Produce  json
// @Security BearerAuth
// @Param rentalID path int true "ID аренды"
// @Success 200 {object} model.Payment
// @Failure 400 {object} sys.ErrorResponse
// @Failure 401 {object} sys.ErrorResponse
// @Failure 404 {object} sys.ErrorResponse
// @Failure 500 {object} sys.ErrorResponse
// @Router /payment/rental/{rentalID}/last [get]
func (api *API) GetLastPaymentByRentalID(c *gin.Context) {
	idParam := c.Param("rentalID")

	rentalID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		sys.HandleError(c, sys.InvalidRequestError)
		return
	}

	lastPayment, err := api.paymentService.GetLastPaymentByRentalID(c, rentalID)
	if err != nil {
		sys.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, lastPayment)
}

// GetOverduePayments получает просроченные платежи
// @Summary Получить просроченные платежи
// @Description Получает список просроченных платежей по rentalID
// @Tags payment
// @Produce  json
// @Security BearerAuth
// @Param rentalID path int true "ID аренды"
// @Success 200 {array} model.Payment
// @Failure 400 {object} sys.ErrorResponse
// @Failure 401 {object} sys.ErrorResponse
// @Failure 404 {object} sys.ErrorResponse
// @Failure 500 {object} sys.ErrorResponse
// @Router /payment/rental/{rentalID}/overdue [get]
func (api *API) GetOverduePayments(c *gin.Context) {
	idParam := c.Param("rentalID")

	rentalID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		sys.HandleError(c, sys.InvalidRequestError)
		return
	}

	overduePayments, err := api.paymentService.GetOverduePayments(c, rentalID)
	if err != nil {
		sys.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, overduePayments)
}
