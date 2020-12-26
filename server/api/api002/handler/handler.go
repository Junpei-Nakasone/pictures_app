package handler

import (
	"net/http"
	"nuxt-dadjokes/api/api002/service"

	"github.com/labstack/echo"
)

// FetchImages is made for sample.
func FetchImages(c echo.Context) error {

	result, err := service.FetchImages()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}
