package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type SomeEvent struct {
	Name string
	Age  int
}

type SomeResponse struct {
	Message string `json:"Answer"`
	Extra   string `json:"Extra"`
	Extra2  string `json:"Extra2"`
}

func handleRequest(ctx context.Context, event events.S3Event) (SomeResponse, error) {
	event.Records
	response := SomeResponse{
		Message: fmt.Sprintf("%s is %d years old", event.Name, event.Age),
	}
	return response, nil

}

func main() {
	lambda.Start(handleRequest)
}
