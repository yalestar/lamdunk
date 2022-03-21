package main

import (
    "log"
    
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
)

type SomeResponse struct {
    EventName   string `json:"event_name"`
    EventSource string `json:"event_source"`
}

func handleRequest(s3Event events.S3Event) (SomeResponse, error) {
    var resp SomeResponse
    for _, record := range s3Event.Records {
        log.Printf("EVENT NAME: %s\n", record.EventName)
        log.Printf("EVENT SOURCE: %s\n", record.EventSource)
        log.Printf("BUCKET: %s", record.S3.Bucket)
        log.Printf("KEY: %s", record.S3.Object.Key)
        log.Printf("Size: %d", record.S3.Object.Size)
        resp = SomeResponse{
            EventName:   record.EventName,
            EventSource: record.EventSource,
        }
    }
    return resp, nil
    
}

func main() {
    lambda.Start(handleRequest)
}
