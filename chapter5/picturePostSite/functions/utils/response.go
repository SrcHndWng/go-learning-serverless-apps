package utils

import (
	"encoding/json"
	"fmt"

	"github.com/SrcHndWng/go-learning-serverless-apps/chapter5/picturePostSite/functions/models"
	"github.com/aws/aws-lambda-go/events"
)

// ErrorResponse returns error response to API Gateway.
func ErrorResponse(err error) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("%+v\n", err)
	return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Internal Server Error!!"}, nil
}

// ItemResponse retuns response with one item data.
func ItemResponse(item models.Item) (events.APIGatewayProxyResponse, error) {
	jsonItem, err := json.Marshal(item)
	if err != nil {
		return ErrorResponse(err)
	}
	return SuccessResponse(string(jsonItem))
}

// ItemsResponse retuns response with items data.
func ItemsResponse(items []models.Item) (events.APIGatewayProxyResponse, error) {
	jsonItem, err := json.Marshal(items)
	if err != nil {
		return ErrorResponse(err)
	}
	return SuccessResponse(string(jsonItem))
}

// SuccessResponse returns response to API Gateway.
func SuccessResponse(body string) (events.APIGatewayProxyResponse, error) {
	headers := map[string]string{
		"Access-Control-Allow-Methods": "*",
		"Access-Control-Allow-Headers": "*",
		"Access-Control-Allow-Origin":  "*",
	}
	return events.APIGatewayProxyResponse{Headers: headers, Body: body, StatusCode: 200}, nil
}

// NoContentResponse returns no content status.
func NoContentResponse() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{StatusCode: 204}, nil
}

// NotFountResponse return not found response.
func NotFountResponse() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{StatusCode: 404}, nil
}
