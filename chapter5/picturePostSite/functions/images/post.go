package main

import (
	"encoding/json"
	"fmt"

	"github.com/SrcHndWng/go-learning-serverless-apps/chapter5/picturePostSite/functions/utils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Body contains request body data.
type Body struct {
	Type string `json:"type"`
	Size int    `json:"size"`
}

// Handler gets Image post requests.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)

	var body Body
	err := json.Unmarshal([]byte(request.Body), &body)
	if err != nil {
		return utils.ErrorResponse(err)
	}

	photoID := utils.GenerateID()
	fmt.Println("photoID = " + photoID)

	return events.APIGatewayProxyResponse{Body: "Images Post Success.\n", StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
