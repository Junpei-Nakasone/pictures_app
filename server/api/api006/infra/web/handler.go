package web

import (
	"io"
	"net/http"
	"os"
	"pictures_app/api/api006/domain"
	"pictures_app/api/api006/usecase"
	"pictures_app/api/common"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

type handler struct {
	s usecase.Service
}

type Handler interface {
	AddImage(c echo.Context) error
}

func NewHandler(s usecase.Service) Handler {
	return &handler{s: s}
}

// AddImage is to Add Image.
func (h *handler) AddImage(c echo.Context) error {

	file, err := c.FormFile("image")
	if err != nil {
		return err
	}

	now := time.Now()
	fileName := now.Format("2006-01-02-15-04-05-06")

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dstFile, err := os.Create("tmp/" + fileName)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	if _, err = io.Copy(dstFile, src); err != nil {
		return err
	}

	// uploadObject("tmp/" + fileName)
	common.UploadObject("tmp/" + fileName)

	os.Remove(dstFile.Name())
	uploadedFileName := "https://20201108-bucket2.s3-ap-northeast-1.amazonaws.com/" + fileName

	userID := c.FormValue("user_id")
	intUserID, _ := strconv.Atoi(userID)
	viewCategoryCd := c.FormValue("view_category_cd")
	prefectureCategoryCd := c.FormValue("prefecture_category_cd")
	imageNote := c.FormValue("image_note")

	data := domain.Picture{
		UserID:               intUserID,
		ImageURL:             uploadedFileName,
		PrefectureCategoryCd: prefectureCategoryCd,
		ViewCategoryCd:       viewCategoryCd,
		ImageNote:            imageNote,
		PublishedAt:          time.Now(),
	}

	err = h.s.AddImage(data)

	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
