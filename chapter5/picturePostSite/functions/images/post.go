package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler gets Image post requests.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)
	return events.APIGatewayProxyResponse{Body: "Images Post Success.\n", StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
