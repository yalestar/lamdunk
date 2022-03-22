package main

import (
    "context"
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/s3"
    "log"
)

func getS3Config() (aws.Config, error) {
    cfg, err := config.LoadDefaultConfig(
        context.TODO(),
        config.WithEndpointResolverWithOptions(
            aws.EndpointResolverWithOptionsFunc(
                func(service, region string, options ...interface{}) (
                    aws.Endpoint, error,
                ) {
                    return aws.Endpoint{
                        URL: "http://localhost:4566",
                    }, nil
                },
            ),
        ),
    )
    return cfg, err
}

func getS3Client(cfg aws.Config) *s3.Client {
    client := s3.NewFromConfig(
        cfg,
        func(options *s3.Options) {
            options.UsePathStyle = true
            options.EndpointOptions.DisableHTTPS = true
        },
    )
    return client
}

func getObject(client s3.Client, bucket, key string) (*s3.GetObjectOutput, error) {
    log.Printf("LOoking for %s in %s", key, bucketName)
    goi := s3.GetObjectInput{Bucket: &bucket, Key: &key}
    obj, err := client.GetObject(context.TODO(), &goi)
    log.Printf("============what of %v", obj)
    if err != nil {
        log.Printf("============== ERROR UP IN THIS PIG")
        log.Println(err)
        return nil, err
    }
    
    return obj, nil
}
