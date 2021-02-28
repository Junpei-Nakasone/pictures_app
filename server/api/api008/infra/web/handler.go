package web

import (
	"io"
	"net/http"
	"os"
	"pictures_app/api/common"
	"time"

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

	file, err := c.FormFile("icon_image")
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

	userName := c.FormValue("user_name")
	password := c.FormValue("password")
	emailAddress := c.FormValue("email_address")
	note := c.FormValue("note")

	// intUserID, _ := strconv.Atoi(userID)

	param := domain.RequestParam{
		UserName:     &userName,
		Password:     &password,
		EmailAddress: &emailAddress,
		Note:         &note,
		IconImage:    &uploadedFileName,
	}

	res, err := h.s.AddNewUser(param)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)

}
