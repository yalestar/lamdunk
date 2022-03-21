**Commands:**

**Start the localstack container for Lambda and S3:**
`docker-compose up`

**Build and zip the task code:**
```shell
GOOS=linux go build -o bin/task && rm -f task.zip && zip -j task.zip bin/task 
```

**Create the Lambda function in Localstack, passing it the zipped task code:**

```shell
aws --endpoint-url=http://localhost:4566 lambda create-function \
--function-name=task --runtime="go1.x" \
--role=fakerole --handler=task \
--zip-file fileb://task.zip
```

**List Lambda functions:**

```shell
aws --endpoint-url=http://localhost:4566 lambda list-functions
```

**Inspect Lambda function:**
```shell
aws --endpoint-url=http://localhost:4566 lambda get-function --function-name task
```

**Create S3 bucket in/on Localstack:**
```shell
aws s3 mb s3://zapp --region us-west-2 --endpoint-url=http://localhost:4566 
```

**List S3 buckets:**
```shell
aws --endpoint-url=http://localhost:4566 s3api list-buckets
```

**Configure the S3 bucket to publish s3:ObjectCreated events:**
```shell
aws s3api put-bucket-notification-configuration --bucket zapp \
--notification-configuration file://test_data/notification.json \
--endpoint-url http://localhost:4566
```

**Confirm the existence of the above notification configuration:**
```shell
aws s3api get-bucket-notification-configuration \
--bucket zapp --endpoint-url http://localhost:4566
```

**Copy a file to S3 bucket on Localstack:**
```shell
aws s3 cp styx_van.jpg s3://zapp --endpoint-url http://localhost:4566
```
Boo. Doesn't work. Doesn't seem to work for anybody yet.

**Invoke the Lambda function, passing it a JSON payload:**
```shell
aws lambda --endpoint-url=http://localhost:4566 invoke \
--function-name task --payload file://test_data/s3event.json \
--region=us-west-2 out.log
```

**Update a Lambda function, passing it the updated zipped task code:**

```shell
aws --endpoint-url=http://localhost:4566 lambda \
update-function-code --function-name=task \
--zip-file fileb://task.zip
```