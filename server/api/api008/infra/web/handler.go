package web

import (
	"net/http"
	"pictures_app/api/api008/domain"
	"pictures_app/api/api008/service"

	"github.com/labstack/echo"
)

// AddNewUser ユーザー新規登録API
func AddNewUser(c echo.Context) error {

	param := domain.RequestParam{}

	if err := c.Bind(&param); err != nil {
		return err
	}

	res, err := service.AddNewUser(param)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
