package web

import (
	"net/http"
	"pictures_app/api/api002/usecase"

	"github.com/labstack/echo"
)

type handler struct {
	s usecase.Service
}

type Handler interface {
	FetchImages(c echo.Context) error
}

func NewHandler(s usecase.Service) Handler {
	return &handler{s: s}
}

// FetchImages is made for sample.
func (h *handler) FetchImages(c echo.Context) error {

	result, err := h.s.FetchImages()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}
