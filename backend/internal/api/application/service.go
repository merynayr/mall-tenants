package application

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/merynayr/mall-tenants/internal/logger"
	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/service"
	"github.com/merynayr/mall-tenants/internal/sys"
)

// API auth структура
type API struct {
	appService service.ApplicationService
}

// NewAPI возвращает новый объект имплементации API-слоя application
func NewAPI(appService service.ApplicationService) *API {
	return &API{appService: appService}
}

// RegisterRoutes регистрирует все маршруты API для работы с платежами
func (api *API) RegisterRoutes(router *gin.Engine) {
	applicationGroup := router.Group("/applications")
	{
		applicationGroup.POST("/", api.Create)
		applicationGroup.GET("/", api.GetAll)
		applicationGroup.PATCH("/:id/status", api.UpdateStatus)
	}
}

// Create создаёт новую заявку
// @Summary      Создать заявку
// @Description  Создаёт новую заявку на аренду помещения
// @Tags         applications
// @Accept       json
// @Produce      json
// @Param        application  body      model.Application  true  "Данные заявки"
// @Success      201  {object}  map[string]string  "message: Application submitted successfully"
// @Failure      400  {object}  sys.ErrorResponse
// @Failure      500  {object}  sys.ErrorResponse
// @Router       /applications/ [post]
func (api *API) Create(c *gin.Context) {
	var input model.Application
	err := c.ShouldBindJSON(&input)
	if err != nil {
		logger.Error(err.Error())
		sys.HandleError(c, sys.InvalidRequestError)
		return
	}

	if err := api.appService.CreateApplication(c.Request.Context(), &input); err != nil {
		sys.HandleError(c, errors.New("Failed to save application"))
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Application submitted successfully"})
}

// GetAll получает все заявки
// @Summary      Получить все заявки
// @Description  Возвращает список всех заявок. Можно указать ?processed=true/false для фильтрации
// @Tags         applications
// @Produce      json
// @Security     BearerAuth
// @Param        processed  query  bool  false  "Фильтрация по статусу обработки"
// @Success      200  {array}   model.Application
// @Failure      500  {object}  sys.ErrorResponse
// @Router       /applications/ [get]
func (api *API) GetAll(c *gin.Context) {
	var processed *bool
	if value := c.Query("processed"); value != "" {
		boolVal, err := strconv.ParseBool(value)
		if err != nil {
			sys.HandleError(c, sys.InvalidRequestError)
			return
		}
		processed = &boolVal
	}

	apps, err := api.appService.GetApplications(c.Request.Context(), processed)
	if err != nil {
		sys.HandleError(c, errors.New("Failed to retrieve applications"))
		return
	}
	c.JSON(http.StatusOK, apps)
}

// UpdateStatus обновляет статус заявки
// @Summary      Обновить статус заявки
// @Description  Обновляет флаг processed заявки по ID
// @Tags         applications
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path     int               true  "ID заявки"
// @Param        body  body     updateStatusInput true  "Статус заявки"
// @Success      200   {object} gin.H
// @Failure      400   {object} sys.ErrorResponse
// @Failure      500   {object} sys.ErrorResponse
// @Router       /applications/{id}/status [patch]
func (api *API) UpdateStatus(c *gin.Context) {
	var input updateStatusInput
	if err := c.ShouldBindJSON(&input); err != nil {
		sys.HandleError(c, sys.InvalidRequestError)
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		logger.Error(err.Error())
		sys.HandleError(c, sys.InvalidRequestError)
		return
	}

	if err := api.appService.UpdateApplicationStatus(c.Request.Context(), id, input.Processed); err != nil {
		sys.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Статус обновлён"})
}

type updateStatusInput struct {
	Processed bool `json:"processed" binding:"required"`
}
