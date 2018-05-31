package main

import (
	"encoding/json"
	"fmt"

	"github.com/SrcHndWng/go-learning-serverless-apps/chapter5/picturePostSite/functions/models"
	"github.com/SrcHndWng/go-learning-serverless-apps/chapter5/picturePostSite/functions/utils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Body contains request body data.
type Body struct {
	ID        string `json:"id"`
	Timestamp int64  `json:"timestamp"`
	Status    string `json:"status"`
}

// Handler gets status update requests.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)

	var body Body
	if err := json.Unmarshal([]byte(request.Body), &body); err != nil {
		return utils.ErrorResponse(err)
	}

	if err := models.UpdateItem(body.ID, body.Timestamp, body.Status); err != nil {
		return utils.ErrorResponse(err)
	}

	item, err := models.GetItem(body.ID)
	if err != nil {
		return utils.ErrorResponse(err)
	}

	return utils.ItemResponse(item)
}

func main() {
	lambda.Start(Handler)
}
