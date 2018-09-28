package dao

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/google/uuid"
	"log"
	"os"
)

type Task struct {
	Id string	`json:"id"`
	Content string `json:"content"`
	Category string `json:"category"`
}

func ListTasks() ([]Task, error) {
	tasks := []Task{}

	svc := GetSession()

	filt := expression.Name("category").Equal(expression.Value("default"))
	expr, err := expression.NewBuilder().WithFilter(filt).Build()

	if err != nil {
		log.Fatal(err.Error())
		return []Task{}, err
	}

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(os.Getenv("TABLE_NAME")),
	}

	// Make the DynamoDB Query API call
	result, err := svc.Scan(params)
	fmt.Println("Result", result)

	if err != nil {
		log.Fatal(err.Error())
		return tasks, err
	}

	numItems := 0
	for _, i := range result.Items {
		task := Task{}

		err = dynamodbattribute.UnmarshalMap(i, &task)
		if err != nil {
			log.Fatal(err.Error())
			return tasks, err
		}

		tasks = append(tasks, task)
		numItems++
	}

	return tasks, nil
}

func GetTask(id string) (Task, error) {
	svc := GetSession()

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})

	if err != nil {
		log.Fatal(err.Error())
		return Task{}, err
	}

	task := Task{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &task)

	if err != nil {
		log.Fatal(err.Error())
		return task, err
	}

	return task, nil
}

func SaveTask(content string, category string) (Task, error) {
	svc := GetSession()

	task := Task {
		Id: uuid.New().String(),
		Content: content,
		Category: category,
	}

	item, err := dynamodbattribute.MarshalMap(task)

	if err != nil {
		log.Fatal(err.Error())
		return task, err
	}

	input := &dynamodb.PutItemInput{
		Item: item,
		TableName: aws.String(os.Getenv("TABLE_NAME")),
	}

	_, err = svc.PutItem(input)

	if err != nil {
		log.Fatal(err.Error())
		return task, err
	}

	return task, nil
}

func GetSession() (*dynamodb.DynamoDB) {
	sess := session.Must(session.NewSession())
	svc := dynamodb.New(sess)

	return svc
}