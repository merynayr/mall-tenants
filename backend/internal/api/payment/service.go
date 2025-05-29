package payment

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/service"
	"github.com/merynayr/mall-tenants/internal/sys"
	"github.com/merynayr/mall-tenants/internal/sys/codes"
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
	paymentGroup := router.Group("/payments")
	{
		paymentGroup.POST("", api.CreatePayment)
		paymentGroup.PATCH(":id/pay", api.MarkAsPaid)
		paymentGroup.PATCH("/mark-paid", api.MarkPaymentsPaid)
		paymentGroup.GET(":id", api.GetPayment)
		paymentGroup.GET("", api.ListPayments)
		paymentGroup.GET("/client", api.ListClientPayments)
	}
}

// CreatePayment godoc
// @Summary Создать платеж
// @Tags payments
// @Accept json
// @Produce json
// @Param payment body model.Payment true "Данные платежа"
// @Success 201 {object} map[string]int64
// @Failure 400,500 {object} sys.ErrorResponse
// @Router /payments [post]
func (api *API) CreatePayment(c *gin.Context) {
	var p model.Payment
	if err := c.ShouldBindJSON(&p); err != nil {
		sys.HandleError(c, err)
		return
	}

	err := api.paymentService.CreatePayment(c.Request.Context(), p)
	if err != nil {
		sys.HandleError(c, err)
		return
	}

	c.Status(http.StatusCreated)
}

// MarkAsPaid godoc
// @Summary Отметить платеж как оплаченный
// @Tags payments
// @Param id path int true "ID платежа"
// @Success 204 "Платеж отмечен как оплаченный"
// @Failure 400,404,500 {object} sys.ErrorResponse
// @Router /payments/{id}/pay [patch]
func (api *API) MarkAsPaid(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		sys.HandleError(c, err)
		return
	}

	if err := api.paymentService.MarkAsPaid(c.Request.Context(), id); err != nil {
		sys.HandleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

// MarkPaymentsPaid помечает указанные платежи как оплаченные
// @Summary Отметить платежи как оплаченные
// @Description Обновляет статус is_paid и увеличивает paid_months у аренды
// @Tags payments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body model.MarkPaymentsPaidRequest true "ID платежей"
// @Success 204 "Успешно"
// @Failure 400 {object} sys.ErrorResponse "Неверный запрос"
// @Failure 500 {object} sys.ErrorResponse "Внутренняя ошибка сервера"
// @Router /payments/mark-paid [patch]
func (api *API) MarkPaymentsPaid(c *gin.Context) {
	var req model.MarkPaymentsPaidRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		sys.HandleError(c, sys.Wrap(err, sys.InvalidRequestError))
		return
	}

	if len(req.PaymentIDs) == 0 {
		sys.HandleError(c, sys.NoPaymentIDsError)
		return
	}

	err := api.paymentService.MarkPaymentsAsPaid(c.Request.Context(), req.PaymentIDs)
	if err != nil {
		sys.HandleError(c, sys.MarkPaymentsFailedError)
		return
	}

	c.Status(http.StatusNoContent)
}

// GetPayment godoc
// @Summary Получить платеж по ID
// @Tags payments
// @Param id path int true "ID платежа"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} model.Payment
// @Failure 400,404,500 {object} sys.ErrorResponse
// @Router /payments/{id} [get]
func (api *API) GetPayment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		sys.HandleError(c, sys.Wrap(err, sys.InvalidRequestError))
		return
	}

	p, err := api.paymentService.GetByID(c.Request.Context(), id)
	if err != nil {
		sys.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, p)
}

