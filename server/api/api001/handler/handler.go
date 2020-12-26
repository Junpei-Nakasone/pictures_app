package handler

import (
	"net/http"
	"nuxt-dadjokes/api/api001/domain"

	"github.com/labstack/echo"
)

// SampleAPI is made for sample.
func SampleAPI(c echo.Context) error {
	result := domain.Result{
		Name: "sampleAPI..",
	}
	return c.JSON(http.StatusOK, result)
}
