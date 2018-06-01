package main

import (
	"github.com/SrcHndWng/go-learning-serverless-apps/chapter5/picturePostSite/functions/models"
	"github.com/SrcHndWng/go-learning-serverless-apps/chapter5/picturePostSite/functions/utils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler gets item by id.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	item, err := models.GetItem(request.PathParameters["id"])
	if err != nil {
		return utils.ErrorResponse(err)
	}

	return utils.ItemResponse(item)
}

func main() {
	lambda.Start(Handler)
}
