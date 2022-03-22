HANDLER_NAME=mungo
BUILD_DIR=bin

S3_EVENT_JSON=ncpdpS3Event.json

make_bucket:
	aws s3 mb s3://ncpdp --region us-west-2 --endpoint-url=http://localhost:4566 \
    aws s3 cp minimas.zip s3://ncpdp --endpoint-url=http://localhost:4566
#    aws s3api head-object --bucket ncpdp --key minimas.zip --endpoint-url=http://localhost:4566

build_lambda_function: clean
	GOOS=linux go build -o ${BUILD_DIR}/${HANDLER_NAME}  && \
	zip -j ${BUILD_DIR}/${HANDLER_NAME}.zip ${BUILD_DIR}/${HANDLER_NAME}

deploy_lambda: build_lambda_function
	aws --endpoint-url=http://localhost:4566 lambda create-function \
		--function-name=${HANDLER_NAME} --runtime="go1.x" \
		--role=fakerole --handler=${HANDLER_NAME} \
		--zip-file fileb://${BUILD_DIR}/${HANDLER_NAME}.zip

update_lambda: build_lambda_function
	aws --endpoint-url=http://localhost:4566 lambda \
    update-function-code --function-name=${HANDLER_NAME} \
    --zip-file fileb://${BUILD_DIR}/${HANDLER_NAME}.zip

clean:
	rm -vf ${BUILD_DIR}/${HANDLER_NAME}*
	rm -f out.log

invoke_lambda:
	aws lambda --endpoint-url=http://localhost:4566 invoke \
    --function-name ${HANDLER_NAME} --payload file://${S3_EVENT_JSON} \
    --region=us-west-2 out.log
