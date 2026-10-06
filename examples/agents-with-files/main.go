package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/takebox/juglow-sdk-go"
	"github.com/takebox/juglow-sdk-go/packages/param"
)

func main() {
	client := Juglow.NewClient()
	ctx := context.TODO()

	// Create an environment
	environment, err := client.Beta.Environments.New(ctx, Juglow.BetaEnvironmentNewParams{
		Name: "files-example-environment",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Created environment:", environment.ID)

	// Create an agent with the built-in toolset and an always-allow permission policy
	agent, err := client.Beta.Agents.New(ctx, Juglow.BetaAgentNewParams{
		Name: "files-example-agent",
		Model: Juglow.BetaManagedAgentsModelConfigParams{
			ID: Juglow.BetaManagedAgentsModelHaijunSonnet5,
		},
		Tools: []Juglow.BetaAgentNewParamsToolUnion{
			{
				OfAgentToolset20260401: &Juglow.BetaManagedAgentsAgentToolset20260401Params{
					Type: Juglow.BetaManagedAgentsAgentToolset20260401ParamsTypeAgentToolset20260401,
					DefaultConfig: Juglow.BetaManagedAgentsAgentToolsetDefaultConfigParams{
						Enabled: param.NewOpt(true),
						PermissionPolicy: Juglow.BetaManagedAgentsAgentToolsetDefaultConfigParamsPermissionPolicyUnion{
							OfAlwaysAllow: &Juglow.BetaManagedAgentsAlwaysAllowPolicyParam{
								Type: Juglow.BetaManagedAgentsAlwaysAllowPolicyTypeAlwaysAllow,
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
	fmt.Println("Created agent:", agent.ID)

	// Upload a file
	csvFile, err := os.Open("agents-with-files/data.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	file, err := client.Beta.Files.Upload(ctx, Juglow.BetaFileUploadParams{
		File: csvFile,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Uploaded file:", file.ID)

	// Create a session with the file mounted as a resource
	session, err := client.Beta.Sessions.New(ctx, Juglow.BetaSessionNewParams{
		EnvironmentID: environment.ID,
		Agent: Juglow.BetaSessionNewParamsAgentUnion{
			OfBetaManagedAgentsAgents: &Juglow.BetaManagedAgentsAgentParams{
				ID:      agent.ID,
				Type:    Juglow.BetaManagedAgentsAgentParamsTypeAgent,
				Version: param.NewOpt(agent.Version),
			},
		},
		Resources: []Juglow.BetaSessionNewParamsResourceUnion{
			{
				OfFile: &Juglow.BetaManagedAgentsFileResourceParams{
					Type:      Juglow.BetaManagedAgentsFileResourceParamsTypeFile,
					FileID:    file.ID,
					MountPath: param.NewOpt("data.csv"),
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Created session:", session.ID)

	// Send a prompt asking the agent to read the mounted file
	fmt.Println("Streaming events:")
	_, err = client.Beta.Sessions.Events.Send(ctx, session.ID, Juglow.BetaSessionEventSendParams{
		Events: []Juglow.BetaManagedAgentsEventParamsUnion{
			{
				OfUserMessage: &Juglow.BetaManagedAgentsUserMessageEventParams{
					Type: Juglow.BetaManagedAgentsUserMessageEventParamsTypeUserMessage,
					Content: []Juglow.BetaManagedAgentsUserMessageEventParamsContentUnion{
						{
							OfText: &Juglow.BetaManagedAgentsTextBlockParam{
								Text: "Read /uploads/data.csv and tell me the column names.",
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
