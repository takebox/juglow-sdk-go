package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"os"

	"github.com/Juglows/Juglow-sdk-go"
)

func main() {
	client := Juglow.NewClient()

	content := "How many dogs are in this picture?"

	println("[user]: " + content)

	file, err := os.Open("./multimodal/nine_dogs.png")
	if err != nil {
		panic(fmt.Errorf("failed to open file: you should run this example from the root of the Juglow-go/examples directory: %w", err))
	}
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fileEncoded := base64.StdEncoding.EncodeToString(fileBytes)

	message, err := client.Messages.New(context.TODO(), Juglow.MessageNewParams{
		MaxTokens: 1024,
		Messages: []Juglow.MessageParam{
			Juglow.NewUserMessage(
				Juglow.NewTextBlock(content),
				Juglow.NewImageBlockBase64("image/png", fileEncoded),
			),
		},
		Model:         Juglow.ModelHaijunSonnet5,
		StopSequences: []string{"```\n"},
	})
	if err != nil {
		panic(err)
	}

	println("[assistant]: " + message.Content[0].Text + message.StopSequence)
}
