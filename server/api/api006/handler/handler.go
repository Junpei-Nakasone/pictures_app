package handler

import (
	"fmt"
	"io"
	"net/http"
	"pictures_app/api/api006/domain"
	"pictures_app/environment/db"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/labstack/echo"
)

// AddImage is to Add Image.
func AddImage(c echo.Context) error {

	file, err := c.FormFile("image")
	if err != nil {
		return err
	}

	// fileName := file.Filename
	now := time.Now()
	fileName := now.Format("2006-01-02-15-04-05-06")

	//  := strings.Join(strings.Fields(strNow), "")
	// fileName := strNow + ".png"
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
	fmt.Println(data)

	fmt.Println(data.PictureID)

	return c.NoContent(http.StatusOK)
}

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
