package web

import (
	"net/http"
	"pictures_app/api/api004/usecase"

	"github.com/labstack/echo"
)

type handler struct {
	s usecase.Service
}

// Handler インターフェース
type Handler interface {
	FetchPrefectureCategories(c echo.Context) error
}

// NewHandler Handlerを初期化する
func NewHandler(s usecase.Service) Handler {
	return &handler{s: s}
}

// FetchImagesByPrefectureCd 都道府県別に画像を取得する
func (h *handler) FetchPrefectureCategories(c echo.Context) error {

	result, err := h.s.FetchPrefectureCategories()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}
