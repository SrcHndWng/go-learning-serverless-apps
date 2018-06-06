package main

import (
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func getPresignedURL(bucket string, photoID string, reqType string) (url string, err error) {
	ext := strings.Split(reqType, "/")[1]
	key := photoID + "." + ext
	svc := s3.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))

	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	url, err = req.Presign(time.Minute * 3600)
	return
}
