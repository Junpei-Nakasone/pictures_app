package di

import (
	api001handler "pictures_app/api/api001/infra/web"
	api002data "pictures_app/api/api002/infra/data"
	api002handler "pictures_app/api/api002/infra/web"
	api002service "pictures_app/api/api002/usecase"
	api003data "pictures_app/api/api003/infra/data"
	api003handler "pictures_app/api/api003/infra/web"
	api003service "pictures_app/api/api003/usecase"
	api004data "pictures_app/api/api004/infra/data"
	api004handler "pictures_app/api/api004/infra/web"
	api004service "pictures_app/api/api004/usecase"
	api006data "pictures_app/api/api006/infra/data"
	api006handler "pictures_app/api/api006/infra/web"
	api006service "pictures_app/api/api006/usecase"
	api007data "pictures_app/api/api007/infra/data"
	api007handler "pictures_app/api/api007/infra/web"
	api007service "pictures_app/api/api007/usecase"
	api008data "pictures_app/api/api008/infra/data"
	api008handler "pictures_app/api/api008/infra/web"
	api008service "pictures_app/api/api008/usecase"
	"pictures_app/environment/waf"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func InitializeHandler(d *gorm.DB, e *echo.Echo) waf.Handler {
	handler001 := api001handler.NewHandler()
	serviceRepository002 := api002data.NewServiceRepository(d)
	servive002 := api002service.NewService(serviceRepository002)
	handler002 := api002handler.NewHandler(servive002)
	serviceRepository003 := api003data.NewServiceRepository(d)
	servive003 := api003service.NewService(serviceRepository003)
	handler003 := api003handler.NewHandler(servive003)
	serviceRepository004 := api004data.NewServiceRepository(d)
	servive004 := api004service.NewService(serviceRepository004)
	handler004 := api004handler.NewHandler(servive004)
	serviceRepository006 := api006data.NewServiceRepository(d)
	servive006 := api006service.NewService(serviceRepository006)
	handler006 := api006handler.NewHandler(servive006)
	serviceRepository007 := api007data.NewServiceRepository(d)
	servive007 := api007service.NewService(serviceRepository007)
	handler007 := api007handler.NewHandler(servive007)
	serviceRepository008 := api008data.NewServiceRepository(d)
	servive008 := api008service.NewService(serviceRepository008)
	handler008 := api008handler.NewHandler(servive008)

	wafHandler := waf.NewHandler(
		handler001,
		handler002,
		handler003,
		handler004,
		handler006,
		handler007,
		handler008,
	)

	return wafHandler
}
