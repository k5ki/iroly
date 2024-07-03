package main

import (
	"iroly/app/infra"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	handler := infra.NewHandler()
	if os.Getenv("DISABLE_LAMBDA") == "1" {
		handler.DebugRunEcho()
	} else {
		lambda.Start(handler.HandleRequest)
	}
}
