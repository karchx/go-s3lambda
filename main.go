package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func response(stringResponse string, statusCode int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{Body: stringResponse, StatusCode: statusCode}
}

func ResponseOK(stringResponse string) events.APIGatewayProxyResponse {
	return response(stringResponse, 200)
}

func ResponseServerError(stringResponse string) events.APIGatewayProxyResponse {
	return response(stringResponse, 500)
}

func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
  log.Println(req)

  return ResponseOK("ready!"), nil
}

func main() {
	lambda.Start(HandleRequest)
}
