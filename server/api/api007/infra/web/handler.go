package web

import (
	"net/http"
	"nuxt-dadjokes/api/api007/domain"
	"nuxt-dadjokes/api/api007/service"

	"github.com/labstack/echo"
)

// Login ログイン用API
func Login(c echo.Context) error {
	param := domain.RequestParam{}

	if err := c.Bind(&param); err != nil {
		return err
	}

	userdata, err := service.FetchUserData(param)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, userdata)
}
