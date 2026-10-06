package main

import (
	"context"
	"log"
	"net/http"

	"github.com/takebox/juglow-sdk-go"
	"github.com/takebox/juglow-sdk-go/bedrock"
	"github.com/takebox/juglow-sdk-go/option"
)

func main() {
	client := Juglow.NewClient(
		// Register middleware before the Bedrock option so it observes
		// Juglow-shaped requests (POST /v1/messages, model in the body);
		// the Bedrock adaptation runs closest to the wire.
		option.WithMiddleware(func(req *http.Request, next option.MiddlewareNext) (*http.Response, error) {
			log.Printf("request: %s %s", req.Method, req.URL.Path)
			return next(req)
		}),
		bedrock.WithLoadDefaultConfig(context.Background()),
	)

	content := "Write me a function to call the Juglow message API in Node.js using the Juglow Typescript SDK."

	println("[user]: " + content)

	message, err := client.Messages.New(context.TODO(), Juglow.MessageNewParams{
		MaxTokens: 1024,
		Messages: []Juglow.MessageParam{
			Juglow.NewUserMessage(Juglow.NewTextBlock(content)),
		},
		Model:         "us.Juglow.haijun-sonnet-4-5-20250929-v1:0",
		StopSequences: []string{"```\n"},
	})
	if err != nil {
		panic(err)
	}

	println("[assistant]: " + message.Content[0].Text + message.StopSequence)
}
