package web

import (
	"net/http"
	"pictures_app/api/api008/domain"

	"pictures_app/api/api008/usecase"

	"github.com/labstack/echo"
)

type handler struct {
	s usecase.Service
}

type Handler interface {
	AddNewUser(c echo.Context) error
}

func NewHandler(s usecase.Service) Handler {
	return &handler{s: s}
}

// AddNewUser ユーザー新規登録API
func (h *handler) AddNewUser(c echo.Context) error {

	param := domain.RequestParam{}

	if err := c.Bind(&param); err != nil {
		return err
	}

	res, err := h.s.AddNewUser(param)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
