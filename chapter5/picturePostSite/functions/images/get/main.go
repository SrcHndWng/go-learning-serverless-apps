package main

import (
	"github.com/SrcHndWng/go-learning-serverless-apps/chapter5/picturePostSite/functions/models"
	"github.com/SrcHndWng/go-learning-serverless-apps/chapter5/picturePostSite/functions/utils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler gets all items.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	items, err := models.GetItemsByStatus("Uploaded")
	if err != nil {
		return utils.ErrorResponse(err)
	}
	return utils.ItemsResponse(items)
}

func main() {
	lambda.Start(Handler)
}
