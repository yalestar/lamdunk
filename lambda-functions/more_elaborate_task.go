package main

import (
    "context"
    "fmt"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/service/s3"
    "log"
)

const (
    keyName    = "minimas.zip"
    bucketName = "ncpdp"
)

func giveHead(client *s3.Client) {
    
    hip := s3.HeadObjectInput{
        Bucket: aws.String(bucketName),
        Key:    aws.String(keyName),
    }
    
    hop, err := client.HeadObject(context.TODO(), &hip)
    if err != nil {
        log.Print("---------- err")
        log.Fatal(err)
    }
    
    log.Println("---------------- HEAD OUTPUT")
    log.Print(hop.ContentType)
    
}

func doTheNeedful() (string, error) {
    
    s3cfg, err := getS3Config()
    if err != nil {
        log.Println(err)
    }
    
    s3Client := getS3Client(s3cfg)
    
    log.Println("-------- GOT CLIENT AND CONFIG")
    // giveHead(s3Client)
    
    log.Printf("-------- looking at %s in %s", keyName, bucketName)
    globject, err := getObject(*s3Client, bucketName, keyName)
    
    log.Println("--------- GLOBJECT?")
    log.Println("-----------", globject.ContentType)
    if err != nil {
        log.Println(err)
    }
    
    log.Print(globject)
    return "good", nil
}

func handler(event events.S3Event) error {
    // Ass-ume a file named minimas.zip is in S3 bucket named ncpdp
    // running on localstack with an endpoint localhost:4566
    
    good, err := doTheNeedful()
    fmt.Println("==================")
    fmt.Println(err)
    fmt.Println("==================")
    fmt.Println(good)
    return err
}

func main() {
    lambda.Start(handler)
}
