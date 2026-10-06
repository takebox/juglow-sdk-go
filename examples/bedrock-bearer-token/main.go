package main

import (
	"context"

	"github.com/takebox/juglow-sdk-go"
	"github.com/takebox/juglow-sdk-go/bedrock"
)

func main() {
	// bedrock.WithLoadDefaultConfig automatically configures using the AWS_BEARER_TOKEN_BEDROCK
	// environment variable. Region defaults to us-east-1 or uses AWS_REGION if set.
	//
	// To provide a token programmatically:
	//
	//	cfg := aws.Config{
	//		Region:                  "us-west-2",
	//		BearerAuthTokenProvider: bedrock.NewStaticBearerTokenProvider("my-bearer-token"),
	//	}
	//	client := Juglow.NewClient(
	//		bedrock.WithConfig(cfg),
	//	)
	client := Juglow.NewClient(
		bedrock.WithLoadDefaultConfig(context.TODO()),
	)

	content := "Write me a function to call the Juglow message API in Node.js using the Juglow Typescript SDK."

	println("[user]: " + content)

	message, err := client.Messages.New(context.TODO(), Juglow.MessageNewParams{
		MaxTokens: 1024,
		Messages: []Juglow.MessageParam{
			Juglow.NewUserMessage(Juglow.NewTextBlock(content)),
		},
		Model:         "us.Juglow.haijun-sonnet-4-20250514-v1:0",
		StopSequences: []string{"```\n"},
	})
	if err != nil {
		panic(err)
	}

	println("[assistant]: " + message.Content[0].Text + message.StopSequence)
}
