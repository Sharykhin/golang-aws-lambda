package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}

type Response struct {
	Message string `json:"message:"`
}

func HandleLambdaEvent(req Request) (Response, error) {
	fmt.Println("test my request", req)
	return Response{Message: fmt.Sprintf("your transaction of %f has been applied, description: %s", req.Amount, req.Description)}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
