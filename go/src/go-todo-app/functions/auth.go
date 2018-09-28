package main

import (
	"context"
	"encoding/base64"
	"errors"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func generatePolicy(principalId, effect, resource string) events.APIGatewayCustomAuthorizerResponse {
	authResponse := events.APIGatewayCustomAuthorizerResponse{PrincipalID: principalId}

	if effect != "" && resource != "" {
		authResponse.PolicyDocument = events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   effect,
					Resource: []string{resource},
				},
			},
		}
	}

	return authResponse
}


func AuthHandler(ctx context.Context, event events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	auth := strings.SplitN(event.AuthorizationToken, " ", 2)

	if len(auth) != 2 || auth[0] != "Basic" {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized")
	}

	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)

	if len(pair) != 2 || !IsAuthorized(pair[0], pair[1]) {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized")
	}

	return generatePolicy("user", "Allow", event.MethodArn), nil
}

func IsAuthorized(username string, password string) (bool) {
	if strings.Compare(username, string(os.Getenv("BASIC_AUTH_USERNAME"))) != 0 {
		log.Printf("Invalid username received: %s", username)
		return false
	}

	if strings.Compare(password, string(os.Getenv("BASIC_AUTH_PASSWORD"))) != 0 {
		log.Printf("Bad password")
		return false
	}

	return true
}

func main() {
	lambda.Start(AuthHandler)
}
