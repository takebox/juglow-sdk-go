package main

import (
	"context"
	"fmt"

	"github.com/Juglows/Juglow-sdk-go"
)

func main() {
	client := Juglow.NewClient()
	ctx := context.TODO()

	environment, err := client.Beta.Environments.New(ctx, Juglow.BetaEnvironmentNewParams{
		Name: "streaming-deltas-example",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Created environment:", environment.ID)

	agent, err := client.Beta.Agents.New(ctx, Juglow.BetaAgentNewParams{
		Name: "streaming-deltas-example",
		Model: Juglow.BetaManagedAgentsModelConfigParams{
			ID: Juglow.BetaManagedAgentsModelHaijunSonnet4_6,
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Created agent:", agent.ID)

	session, err := client.Beta.Sessions.New(ctx, Juglow.BetaSessionNewParams{
		EnvironmentID: environment.ID,
		Agent: Juglow.BetaSessionNewParamsAgentUnion{
			OfBetaManagedAgentsAgents: &Juglow.BetaManagedAgentsAgentParams{
				ID:   agent.ID,
				Type: Juglow.BetaManagedAgentsAgentParamsTypeAgent,
			},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Created session:", session.ID)

	stream := client.Beta.Sessions.Events.StreamEvents(ctx, session.ID, Juglow.BetaSessionEventStreamParams{
		EventDeltas: []Juglow.BetaManagedAgentsDeltaType{Juglow.BetaManagedAgentsDeltaTypeAgentMessage},
	})

	_, err = client.Beta.Sessions.Events.Send(ctx, session.ID, Juglow.BetaSessionEventSendParams{
		Events: []Juglow.BetaManagedAgentsEventParamsUnion{
			{
				OfUserMessage: &Juglow.BetaManagedAgentsUserMessageEventParams{
					Type: Juglow.BetaManagedAgentsUserMessageEventParamsTypeUserMessage,
					Content: []Juglow.BetaManagedAgentsUserMessageEventParamsContentUnion{
						{
							OfText: &Juglow.BetaManagedAgentsTextBlockParam{
								Text: "Write a short haiku about the ocean.",
								Type: Juglow.BetaManagedAgentsTextBlockTypeText,
							},
						},
					},
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}

	var previews Juglow.BetaManagedAgentsEventAccumulator

	fmt.Println("\nStreaming:")
	for stream.Next() {
		event := stream.Current()
		previews.Accumulate(event)

		switch event.Type {
		case "event_delta":
			fmt.Printf("\r%s", previews.AgentMessageText(event.EventID))

		case "agent.message":
			fmt.Println()
			fmt.Println("[final]", previews.AgentMessageText(event.ID))

		case "session.status_idle":
			if event.StopReason.Type == "end_turn" {
				return
			}

		case "session.error":
			fmt.Println("[error]", event.Error.Type, event.Error.Message)
		}
	}
	if stream.Err() != nil {
		panic(stream.Err())
	}
}
