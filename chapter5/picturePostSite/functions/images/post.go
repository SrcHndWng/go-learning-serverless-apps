package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/SrcHndWng/go-learning-serverless-apps/chapter5/picturePostSite/functions/utils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Body contains request body data.
type Body struct {
	Type string `json:"type"`
	Size int    `json:"size"`
}

// Item contains data to register to dynamoDB.
type Item struct {
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Status    string `json:"status"`
	Type      string `json:"type"`
	Size      int    `json:"size"`
	SignedURL string `json:"signed_url"`
}

func nowDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func register(item Item) error {
	tbl := utils.Table(os.Getenv("TABLE_NAME"))
	return tbl.Put(item).Run()
}

// Handler gets Image post requests.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)

	var body Body
	if err := json.Unmarshal([]byte(request.Body), &body); err != nil {
		return utils.ErrorResponse(err)
	}

	photoID := utils.GenerateID()
	bucket := os.Getenv("BUCKET_NAME")

	url, err := utils.GetPresignedURL(bucket, photoID, body.Type)
	if err != nil {
		return utils.ErrorResponse(err)
	}

	item := Item{ID: photoID, Timestamp: nowDateTime(), Status: "Waiting", Type: body.Type, Size: body.Size, SignedURL: url}
	if err = register(item); err != nil {
		return utils.ErrorResponse(err)
	}

	jsonItem, err := json.Marshal(item)
	if err != nil {
		return utils.ErrorResponse(err)
	}
	return utils.SuccessResponse(string(jsonItem))
}

func main() {
	lambda.Start(Handler)
}
