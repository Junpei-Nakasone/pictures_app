package environment

import (
	"pictures_app/di"
	"pictures_app/environment/waf"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type Application struct {
	App *echo.Echo
}

func NewApp(db *gorm.DB) *Application {
	echo := waf.NewEcho()
	handler := di.InitializeHandler(db, echo)

	waf.NewRouter(echo, handler)

	return &Application{
		App: echo,
	}
}
