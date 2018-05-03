package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	uuid "github.com/satori/go.uuid"
)

func main() {
	const region = "ap-northeast-1"
	partitionKey := uuid.NewV4().String()
	const streamName = "cloudwatch-alarm-sample-stream"
	data := time.Now().Format("2006-01-02 15:04:05")

	s := session.New(&aws.Config{Region: aws.String(region)})
	c := kinesis.New(s)

	for i := 0; i < 15; i++ {
		out, err := c.PutRecord(&kinesis.PutRecordInput{
			Data:         []byte(data),
			StreamName:   aws.String(streamName),
			PartitionKey: aws.String(partitionKey),
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(out)
	}

	fmt.Println("put record end.")
}
