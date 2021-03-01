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
	api005data "pictures_app/api/api005/infra/data"
	api005handler "pictures_app/api/api005/infra/web"
	api005service "pictures_app/api/api005/usecase"
	api006data "pictures_app/api/api006/infra/data"
	api006handler "pictures_app/api/api006/infra/web"
	api006service "pictures_app/api/api006/usecase"
	api007data "pictures_app/api/api007/infra/data"
	api007handler "pictures_app/api/api007/infra/web"
	api007service "pictures_app/api/api007/usecase"
	api008data "pictures_app/api/api008/infra/data"
	api008handler "pictures_app/api/api008/infra/web"
	api008service "pictures_app/api/api008/usecase"
	api009data "pictures_app/api/api009/infra/data"
	api009handler "pictures_app/api/api009/infra/web"
	api009service "pictures_app/api/api009/usecase"
	api010data "pictures_app/api/api010/infra/data"
	api010handler "pictures_app/api/api010/infra/web"
	api010service "pictures_app/api/api010/usecase"
	api011data "pictures_app/api/api011/infra/data"
	api011handler "pictures_app/api/api011/infra/web"
	api011service "pictures_app/api/api011/usecase"
	api012data "pictures_app/api/api012/infra/data"
	api012handler "pictures_app/api/api012/infra/web"
	api012service "pictures_app/api/api012/usecase"

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
	serviceRepository005 := api005data.NewServiceRepository(d)
	servive005 := api005service.NewService(serviceRepository005)
	handler005 := api005handler.NewHandler(servive005)
	serviceRepository006 := api006data.NewServiceRepository(d)
	servive006 := api006service.NewService(serviceRepository006)
	handler006 := api006handler.NewHandler(servive006)
	serviceRepository007 := api007data.NewServiceRepository(d)
	servive007 := api007service.NewService(serviceRepository007)
	handler007 := api007handler.NewHandler(servive007)
	serviceRepository008 := api008data.NewServiceRepository(d)
	servive008 := api008service.NewService(serviceRepository008)
	handler008 := api008handler.NewHandler(servive008)
	serviceRepository009 := api009data.NewServiceRepository(d)
	servive009 := api009service.NewService(serviceRepository009)
	handler009 := api009handler.NewHandler(servive009)
	serviceRepository010 := api010data.NewServiceRepository(d)
	servive010 := api010service.NewService(serviceRepository010)
	handler010 := api010handler.NewHandler(servive010)
	serviceRepository011 := api011data.NewServiceRepository(d)
	servive011 := api011service.NewService(serviceRepository011)
	handler011 := api011handler.NewHandler(servive011)
	serviceRepository012 := api012data.NewServiceRepository(d)
	servive012 := api012service.NewService(serviceRepository012)
	handler012 := api012handler.NewHandler(servive012)

	wafHandler := waf.NewHandler(
		handler001,
		handler002,
		handler003,
		handler004,
		handler005,
		handler006,
		handler007,
		handler008,
		handler009,
		handler010,
		handler011,
		handler012,
	)

	return wafHandler
}
