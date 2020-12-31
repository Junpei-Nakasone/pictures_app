package waf

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// NewEcho Echoアプリを作成
func NewEcho() *echo.Echo {
	echo := echo.New()

	echo.Use(middleware.CORS())

	return echo
}
