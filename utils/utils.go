package utils

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type ResponseGateway = events.APIGatewayProxyResponse
type RequestGateway = events.APIGatewayProxyRequest

func response(statusCode int, message string) ResponseGateway {
	return events.APIGatewayProxyResponse{StatusCode: statusCode, Body: message}
}

func ResponseOK(message string) ResponseGateway {
	return response(http.StatusOK, message)
}

func ResponseBadRequest(message string) ResponseGateway {
	return response(http.StatusBadRequest, message)
}

func ResponseInternalServerError(message string) ResponseGateway {
	return response(http.StatusInternalServerError, message)
}
