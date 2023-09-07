#!/bin/bash

aws --endpoint-url=http://localstack:4566 sqs create-queue --queue-name ${SQS_QUEUE_NAME}
