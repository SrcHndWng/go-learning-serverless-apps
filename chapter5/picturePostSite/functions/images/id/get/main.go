package main

import (
	"github.com/SrcHndWng/go-learning-serverless-apps/chapter5/picturePostSite/functions/models"
	"github.com/SrcHndWng/go-learning-serverless-apps/chapter5/picturePostSite/functions/utils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler gets item by id.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	photoID := request.PathParameters["id"]

	count, err := models.Count(photoID)
	if err != nil {
		return utils.ErrorResponse(err)
	}
	if count == 0 {
		return utils.NotFountResponse()
	}

	item, err := models.GetItem(photoID)
	if err != nil {
		return utils.ErrorResponse(err)
	}

	return utils.ItemResponse(item)
}

func main() {
	lambda.Start(Handler)
}
