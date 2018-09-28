package main

import (
	"fmt"
	"log"
	"encoding/json"

	"go-todo-app/dao"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func ListHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	tasks, err := dao.ListTasks()

	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to find tasks, %v", err))
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("%s", string(err.Error())),
			StatusCode: 500,
		}, nil
	}

	jsonItem, _ := json.Marshal(tasks)
	stringItem := string(jsonItem) + "\n"


	fmt.Println("Found items: ", stringItem)
	return events.APIGatewayProxyResponse{
		Body: stringItem,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(ListHandler)
}
