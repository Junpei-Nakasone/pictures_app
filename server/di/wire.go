package di

import (
	api002data "pictures_app/api/api002/infra/data"
	api002handler "pictures_app/api/api002/infra/web"
	api002service "pictures_app/api/api002/usecase"
	"pictures_app/environment/waf"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func InitializeHandler(d *gorm.DB, e *echo.Echo) waf.Handler {
	serviceRepository002 := api002data.NewServiceRepository(d)
	servive002 := api002service.NewService(serviceRepository002)
	handler002 := api002handler.NewHandler(servive002)

	wafHandler := waf.NewHandler(handler002)

	return wafHandler
}
