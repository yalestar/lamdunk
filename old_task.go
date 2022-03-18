package main

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"os"

// 	"github.com/aws/aws-lambda-go/events"
// 	runtime "github.com/aws/aws-lambda-go/lambda"
// 	"github.com/aws/aws-lambda-go/lambdacontext"
// 	"github.com/aws/aws-sdk-go/aws/session"
// 	"github.com/aws/aws-sdk-go/service/lambda"
// )

// var client = lambda.New(session.New())

// func callLambda() (string, error) {
// 	input := &lambda.GetAccountSettingsInput{}
// 	req, resp := client.GetAccountSettingsRequest(input)
// 	err := req.Send()
// 	output, _ := json.Marshal(resp.AccountUsage)
// 	return string(output), err
// }

// func handleRequest(ctx context.Context, event events.S3Event) (string, error) {
// 	eventJSON, _ := json.MarshalIndent(event, "", " ")
// 	fmt.Printf("EVENT: %s", eventJSON)

// 	fmt.Printf("REGION: %s", os.Getenv("AWS_REGION"))

// 	fmt.Println("ALL ENV VARS:")
// 	for _, element := range os.Environ() {
// 		fmt.Println("OH! ", element)
// 	}

// 	lc, _ := lambdacontext.FromContext(ctx)
// 	fmt.Printf("REQUEST ID: %s\n", lc.AwsRequestID)
// 	fmt.Printf("FUNCTIONNAME: %s\n", lambdacontext.FunctionName)

// 	deadline, _ := ctx.Deadline()
// 	fmt.Printf("DEADLINE: %s\n", deadline)

// 	// usage, err := callLambda()

// 	// if err != nil {
// 	// 	return "ERROR", err
// 	// }

// 	return "THE WORD USAAGE", nil

// }

// func main() {
// 	runtime.Start(handleRequest)
// }
