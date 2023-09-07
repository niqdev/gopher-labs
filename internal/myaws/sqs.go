package myaws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

func Example() {

}

func defaultSqsClient(ctx context.Context) (*sqs.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("error loading default configuration: %w", err)
	}

	client := sqs.NewFromConfig(cfg)
	return client, err
}

// https://github.com/localstack/localstack-aws-sdk-examples/blob/main/go/s3-basic-v2.go
// https://towardsaws.com/sns-and-sqs-with-localstack-using-golang-16b291f45e0b
func localSqsClient(ctx context.Context, awsEndpoint string, awsRegion string) (*sqs.Client, error) {
	// TODO deprecated
	// TODO catch error ???
	endpointResolver :=
		aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{
				PartitionID:       "aws",
				URL:               awsEndpoint,
				SigningRegion:     awsRegion,
				HostnameImmutable: true,
			}, nil
		})

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(awsRegion),
		config.WithEndpointResolver(endpointResolver),
	)
	if err != nil {
		return nil, fmt.Errorf("error loading local configuration: %w", err)
	}

	client := sqs.NewFromConfig(cfg)
	return client, err
}

func getQueueUrl(ctx context.Context, client *sqs.Client, queueName string) (*string, error) {
	getQueueUrlRequest := &sqs.GetQueueUrlInput{
		QueueName: &queueName,
	}

	urlResult, err := client.GetQueueUrl(ctx, getQueueUrlRequest)
	if err != nil {
		return nil, fmt.Errorf("error getting queue url: %w", err)
	}

	return urlResult.QueueUrl, nil
}

func sendMessage(ctx context.Context, client *sqs.Client, queueUrl string, body string) (*string, error) {
	messageInput := &sqs.SendMessageInput{
		DelaySeconds: 10,
		MessageBody:  aws.String(body),
		QueueUrl:     &queueUrl,
	}

	resp, err := client.SendMessage(ctx, messageInput)
	if err != nil {
		return nil, fmt.Errorf("error sending message: %w", err)
	}
	return resp.MessageId, err
}

// exampleSend https://github.com/awsdocs/aws-doc-sdk-examples/blob/main/gov2/sqs/SendMessage/SendMessagev2.go
func exampleSend() (string, error) {
	// TODO param
	ctx := context.TODO()
	awsEndpoint := "http://localhost:4566"
	awsRegion := "us-east-1"
	queueName := "go-sqs-example"
	messageBody := "hello"

	fmt.Println(fmt.Sprintf("Sending message in queue: %v", queueName))

	client, err := localSqsClient(ctx, awsEndpoint, awsRegion)
	if err != nil {
		return "", err
	}

	queueUrl, err := getQueueUrl(ctx, client, queueName)
	if err != nil {
		return "", err
	}

	messageId, err := sendMessage(ctx, client, *queueUrl, messageBody)
	if err != nil {
		return "", err
	}

	return *messageId, nil
}

func Send() {
	messageIdSent, err := exampleSend()
	if err != nil {
		fmt.Println(fmt.Sprintf("FAILURE send %v", err))
	}
	fmt.Println(fmt.Sprintf("SEND messageId: %v", messageIdSent))
}

func receiveMessage(ctx context.Context, client *sqs.Client, queueUrl string) (string, error) {
	messageInput := &sqs.ReceiveMessageInput{
		MessageAttributeNames: []string{
			string(types.QueueAttributeNameAll),
		},
		QueueUrl:            &queueUrl,
		MaxNumberOfMessages: 1,
		VisibilityTimeout:   10,
	}

	msgResult, err := client.ReceiveMessage(ctx, messageInput)
	if err != nil {
		return "", fmt.Errorf("error receiving messages: %w", err)
	}

	// TODO map collection []string{}
	if msgResult.Messages != nil {
		return *msgResult.Messages[0].MessageId, nil
	} else {
		return "NONE", nil
	}
}

// exampleReceive https://github.com/awsdocs/aws-doc-sdk-examples/blob/main/gov2/sqs/ReceiveMessage/ReceiveMessagev2.go
// TODO loop https://github.com/awsdocs/aws-doc-sdk-examples/blob/main/gov2/sqs/ReceiveLPMessage/ReceiveLPMessagev2.go
func exampleReceive() (string, error) {
	ctx := context.TODO()
	awsEndpoint := "http://localhost:4566"
	awsRegion := "us-east-1"
	queueName := "go-sqs-example"

	fmt.Println(fmt.Sprintf("Receiving message from queue: %v", queueName))

	client, err := localSqsClient(ctx, awsEndpoint, awsRegion)
	if err != nil {
		return "", err
	}

	queueUrl, err := getQueueUrl(ctx, client, queueName)
	if err != nil {
		return "", err
	}

	messageId, err := receiveMessage(ctx, client, *queueUrl)
	if err != nil {
		return "", err
	}

	return messageId, nil
}

func Receive() {
	// TODO ExampleReceiveLoop
	messageIdReceived, err := exampleReceive()
	if err != nil {
		fmt.Println(fmt.Sprintf("FAILURE receive %v", err))
	}
	fmt.Println(fmt.Sprintf("RECEIVE messageId: %v", messageIdReceived))
}
