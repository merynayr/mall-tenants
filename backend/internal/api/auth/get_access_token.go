package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merynayr/mall-tenants/internal/sys"
)

// GetAccessToken обрабатывает HTTP-запрос на получение access токена
func (a *API) GetAccessToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		sys.HandleError(c, sys.Wrap(err, sys.InvalidRequestError))
		return
	}

	accessToken, err := a.authService.GetAccessToken(c.Request.Context(), refreshToken)
	if err != nil {
		sys.HandleError(c, err)
		return
	}

	refreshToken, err = a.authService.GetRefreshToken(c.Request.Context(), refreshToken)
	if err != nil {
		sys.HandleError(c, sys.InvalidRefreshTokenError)
		return
	}

	a.setCookies(c, refreshToken, accessToken)

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken, "refresh_token": refreshToken})
}
