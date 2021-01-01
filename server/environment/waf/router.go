package waf

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// NewRouter returns router
func NewRouter(e *echo.Echo, h Handler) {

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "PONG")
	})

	e.GET("sample", h.api001.SampleAPI)
	e.GET("fetchLatestImages", h.api002.FetchImages)
	e.POST("addImage", h.api006.AddImage)
	e.POST("login", h.api007.Login)
	e.POST("addNewUser", h.api008.AddNewUser)

	e.Use(middleware.CORS())

	// return e
}
