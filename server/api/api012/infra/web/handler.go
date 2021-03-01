package web

import (
	"net/http"
	"pictures_app/api/api012/domain"
	"pictures_app/api/api012/usecase"

	"github.com/labstack/echo"
)

type handler struct {
	s usecase.Service
}

// Handler インターフェース
type Handler interface {
	DeletePicture(c echo.Context) error
}

// NewHandler Handlerを初期化する
func NewHandler(s usecase.Service) Handler {
	return &handler{s: s}
}

// DeletePicture 画像を削除する
func (h *handler) DeletePicture(c echo.Context) error {

	param := domain.RequestParam{}

	if err := c.Bind(&param); err != nil {
		return err
	}

	err := h.s.DeletePicture(param)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusAccepted)
}
