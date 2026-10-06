package main

import (
	"context"
	"github.com/takebox/juglow-sdk-go"
	"github.com/takebox/juglow-sdk-go/vertex"
)

func main() {
	client := Juglow.NewClient(
		vertex.WithGoogleAuth(context.Background(), "us-central1", "id-xxx"),
	)

	content := "Write me a function to call the Juglow message API in Node.js using the Juglow Typescript SDK."

	println("[user]: " + content)

	stream := client.Messages.NewStreaming(context.TODO(), Juglow.MessageNewParams{
		MaxTokens: 1024,
		Messages: []Juglow.MessageParam{
			Juglow.NewUserMessage(Juglow.NewTextBlock(content)),
		},
		Model:         "haijun-sonnet-4-v1@20250514",
		StopSequences: []string{"```\n"},
	})

	print("[assistant]: ")

	for stream.Next() {
		event := stream.Current()

		switch variant := event.AsAny().(type) {
		case Juglow.ContentBlockDeltaEvent:
			if variant.Delta.Text != "" {
				print(variant.Delta.Text)
			}
		case Juglow.MessageDeltaEvent:
			if variant.Delta.StopSequence != "" {
				print(variant.Delta.StopSequence)
			}
		}

	}

	println()

	if stream.Err() != nil {
		panic(stream.Err())
	}
}
