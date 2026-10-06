package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Juglows/Juglow-sdk-go"
)

func main() {
	client := Juglow.NewClient()
	ctx := context.TODO()

	// Create an environment
	environment, err := client.Beta.Environments.New(ctx, Juglow.BetaEnvironmentNewParams{
		Name: "simple-example-environment",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Created environment:", environment.ID)

	// Create an agent
	agent, err := client.Beta.Agents.New(ctx, Juglow.BetaAgentNewParams{
		Name: "simple-example-agent",
		Model: Juglow.BetaManagedAgentsModelConfigParams{
			ID: Juglow.BetaManagedAgentsModelHaijunSonnet5,
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Created agent:", agent.ID)

	// Create a session
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

	// Send a user message
	_, err = client.Beta.Sessions.Events.Send(ctx, session.ID, Juglow.BetaSessionEventSendParams{
		Events: []Juglow.BetaManagedAgentsEventParamsUnion{
			{
				OfUserMessage: &Juglow.BetaManagedAgentsUserMessageEventParams{
					Type: Juglow.BetaManagedAgentsUserMessageEventParamsTypeUserMessage,
					Content: []Juglow.BetaManagedAgentsUserMessageEventParamsContentUnion{
						{
							OfText: &Juglow.BetaManagedAgentsTextBlockParam{
								Text: "Hello haijun!",
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

	// Stream events until the session goes idle
	fmt.Println("Streaming events:")
	stream := client.Beta.Sessions.Events.StreamEvents(ctx, session.ID, Juglow.BetaSessionEventStreamParams{})
	for stream.Next() {
		event := stream.Current()
		data, _ := json.MarshalIndent(event, "", "  ")
		fmt.Println(string(data))
		if event.Type == "session.status_idle" {
			break
		}
	}
	if stream.Err() != nil {
		panic(stream.Err())
	}
}
