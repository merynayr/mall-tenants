package floorplan

import (
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/service"
	"github.com/merynayr/mall-tenants/internal/sys"
	"github.com/merynayr/mall-tenants/internal/sys/codes"
)

// API premise структура
type API struct {
	floorplanService service.FloorPlanService
}

// NewAPI возвращает новый объект имплементации API-слоя mall
func NewAPI(floorplanService service.FloorPlanService) *API {
	return &API{
		floorplanService: floorplanService,
	}
}

// RegisterRoutes регистрирует маршруты
func (api *API) RegisterRoutes(router *gin.Engine) {
	floorGroup := router.Group("/floor-plan")
	{
		floorGroup.POST("/:floor", api.UploadFloorPlan)
		floorGroup.GET("/:floor", api.GetFloorPlan)
		floorGroup.POST("/polygons/:code", api.AddPolygon)
		floorGroup.GET("/polygons/:floor", api.GetPolygons)
		floorGroup.DELETE("polygons/:code", api.DeletePolygon)
	}
}

// UploadFloorPlan загружает план этажа (SVG или PNG)
// @Summary      Загрузить план этажа
// @Description  Принимает multipart/form-data с полем file и сохраняет файл на диск + в БД
// @Tags         floor-plans
// @Accept       multipart/form-data
// @Produce      json
// @Security     BearerAuth
// @Param        floor path     int  true  "Номер этажа"
// @Param        file  formData file true  "Файл плана этажа (.svg или .png)"
// @Success      201 {object}   map[string]string  "{"message":"floor plan uploaded"}"
// @Failure      400 {object}   sys.ErrorResponse  "Неверный запрос"
// @Failure      500 {object}   sys.ErrorResponse  "Ошибка сервера при сохранении"
// @Router       /floor-plan/{floor} [post]
func (api *API) UploadFloorPlan(c *gin.Context) {
	floorParam := c.Param("floor")
	n, err := strconv.ParseInt(floorParam, 10, 64)
	if err != nil {
		sys.HandleError(c, sys.Wrap(err, sys.InvalidRequestError))
		return
	}

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
	defer func() {
		if err := file.Close(); err == nil {
			sys.HandleError(c, err)
			return
		}
	}()
	// Определим тип файла
	buffer := make([]byte, 512)
	if _, err := file.Read(buffer); err != nil {
		sys.HandleError(c, err)
		return
	}
	filetype := http.DetectContentType(buffer)

	// Поддерживаемые типы
	if filetype != "image/svg+xml" && filetype != "image/png" {
		sys.HandleError(c, sys.NewCommonError("допустимы только SVG или PNG", codes.BadRequest))
		return
	}

	// Вернуть начало файла
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		sys.HandleError(c, err)
		return
	}

	if err := api.floorplanService.SaveFloorPlan(c.Request.Context(), n, file, filetype); err != nil {
		sys.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "floor plan uploaded"})
}

// GetFloorPlan отдает PNG-план этажа
// @Summary      Получить PNG‑план этажа
// @Description  Возвращает raw PNG‑контент для указанного этажа
// @Tags         floor-plans
// @Accept       json
// @Produce      image/png
// @Security BearerAuth
// @Param        floor path int64 true "Номер этажа"
// @Success      200 {file}    string              "PNG‑план этажа"
// @Failure      400 {object}  sys.ErrorResponse   "Неверный номер этажа"
// @Failure      404 {object}  sys.ErrorResponse   "План не найден"
// @Router       /floor-plan/{floor} [get]
func (api *API) GetFloorPlan(c *gin.Context) {
	floorParam := c.Param("floor")
	floor, err := strconv.ParseInt(floorParam, 10, 64)
	if err != nil {
		sys.HandleError(c, sys.Wrap(err, sys.InvalidRequestError))
		return
	}

	content, err := api.floorplanService.GetFloorPlanContent(c.Request.Context(), floor)
	if err != nil {
		sys.HandleError(c, err)
		return
	}

	c.Data(http.StatusOK, "image/png", content)
}

// AddPolygon сохраняет полигон
// @Summary      Добавить полигон для помещения
// @Description  Сохраняет новые координаты и подпись полигона
// @Tags         floor-plans
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Param        code path     int64                   true  "Код помещения"
// @Param        poly body     model.PremisePolygon  true  "Данные полигона"
// @Success      201  {object}  model.PremisePolygon
// @Failure      400  {object}  sys.ErrorResponse
// @Failure      500  {object}  sys.ErrorResponse
// @Router       /floor-plan/polygons/{code} [post]
func (api *API) AddPolygon(c *gin.Context) {
	codeParam := c.Param("code")
	code, err := strconv.ParseInt(codeParam, 10, 64)
	if err != nil {
		sys.HandleError(c, sys.Wrap(err, sys.InvalidRequestError))
		return
	}

	var poly model.PremisePolygon
	if err := c.ShouldBindJSON(&poly); err != nil {
		sys.HandleError(c, sys.Wrap(err, sys.InvalidRequestError))
		return
	}
	poly.PremiseCode = code

	if err := api.floorplanService.AddPolygon(c.Request.Context(), &poly); err != nil {
		sys.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, poly)
}

// GetPolygons возвращает все полигоны по помещению
// @Summary      Получить полигоны помещения
// @Description  Возвращает список полигонов (points+label) для указанного помещения
// @Tags         floor-plans
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Param        floor path int64 true "Номер этажа"
// @Success      200  {array}   model.PremisePolygon
// @Failure      400  {object}  sys.ErrorResponse
// @Failure      500  {object}  sys.ErrorResponse
// @Router       /floor-plan/polygons/{floor} [get]
func (api *API) GetPolygons(c *gin.Context) {
	floorParam := c.Param("floor")
	floor, err := strconv.ParseInt(floorParam, 10, 64)
	if err != nil {
		sys.HandleError(c, sys.Wrap(err, sys.InvalidRequestError))
		return
	}

	polys, err := api.floorplanService.GetPolygons(c.Request.Context(), floor)
	if err != nil {
		sys.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, polys)
}

// DeletePolygon возвращает все полигоны по помещению
// @Summary      Удалить полигоны по code помещения
// @Description  Удаляет полигоны для указанного помещения
// @Tags         floor-plans
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Param        code path int64  true  "Код помещения"
// @Success      200
// @Failure      400  {object}  sys.ErrorResponse
// @Failure      500  {object}  sys.ErrorResponse
// @Router       /floor-plan/polygons/{code} [DELETE]
func (api *API) DeletePolygon(c *gin.Context) {
	codeParam := c.Param("code")
	code, err := strconv.ParseInt(codeParam, 10, 64)
	if err != nil {
		sys.HandleError(c, sys.Wrap(err, sys.InvalidRequestError))
		return
	}

	err = api.floorplanService.DeletPolygon(c.Request.Context(), code)
	if err != nil {
		sys.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, "")
}
