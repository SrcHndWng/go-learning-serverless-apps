package utils

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

// ErrorResponse returns error response to API Gateway.
func ErrorResponse(err error) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("%+v\n", err)
	return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Internal Server Error!!"}, nil
}
