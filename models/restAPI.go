package models

import "github.com/aws/aws-lambda-go/events"

type RestApi struct {
	Status 		int
	Message 	string
	CustomResp *events.APIGatewayProxyResponse
}