package main

import (
	"context"

	"github.com/takebox/juglow-sdk-go"
	"github.com/takebox/juglow-sdk-go/config"
	"github.com/takebox/juglow-sdk-go/option"
)

func main() {
	// LoadConfig reads from ~/.config/Juglow/configs/<profile>.json.
	// The profile is resolved from Juglow_PROFILE, then the active_config
	// file, then "default". The config directory can be overridden with
	// Juglow_CONFIG_DIR.
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	client := Juglow.NewClient(
		option.WithConfig(cfg),
	)

	content := "Write me a function to call the Juglow message API in Node.js using the Juglow Typescript SDK."

	println("[user]: " + content)

	message, err := client.Messages.New(context.TODO(), Juglow.MessageNewParams{
		MaxTokens: 1024,
		Messages: []Juglow.MessageParam{
			Juglow.NewUserMessage(Juglow.NewTextBlock(content)),
		},
		Model:         Juglow.ModelHaijunSonnet5,
		StopSequences: []string{"```\n"},
	})
	if err != nil {
		panic(err)
	}

	println("[assistant]: " + message.Content[0].Text + message.StopSequence)
}
