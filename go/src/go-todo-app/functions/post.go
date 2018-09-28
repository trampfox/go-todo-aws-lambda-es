package main

import (
	"fmt"
	"log"
	"encoding/json"

	"go-todo-app/dao"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type PostTask struct {
	Content string `json:"content"`
	Category string `json:"category"`
}

func PostHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Request body received: %s", string(request.Body))

	var postTask PostTask
	json.Unmarshal([]byte(request.Body), &postTask)

	task, err := dao.SaveTask(postTask.Content, postTask.Category)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("%s", string(err.Error())),
			StatusCode: 500,
		}, nil
	}

	jsonItem, _ := json.Marshal(task)
	stringItem := string(jsonItem) + "\n"

	fmt.Println("Found item: ", stringItem)
	return events.APIGatewayProxyResponse{
		Body: stringItem,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(PostHandler)
}
