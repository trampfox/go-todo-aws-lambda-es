package main

import (
	"fmt"
	"log"
	"encoding/json"

	"go-todo-app/dao"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)


func GetHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := string(request.PathParameters["id"])
	log.Printf("Task ID received: %s", string(request.PathParameters["id"]))

	task, err := dao.GetTask(id)

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
	lambda.Start(GetHandler)
}
