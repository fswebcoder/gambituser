package main

import (
	"context"
	"errors"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/fswebcoder/awsgo"
)

func main() {
	lambda.Start(ExecuteLambda)
}

func ExecuteLambda(ctx context.Context, events events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.StarAws()
	if !ValidateParameters() {
		err := errors.New("SecretName parameter is not set")
		return events, err
	}
}

func ValidateParameters() bool {
	var getParameter bool
	_, getParameter = os.LookupEnv("SecretName")
	return getParameter
}
