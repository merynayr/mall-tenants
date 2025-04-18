package mall

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/service"
	"github.com/merynayr/mall-tenants/internal/sys"
)

// API premise структура
type API struct {
	premiseService service.PremiseService
}

// NewAPI возвращает новый объект имплементации API-слоя mall
func NewAPI(premiseService service.PremiseService) *API {
	return &API{
		premiseService: premiseService,
	}
}

// RegisterRoutes регистрирует маршруты
func (api *API) RegisterRoutes(router *gin.Engine) {
	mallGroup := router.Group("/premise")
	{
		mallGroup.GET("/:code", api.GetPremisesByCode)
		mallGroup.GET("/", api.GetAllPremises)
		mallGroup.PATCH("/", api.UpdatePremise)
		mallGroup.POST("/", api.CreatePremise)
		mallGroup.POST("/:code/polygons", api.AddPolygon)
		mallGroup.GET("/polygons", api.GetPolygons)
	}
	floorGroup := router.Group("/floor-plan/:floor")
	{
		floorGroup.POST("", api.UploadFloorPlan)
		floorGroup.GET("", api.GetFloorPlan)
	}
}

// UploadFloorPlan загружает SVG‑план этажа
// @Summary      Загрузить SVG‑план этажа
// @Description  Принимает multipart/form-data с полем file и сохраняет SVG на диск + в БД
// @Tags         floor-plans
// @Accept       multipart/form-data
// @Produce      json
// @Param        floor path     int  true  "Номер этажа"
// @Param        file  formData file true  "SVG‑файл плана этажа"
// @Success      201 {object}   map[string]string  "{"message":"floor plan uploaded"}"
// @Failure      400 {object}   sys.ErrorResponse  "Неверный запрос"
// @Failure      500 {object}   sys.ErrorResponse  "Ошибка сервера при сохранении"
// @Router       /floor-plan/{floor} [post]
func (api *API) UploadFloorPlan(c *gin.Context) {
	// Парсим этаж
	floorParam := c.Param("floor")
	n, err := strconv.ParseInt(floorParam, 10, 64)
	if err != nil {
		sys.HandleError(c, err)
		return
	}

	// Забираем файл, но не читаем его полностью — передаём reader
	fileHeader, err := c.FormFile("file")
	if err != nil {
		sys.HandleError(c, err)
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		sys.HandleError(c, err)
		return
	}
	defer file.Close()

	// Передаём файл-создатель (io.Reader) в сервис
	if err := api.premiseService.SaveFloorPlan(c.Request.Context(), n, file); err != nil {
		sys.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "floor plan uploaded"})
}

// GetFloorPlan отдает SVG‑план этажа
// @Summary      Получить SVG‑план этажа
// @Description  Возвращает raw SVG‑контент для указанного этажа
// @Tags         floor-plans
// @Accept       json
// @Produce      image/svg+xml
// @Param        floor path int64 true "Номер этажа"
// @Success      200 {file}    string              "SVG‑план этажа"
// @Failure      400 {object}  sys.ErrorResponse   "Неверный номер этажа"
// @Failure      404 {object}  sys.ErrorResponse   "План не найден"
// @Router       /floor-plan/{floor} [get]
func (api *API) GetFloorPlan(c *gin.Context) {
	floorParam := c.Param("floor")
	floor, err := strconv.ParseInt(floorParam, 10, 64)
	if err != nil {
		sys.HandleError(c, err)
		return
	}

	// Получаем байты SVG из сервиса
	content, err := api.premiseService.GetFloorPlanContent(c.Request.Context(), floor)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Отдаём как SVG
	c.Data(http.StatusOK, "image/svg+xml", content)
}

// AddPolygonHandler сохраняет полигон
// @Summary      Добавить полигон для помещения
// @Description  Сохраняет новые координаты и подпись полигона
// @Tags         premises
// @Accept       json
// @Produce      json
// @Param        code path     int64                   true  "Код помещения"
// @Param        poly body     model.PremisePolygon  true  "Данные полигона"
// @Success      201  {object}  model.PremisePolygon
// @Failure      400  {object}  sys.ErrorResponse
// @Failure      500  {object}  sys.ErrorResponse
// @Router       /premise/{code}/polygons [post]
func (api *API) AddPolygon(c *gin.Context) {
	codeParam := c.Param("code")
	code, err := strconv.ParseInt(codeParam, 10, 64)
	if err != nil {
		sys.HandleError(c, err)
		return
	}

	var poly model.PremisePolygon
	if err := c.ShouldBindJSON(&poly); err != nil {
		sys.HandleError(c, err)
		return
	}
	poly.PremiseCode = code

	if err := api.premiseService.AddPolygon(c.Request.Context(), &poly); err != nil {
		sys.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, poly)
}

// GetPolygonsHandler возвращает все полигоны по помещению
// @Summary      Получить полигоны помещения
// @Description  Возвращает список полигонов (points+label) для указанного помещения
// @Tags         premises
// @Accept       json
// @Produce      json
// @Success      200  {array}   model.PremisePolygon
// @Failure      400  {object}  sys.ErrorResponse
// @Failure      500  {object}  sys.ErrorResponse
// @Router       /premise/polygons [get]
func (api *API) GetPolygons(c *gin.Context) {

	polys, err := api.premiseService.GetPolygons(c.Request.Context())
	if err != nil {
		sys.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, polys)
}
