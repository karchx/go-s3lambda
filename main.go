package main

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Even struct {
	Number *int `json:"number"`
}

type EvenResponse struct {
	IsEven bool `json:"is_event"`
}

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
	if req.Body == "" {
		return events.APIGatewayProxyResponse{}, errors.New("empty body")
	}

	var body Even
	err := json.Unmarshal([]byte(req.Body), &body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	var check bool

	if *body.Number%2 == 0 {
		check = true
	} else {
		check = false
	}

	res := EvenResponse{
		IsEven: check,
	}

	bytesRes, err := json.Marshal(res)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return ResponseOK(string(bytesRes)), nil
}

func main() {
	lambda.Start(HandleRequest)
}
