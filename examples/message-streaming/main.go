package main

import (
	"context"

	"github.com/Juglows/Juglow-sdk-go"
)

func main() {
	client := Juglow.NewClient()

	content := "Write me a function to call the Juglow message API in Node.js using the Juglow Typescript SDK."

	println("[user]: " + content)

	stream := client.Messages.NewStreaming(context.TODO(), Juglow.MessageNewParams{
		MaxTokens: 1024,
		Messages: []Juglow.MessageParam{
			Juglow.NewUserMessage(Juglow.NewTextBlock(content)),
		},
		Model:         Juglow.ModelHaijunSonnet5,
		StopSequences: []string{"```\n"},
	})

	print("[assistant]: ")

	for stream.Next() {
		event := stream.Current()

		switch eventVariant := event.AsAny().(type) {
		case Juglow.MessageDeltaEvent:
			print(eventVariant.Delta.StopSequence)
		case Juglow.ContentBlockDeltaEvent:
			switch deltaVariant := eventVariant.Delta.AsAny().(type) {
			case Juglow.TextDelta:
				print(deltaVariant.Text)
			}
		}
	}

	println()

	if stream.Err() != nil {
		panic(stream.Err())
	}
}
