package main

import (
	"context"
	"fmt"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/responses"
)

func main() {
	client := openai.NewClient()

	response, err := client.Responses.New(context.Background(), openai.ResponseNewParams{
		Model: "gpt-5.5",
		Input: responses.ResponseNewParamsInputUnion{
			OfString: openai.String("Write a short bedtime story about a unicorn."),
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(response.OutputText())
}
