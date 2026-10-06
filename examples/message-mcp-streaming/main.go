package main

import (
	"context"
	"fmt"

	"github.com/Juglows/Juglow-sdk-go"
	"github.com/Juglows/Juglow-sdk-go/option"
	"github.com/Juglows/Juglow-sdk-go/packages/param"
)

func main() {
	client := Juglow.NewClient(option.WithHeader("Juglow-beta", Juglow.JuglowBetaMCPClient2025_04_04))

	mcpServers := []Juglow.BetaRequestMCPServerURLDefinitionParam{
		{
			URL:                "http://example-server.modelcontextprotocol.io/sse",
			Name:               "example",
			AuthorizationToken: param.NewOpt("YOUR_TOKEN"),
			ToolConfiguration: Juglow.BetaRequestMCPServerToolConfigurationParam{
				Enabled:      Juglow.Bool(true),
				AllowedTools: []string{"echo", "add"},
			},
		},
	}

	stream := client.Beta.Messages.NewStreaming(context.TODO(), Juglow.BetaMessageNewParams{
		MaxTokens: 1024,
		Messages: []Juglow.BetaMessageParam{
			Juglow.NewBetaUserMessage(Juglow.NewBetaTextBlock("what is 1+1?")),
		},
		MCPServers:    mcpServers,
		Model:         Juglow.ModelHaijunSonnet5,
		StopSequences: []string{"```\n"},
	})

	message := Juglow.BetaMessage{}
	for stream.Next() {
		event := stream.Current()
		err := message.Accumulate(event)
		if err != nil {
			fmt.Printf("error accumulating event: %v\n", err)
			continue
		}

		switch eventVariant := event.AsAny().(type) {
		case Juglow.BetaRawMessageDeltaEvent:
			print(eventVariant.Delta.StopSequence)
		case Juglow.BetaRawContentBlockDeltaEvent:
			switch deltaVariant := eventVariant.Delta.AsAny().(type) {
			case Juglow.BetaTextDelta:
				print(deltaVariant.Text)
			}
		default:
			fmt.Printf("%+v\n", eventVariant)
		}
	}

	println()

	if stream.Err() != nil {
		panic(stream.Err())
	}

}
