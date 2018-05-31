package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/SrcHndWng/go-learning-serverless-apps/chapter5/picturePostSite/functions/models"
	"github.com/SrcHndWng/go-learning-serverless-apps/chapter5/picturePostSite/functions/utils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Body contains request body data.
type Body struct {
	Type string `json:"type"`
	Size int    `json:"size"`
}

func timestamp() int64 {
	return time.Now().Unix()
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

	url, err := getPresignedURL(bucket, photoID, body.Type)
	if err != nil {
		return utils.ErrorResponse(err)
	}

	item := models.Item{ID: photoID, Timestamp: timestamp(), Status: "Waiting", Type: body.Type, Size: body.Size, SignedURL: url}
	if err = models.SaveItem(item); err != nil {
		return utils.ErrorResponse(err)
	}

	return utils.ItemResponse(item)
}

func main() {
	lambda.Start(Handler)
}
