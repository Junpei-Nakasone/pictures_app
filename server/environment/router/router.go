package router

import (
	"net/http"

	api001 "pictures_app/api/api001/handler"
	api002 "pictures_app/api/api002/handler"
	api003 "pictures_app/api/api003/handler"
	api004 "pictures_app/api/api004/handler"
	api005 "pictures_app/api/api005/handler"
	api006 "pictures_app/api/api006/handler"
	api007 "pictures_app/api/api007/infra/web"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// NewRouter returns router
func NewRouter() *echo.Echo {
	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "PONG")
	})

	e.GET("sample", api001.SampleAPI)
	e.GET("fetchLatestImages", api002.FetchImages)
	e.POST("addJoke", api003.AddJoke)
	e.PUT("updateJoke", api004.UpdateJoke)
	e.DELETE("deleteJoke", api005.DeleteJoke)
	e.POST("addImage", api006.AddImage)
	e.POST("login", api007.Login)

	e.Use(middleware.CORS())

	return e
}
