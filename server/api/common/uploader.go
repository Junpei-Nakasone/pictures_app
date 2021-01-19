package common

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// UploadObject S3にデータをアップロード
func UploadObject(filename string) (resp *s3.PutObjectOutput) {
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
