package main

import (
	"encoding/json"
	"fmt"
	"time"
	"unsafe"

	"github.com/SrcHndWng/go-todo-sample-api-gateway/model"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/guregu/dynamo"
)

// Response has lambda return message.
type Response struct {
	Message string `json:"message"`
}

// KinesisData describes data from kinesis.
type KinesisData struct {
	ID        int64  `json:"id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}

// DynamoData describes data that register to dynamodb.
type DynamoData struct {
	ID        int64  `json:"id"`
	Timestamp int64  `json:"timestamp"`
	Text      string `json:"text"`
}

var tbl dynamo.Table

func bstring(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func register(data *KinesisData) error {
	tbl = table()
	t, err := time.Parse("Mon Jan 2 15:04:05 -0700 2006", data.CreatedAt)
	if err != nil {
		return err
	}
	d := DynamoData{ID: data.ID, Timestamp: t.Unix(), Text: data.Text}
	return tbl.Put(d).Run()
}

func table() dynamo.Table {
	if tbl.Name() == "" {
		return model.Table("tweet-data")
	}
	return tbl
}

// Handler is lambda function handler.
func Handler(request events.KinesisEvent) (Response, error) {
	for _, rec := range request.Records {
		data := new(KinesisData)
		json.Unmarshal(rec.Kinesis.Data, data)
		err := register(data)
		if err != nil {
			fmt.Printf("error raise. %#v", err)
		}
	}

	return Response{
		Message: "kinesisToDynamo success!",
	}, nil
}

func main() {
	lambda.Start(Handler)
}
