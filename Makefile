HANDLER_NAME=task
BUILD_DIR=bin
S3_EVENT_JSON=test_data/s3event.json

build_lambda_function: clean
	GOOS=linux go build -o ${BUILD_DIR}/${HANDLER_NAME}  && \
	zip -j ${HANDLER_NAME}.zip ${BUILD_DIR}/${HANDLER_NAME}

create_lambda: build_lambda_function
	aws --endpoint-url=http://localhost:4566 lambda create-function \
		--function-name=${HANDLER_NAME} --runtime="go1.x" \
		--role=fakerole --handler=${HANDLER_NAME} \
		--zip-file fileb://${HANDLER_NAME}.zip

rebuild_lambda: build_lambda_function
	aws --endpoint-url=http://localhost:4566 lambda \
    update-function-code --function-name=${HANDLER_NAME} \
    --zip-file fileb://${HANDLER_NAME}.zip

clean:
	rm -vf ${BUILD_DIR}/${HANDLER_NAME}*

invoke_lambda:
	aws lambda --endpoint-url=http://localhost:4566 invoke \
    --function-name ${HANDLER_NAME} --payload file://${S3_EVENT_JSON} \
    --region=us-west-2 out.log
