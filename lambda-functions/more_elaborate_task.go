package main

import (
    "fmt"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "log"
)

const (
    keyName          = "2021-10-29-NCPDP_Monthly_Master_20211001.zip"
    bucketName       = "ncpdp"
    internalEndpoint = "http://172.17.0.1:4566"
)

func doTheNeedful() (string, error) {
    s3cfg, err := getS3Config(internalEndpoint)
    if err != nil {
        log.Println(err)
    }
    
    s3Client := getS3Client(s3cfg)
    
    log.Printf("-------- looking at %s in %s", keyName, bucketName)
    globject, err := getObject(*s3Client, bucketName, keyName)
    
    log.Println("--------- GLOBJECT?")
    log.Println(">>>>>>>>>>", *globject.ContentType)
    
    if err != nil {
        log.Println(err)
    }
    
    return "good", nil
}

func handler(event events.S3Event) error {
    // ASS-ume a file named 2021-10-29-NCPDP_Monthly_Master_20211001.zip
    // is in S3 bucket named ncpdp running on localstack with an
    // internalEndpoint localhost:4566
    
    good, err := doTheNeedful()
    
    if err != nil {
        log.Println(err)
    }
    fmt.Println(good)
    return err
}

func main() {
    lambda.Start(handler)
}