// ListClientPayments godoc
// @Summary      Получить список платёжных операций клиента по ID
// @Description  Возвращает список платежей конкретного клиента с фильтрацией, поиском и сортировкой
// @Tags         payments
// @Accept       json
// @Produce      json
// @Param        rental_id       query     string false  "Поиск по номеру договора (нечёткий поиск)"
// @Param        is_paid      query     bool   false  "Фильтрация по статусу оплаты (true/false)"
// @Param        sort_by      query     string false  "Поле сортировки (например, period_start)"
// @Param        sort_order   query     string false  "Порядок сортировки: asc или desc"
// @Param        limit        query     int    false  "Количество записей на страницу (по умолчанию 20)"
// @Param        offset       query     int    false  "Смещение (offset) от начала выборки"
// @Success      200          {array}  model.Payment
// @Failure      400          {object}  sys.ErrorResponse "Некорректный запрос"
// @Failure      500          {object}  sys.ErrorResponse "Внутренняя ошибка сервера"
// @Router       /payments/client/ [get]
func (api *API) ListClientPayments(c *gin.Context) {
	var filter model.PaymentFilter

	clientID := c.GetInt64("user_id")

	if isPaidStr := c.Query("is_paid"); isPaidStr != "" {
		isPaid, err := strconv.ParseBool(isPaidStr)
		if err != nil {
			sys.HandleError(c, sys.Wrap(err, sys.InvalidRequestError))
			return
		}
		filter.IsPaid = &isPaid
	}

	if searchStr := c.Query("rental_id"); searchStr != "" {
		search, err := strconv.ParseInt(searchStr, 10, 64)
		if err != nil {
			sys.HandleError(c, sys.Wrap(err, sys.NewCommonError("Неверный формат поиска", codes.BadRequest)))
			return
		}
		filter.Search = search
	}
	filter.SortBy = c.DefaultQuery("sort_by", "period_start")
	filter.SortOrder = c.DefaultQuery("sort_order", "desc")

	limitStr := c.DefaultQuery("limit", "20")
	limit, err := strconv.ParseUint(limitStr, 10, 64)
	if err != nil {
		sys.HandleError(c, sys.Wrap(err, sys.InvalidRequestError))
		return
	}
	filter.Limit = limit

	offsetStr := c.DefaultQuery("offset", "0")
	offset, err := strconv.ParseUint(offsetStr, 10, 64)
	if err != nil {
		sys.HandleError(c, sys.Wrap(err, sys.InvalidRequestError))
		return
	}
	filter.Offset = offset

	payments, err := api.paymentService.GetByClientID(c.Request.Context(), clientID, filter)
	if err != nil {
		sys.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, payments)
}

// ListPayments godoc
// @Summary      Получить список платёжных операций
// @Description  Возвращает список платежей с фильтрацией, поиском и сортировкой
// @Tags         payments
// @Accept       json
// @Produce      json
// @Param        is_paid     query     bool   false  "Фильтрация по статусу оплаты (true/false)"
// @Param        client      query     string false  "Поиск по имени клиента (нечёткий поиск)"
// @Param        sort_by     query     string false  "Поле сортировки (например, period_start)"
// @Param        sort_order  query     string false  "Порядок сортировки: asc или desc"
// @Param        limit       query     int    false  "Количество записей на страницу (по умолчанию 20)"
// @Param        offset      query     int    false  "Смещение (offset) от начала выборки"
// @Success      200         {array}   model.Payment
// @Failure      400         {object}  sys.ErrorResponse "Некорректный запрос"
// @Failure      500         {object}  sys.ErrorResponse "Внутренняя ошибка сервера"
// @Router       /payments [get]
func (api *API) ListPayments(c *gin.Context) {
	var filter model.PaymentFilter

	if isPaidStr := c.Query("is_paid"); isPaidStr != "" {
		isPaid, err := strconv.ParseBool(isPaidStr)
		if err != nil {
			sys.HandleError(c, sys.Wrap(err, sys.InvalidRequestError))
			return
		}
		filter.IsPaid = &isPaid
	}

	filter.Search = c.Query("client")
	filter.SortBy = c.DefaultQuery("sort_by", "period_start")
	filter.SortOrder = c.DefaultQuery("sort_order", "desc")

	limitStr := c.DefaultQuery("limit", "20")
	limit, err := strconv.ParseUint(limitStr, 10, 64)
	if err != nil {
		sys.HandleError(c, sys.Wrap(err, sys.InvalidRequestError))
		return
	}
	filter.Limit = limit

	offsetStr := c.DefaultQuery("offset", "0")
	offset, err := strconv.ParseUint(offsetStr, 10, 64)
	if err != nil {
		sys.HandleError(c, sys.Wrap(err, sys.InvalidRequestError))
		return
	}
	filter.Offset = offset

	payments, err := api.paymentService.ListPayments(c.Request.Context(), filter)
	if err != nil {
		sys.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, payments)
}
