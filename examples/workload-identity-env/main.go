package main

import (
	"context"

	"github.com/Juglows/Juglow-sdk-go"
)

func main() {
	// Zero-config workload identity authentication via environment variables.
	// Set the following env vars before running:
	//
	//   Juglow_FEDERATION_RULE_ID â€” the federation rule ID
	//   Juglow_ORGANIZATION_ID   â€” the organization ID
	//   Juglow_IDENTITY_TOKEN    â€” a literal JWT identity token
	//     (or Juglow_IDENTITY_TOKEN_FILE â€” path to a file containing the JWT)
	//
	// Optional:
	//   Juglow_SERVICE_ACCOUNT_ID â€” service account ID
	//
	// When these are set, NewClient() automatically exchanges the identity token
	// for a short-lived Juglow access token. If an API key is also set, the
	// API key takes precedence.
	client := Juglow.NewClient()

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
