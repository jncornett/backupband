package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jncornett/backupband/backupband"
)

func main() {
	lambda.Start(backupband.Handler)
}
