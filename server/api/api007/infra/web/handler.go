package web

import (
	"net/http"
	"pictures_app/api/api007/domain"
	"pictures_app/api/api007/usecase"

	"github.com/labstack/echo"
)

type handler struct {
	s usecase.Service
}

type Handler interface {
	Login(c echo.Context) error
}

func NewHandler(s usecase.Service) Handler {
	return &handler{s: s}
}

// Login ログイン用API
func (h *handler) Login(c echo.Context) error {
	param := domain.RequestParam{}

	if err := c.Bind(&param); err != nil {
		return err
	}

	userdata, err := h.s.FetchUserData(param)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, userdata)
}
