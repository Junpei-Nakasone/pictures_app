package web

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"pictures_app/api/api006/domain"
	"pictures_app/api/api006/usecase"
	"pictures_app/environment/db"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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

	uploadObject("tmp/" + fileName)

	os.Remove(dstFile.Name())
	uploadedFileName := "https://20201108-bucket2.s3-ap-northeast-1.amazonaws.com/" + fileName

	db := db.CreateDBConnection()
	defer db.Close()

	// param := domain.RequestParam{}

	// param.ImageURL = uploadedFileName

	userID := c.FormValue("user_id")
	intUserID, _ := strconv.Atoi(userID)
	// param.UserID = userID

	// now := time.Now()

	data := domain.Pictures{
		UserID:      intUserID,
		ImageURL:    uploadedFileName,
		PublishedAt: time.Now(),
	}

	err = db.Table("pictures").Create(&data).Error
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

// TODO: api/common配下に移動する
func uploadObject(filename string) (resp *s3.PutObjectOutput) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var (
		s3session *s3.S3
	)
	const (
		BUCKETNAME = "20201108-bucket2"
		REGION     = "ap-northeast-1"
	)
	s3session = s3.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(REGION),
	})))

	fmt.Println("Uploading:", filename)
	resp, err = s3session.PutObject(&s3.PutObjectInput{
		Body:   f,
		Bucket: aws.String(BUCKETNAME),
		Key:    aws.String(strings.Split(filename, "/")[1]),
		ACL:    aws.String(s3.BucketCannedACLPublicRead),
	})

	if err != nil {
		panic(err)
	}
	return resp
}
