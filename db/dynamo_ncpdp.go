package db

import (
    "context"
    "log"
    
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
    NcpdpTable = "Ncpdp"
    // DynamoDBEndpoint = viper.GetString("DYNAMODB_ENDPOINT")
    DynamoDBEndpoint = "http://localhost:4566"
    internalEndpoint = "http://172.17.0.1:4566"
)

// var client *dynamodb.Client
//
// func init() {
//     cfg, err := config.LoadDefaultConfig(context.Background())
//     if err != nil {
//         log.Fatal(err)
//     }
//     log.Println("AT THIS MOMENT WHAT IS DYNAMODB_ENDPOINT", DynamoDBEndpoint)
//     client = dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
//         o.EndpointResolver = dynamodb.EndpointResolverFromURL(DynamoDBEndpoint)
//     })
// }

func connecto(endpoint string) *dynamodb.Client {
    var client *dynamodb.Client
    cfg, err := config.LoadDefaultConfig(context.Background())
    if err != nil {
        log.Fatal(err)
    }
    client = dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
        o.EndpointResolver = dynamodb.EndpointResolverFromURL(endpoint)
    })
    
    return client
}
func DescribeTable() (string, error) {
    svc := connecto(internalEndpoint)
    log.Printf("Asking as to the presence of %s", NcpdpTable)
    input := dynamodb.DescribeTableInput{TableName: aws.String(NcpdpTable)}
    out, err := svc.DescribeTable(context.Background(), &input)
    log.Println("OUT: ", out)
    if err != nil {
        log.Println("A ERROR IS BORN", err)
        return "unavailable", err
    }
    // TODO: return boolean maybe
    return "ok", nil
}

// func GetNcpdpItem(ncpdpProviderId string) (map[string]types.AttributeValue, error) {
//     log.Println("Fetching", ncpdpProviderId)
//     svc := client
//     out, err := svc.GetItem(context.Background(), &dynamodb.GetItemInput{
//         TableName: aws.String(NcpdpTable),
//         Key: map[string]types.AttributeValue{
//             "ncpdp_provider_id": &types.AttributeValueMemberS{Value: ncpdpProviderId},
//         },
//     })
//     if err != nil {
//         log.Fatal(err)
//     }
//     return out.Item, nil
// }
//
// func CreateNcpdpTable() {
//     log.Println("CREATING TABLE", NcpdpTable)
//     svc := client
//     _, err := svc.CreateTable(context.Background(), &dynamodb.CreateTableInput{
//         BillingMode: types.BillingModePayPerRequest,
//         AttributeDefinitions: []types.AttributeDefinition{
//             {
//                 AttributeName: aws.String("ncpdp_provider_id"),
//                 AttributeType: types.ScalarAttributeTypeS,
//             },
//         },
//         KeySchema: []types.KeySchemaElement{
//             {
//                 AttributeName: aws.String("ncpdp_provider_id"),
//                 KeyType:       types.KeyTypeHash,
//             },
//         },
//         TableName: aws.String(NcpdpTable),
//     })
//     if err != nil {
//         log.Fatal(err)
//     }
// }
//
// func PutNcpdpItem(ncpdp *models.Ncpdp) {
//     svc := client
//
//     // convert our Ncpdp model object into the format for inserting into DynamoDB
//     // (i.e. a map of strings to AttributeValues)
//     marshalledNcpdp, err := av.MarshalMap(ncpdp)
//     putItemInput := dynamodb.PutItemInput{
//         TableName: aws.String(NcpdpTable),
//         Item:      marshalledNcpdp,
//     }
//
//     _, err = svc.PutItem(context.Background(), &putItemInput)
//
//     if err != nil {
//         log.Fatal(err)
//     }
// }
//
// func DropNcpdpTable() {
//     log.Println("DROPPING TABLE", NcpdpTable)
//     svc := client
//     _, err := svc.DeleteTable(context.Background(), &dynamodb.DeleteTableInput{
//         TableName: aws.String(NcpdpTable),
//     })
//
//     if err != nil {
//         var oe *smithy.OperationError
//         if errors.As(err, &oe) {
//             log.Println("SMITHY ERROR")
//         }
//     }
// }
