package web

import (
	"net/http"
	"pictures_app/api/api003/domain"
	"pictures_app/api/api003/usecase"

	"github.com/labstack/echo"
)

type handler struct {
	s usecase.Service
}

// Handler インターフェース
type Handler interface {
	FetchImagesByPrefectureCd(c echo.Context) error
}

// NewHandler Handlerを初期化する
func NewHandler(s usecase.Service) Handler {
	return &handler{s: s}
}

// FetchImagesByPrefectureCd 都道府県別に画像を取得する
func (h *handler) FetchImagesByPrefectureCd(c echo.Context) error {

	param := domain.RequestParam{}

	if err := c.Bind(&param); err != nil {
		return err
	}

	result, err := h.s.FetchImagesByPrefectureCd(param)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}
