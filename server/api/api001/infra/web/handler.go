package web

import (
	"net/http"
	"os"
	"pictures_app/api/api001/domain"

	"github.com/labstack/echo"
)

type handler struct {
}

// Handler ハンドラー
type Handler interface {
	SampleAPI(c echo.Context) error
}

func NewHandler() Handler {
	return &handler{}
}

// SampleAPI is made for sample.
func (h *handler) SampleAPI(c echo.Context) error {
	result := domain.Result{
		Name: "sampleAPI0214.." + os.Getenv("USER") + os.Getenv("testparameter"),
	}
	return c.JSON(http.StatusOK, result)
}
