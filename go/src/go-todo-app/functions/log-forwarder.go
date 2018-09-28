package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/sha1sum/aws_signing_client"
	"gopkg.in/olivere/elastic.v3"
)

func newElasticClient(creds *credentials.Credentials) (*elastic.Client, error) {
	signer := v4.NewSigner(creds)
	awsClient, err := aws_signing_client.New(signer, nil, "es", "us-east-1")
	if err != nil {
		return nil, err
	}
	return elastic.NewClient(
		elastic.SetURL("https://my-aws-endpoint.us-east-1.es.amazonaws.com"),
		elastic.SetScheme("https"),
		elastic.SetHttpClient(awsClient),
		elastic.SetSniff(false), // See note below
	)
}

type Event struct {
	AwsLogs string `json:"awslogs"`
}

type AwsLogs struct {
	Data string `json:"data"`
}


func LogHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Request body received: %s", string(request.Body))

	var awsLogs Event
	json.Unmarshal([]byte(request.Body), &awsLogs)

	fmt.Println("Found item: ", awsLogs)
	return events.APIGatewayProxyResponse{
		Body: string(awsLogs.AwsLogs),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(PostHandler)
}