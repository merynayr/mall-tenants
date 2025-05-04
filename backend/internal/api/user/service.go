package user

import (
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
	userService service.UserService
}

// NewAPI возвращает новый объект имплементации API-слоя auth
func NewAPI(userService service.UserService) *API {
	return &API{
		userService: userService,
	}
}

// RegisterRoutes регистрирует маршруты
func (api *API) RegisterRoutes(router *gin.Engine) {
	authGroup := router.Group("/clients")
	{
		authGroup.GET("/", api.GetClients)
		authGroup.POST("/", api.CreateClient)
		authGroup.GET("/profile", api.GetProfile)
	}
}

// GetClients godoc
// @Summary      Получить список клиентов
// @Description  Возвращает список клиентов с пагинацией
// @Tags         Клиенты
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Param        limit  query     int  false  "Максимальное количество клиентов" default(20)
// @Param        offset query     int  false  "Смещение для пагинации" default(0)
// @Success      200    {array}   model.Client
// @Failure      400    {object}  sys.ErrorResponse
// @Failure      500    {object}  sys.ErrorResponse
// @Router       /clients/ [get]
func (api *API) GetClients(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "20")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.ParseUint(limitStr, 10, 64)
	if err != nil {
		sys.HandleError(c, sys.InvalidRequestError)
		return
	}

	offset, err := strconv.ParseUint(offsetStr, 10, 64)
	if err != nil {
		sys.HandleError(c, sys.InvalidRequestError)
		return
	}

	clients, err := api.userService.GetClients(c.Request.Context(), limit, offset)
	if err != nil {
		sys.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, clients)
}

// CreateClient создание клиента
// @Summary Создать клиента
// @Description Создаёт нового клиента
// @Tags Клиенты
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param body body model.RegisterRequest true "Данные для регистрации"
// @Success 200 {object} model.AuthResponse
// @Failure 400 {object} sys.ErrorResponse
// @Failure 401 {object} sys.ErrorResponse
// @Failure 500 {object} sys.ErrorResponse
// @Router /clients/ [post]
func (api *API) CreateClient(c *gin.Context) {
	var req model.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Debug(err.Error())
		sys.HandleError(c, sys.InvalidRequestError)
		return
	}

	err := api.userService.CreateClient(c.Request.Context(), req)
	if err != nil {
		sys.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

// GetProfile godoc
// @Summary      Получить информацию о клиенте
// @Description  Возвращает информацию о клиенте по email (query-параметр)
// @Tags         Клиенты
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200    {object}  model.Client
// @Failure      400    {object}  sys.ErrorResponse
// @Failure      500    {object}  sys.ErrorResponse
// @Router       /clients/profile [get]
func (api *API) GetProfile(c *gin.Context) {
	email := c.GetString("email")
	client, err := api.userService.GetUserByEmail(c.Request.Context(), email)
	if err != nil {
		sys.HandleError(c, err)
		return
	}
	client.Password = ""

	c.JSON(http.StatusOK, client)
}
