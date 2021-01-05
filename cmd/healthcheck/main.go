package main

import (
	"github.com/Pranc1ngPegasus/lambda-chi/internal/adapter/configuration"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	configuration.Load()
	lambda.Start(initialize().Get)
}
