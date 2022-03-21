package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambdacontext"
)

func handleARequest(ctx context.Context, event events.S3Event) (string, error) {
    eventJSON, _ := json.MarshalIndent(event, "", " ")
    fmt.Printf("EVENT: %s", eventJSON)
    
    fmt.Printf("REGION: %s", os.Getenv("AWS_REGION"))
    
    fmt.Println("ALL ENV VARS:")
    for _, element := range os.Environ() {
        fmt.Println("OH! ", element)
    }
    
    lc, _ := lambdacontext.FromContext(ctx)
    fmt.Printf("REQUEST ID: %s\n", lc.AwsRequestID)
    fmt.Printf("FUNCTIONNAME: %s\n", lambdacontext.FunctionName)
    
    deadline, _ := ctx.Deadline()
    fmt.Printf("DEADLINE: %s\n", deadline)
    
    return "THE WORD USAAGE", nil
    
}

// func main() {
//     lambda.Start(handleARequest)
// }
