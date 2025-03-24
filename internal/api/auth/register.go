package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/sys"
)

// Register регистрация и получение JWT-токена
// @Summary Регистрация
// @Description Регистрирует пользователя и возвращает JWT-токен
// @Tags user
// @Accept  json
// @Produce  json
// @Param body body model.RegisterRequest true "Данные для регистрации"
// @Success 200 {object} model.AuthResponse
// @Failure 400 {object} sys.ErrorResponse
// @Failure 401 {object} sys.ErrorResponse
// @Failure 500 {object} sys.ErrorResponse
// @Router /api/register [post]
func (a *API) Register(c *gin.Context) {
	var req model.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		sys.HandleError(c, sys.InvalidRequestError)
		return
	}

	authResponse, err := a.authService.Register(c.Request.Context(), req)
	if err != nil {
		sys.HandleError(c, err)
		return
	}

	a.setCookies(c, authResponse.RefreshToken, authResponse.AccessToken)

	c.JSON(http.StatusOK, authResponse)
}
