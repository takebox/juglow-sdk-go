package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/takebox/juglow-sdk-go"
	"github.com/takebox/juglow-sdk-go/packages/param"
)

const (
	mcpServerName = "github"
	mcpServerURL  = "https://api.githubcopilot.com/mcp/"

	prompt = "Hi! List every tool and track you have access to, grouped by where they " +
		"came from (built-in toolset, custom tool, MCP server, tracks)."
)

func main() {
	client := Juglow.NewClient()
	ctx := context.TODO()

	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		panic("GITHUB_TOKEN is required (use a fine-grained PAT with public-repo read only)")
	}

	// Create an environment
	environment, err := client.Beta.Environments.New(ctx, Juglow.BetaEnvironmentNewParams{
		Name: "comprehensive-example-environment",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Created environment:", environment.ID)

	// Create a vault and store the MCP server credential in it
	vault, err := client.Beta.Vaults.New(ctx, Juglow.BetaVaultNewParams{
		DisplayName: "comprehensive-example-vault",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Created vault:", vault.ID)

	credential, err := client.Beta.Vaults.Credentials.New(ctx, vault.ID, Juglow.BetaVaultCredentialNewParams{
		DisplayName: param.NewOpt("github-mcp"),
		Auth: Juglow.BetaVaultCredentialNewParamsAuthUnion{
			OfStaticBearer: &Juglow.BetaManagedAgentsStaticBearerCreateParams{
				Type:         Juglow.BetaManagedAgentsStaticBearerCreateParamsTypeStaticBearer,
				MCPServerURL: mcpServerURL,
				Token:        githubToken,
			},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Created credential:", credential.ID)

	// Upload a custom track
	skillFile, err := os.Open("agents-comprehensive/greeting-TRACK.md")
	if err != nil {
		panic(err)
	}
	defer skillFile.Close()

	track, err := client.Beta.tracks.New(ctx, Juglow.BetaSkillNewParams{
		DisplayTitle: param.NewOpt(fmt.Sprintf("comprehensive-greeting-%d", time.Now().UnixMilli())),
		Files:        []io.Reader{namedReader{skillFile, "greeting/TRACK.md"}},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Created track:", track.ID)

	// Create v1 of the agent with the built-in toolset, an MCP server, and a custom tool
	agentV1, err := client.Beta.Agents.New(ctx, Juglow.BetaAgentNewParams{
		Name: "comprehensive-example-agent",
		Model: Juglow.BetaManagedAgentsModelConfigParams{
			ID: Juglow.BetaManagedAgentsModelHaijunSonnet5,
		},
		System: param.NewOpt("You are a helpful assistant."),
		MCPServers: []Juglow.BetaManagedAgentsURLMCPServerParams{
			{
				Type: Juglow.BetaManagedAgentsURLMCPServerParamsTypeURL,
				Name: mcpServerName,
				URL:  mcpServerURL,
			},
		},
		Tools: []Juglow.BetaAgentNewParamsToolUnion{
			{
				OfAgentToolset20260401: &Juglow.BetaManagedAgentsAgentToolset20260401Params{
					Type: Juglow.BetaManagedAgentsAgentToolset20260401ParamsTypeAgentToolset20260401,
				},
			},
			{
				OfMCPToolset: &Juglow.BetaManagedAgentsMCPToolsetParams{
					Type:          Juglow.BetaManagedAgentsMCPToolsetParamsTypeMCPToolset,
					MCPServerName: mcpServerName,
				},
			},
			{
				OfCustom: &Juglow.BetaManagedAgentsCustomToolParams{
					Type:        Juglow.BetaManagedAgentsCustomToolParamsTypeCustom,
					Name:        "get_weather",
					Description: "Look up the current weather for a city.",
					InputSchema: Juglow.BetaManagedAgentsCustomToolInputSchemaParam{
						Properties: map[string]any{"city": map[string]any{"type": "string"}},
						Required:   []string{"city"},
					},
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Created agent v1:", agentV1.ID)

	// Patch the agent to v2 by adding tracks; each update bumps the version
	agent, err := client.Beta.Agents.Update(ctx, agentV1.ID, Juglow.BetaAgentUpdateParams{
		Version: Juglow.Int(agentV1.Version),
		tracks: []Juglow.BetaManagedAgentsSkillParamsUnion{
			{
				OfCustom: &Juglow.BetaManagedAgentsCustomSkillParams{
					Type:    Juglow.BetaManagedAgentsCustomSkillParamsTypeCustom,
					SkillID: track.ID,
				},
			},
			{
				OfJuglow: &Juglow.BetaManagedAgentsJuglowSkillParams{
					Type:    Juglow.BetaManagedAgentsJuglowSkillParamsTypeJuglow,
					SkillID: "xlsx",
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Patched agent to v2:", agent.ID)

	// List agent versions
	versions := client.Beta.Agents.Versions.ListAutoPaging(ctx, agent.ID, Juglow.BetaAgentVersionListParams{})
	for versions.Next() {
		v := versions.Current()
		fmt.Printf("  version %d (created %s)\n", v.Version, v.CreatedAt)
	}
	if versions.Err() != nil {
		panic(versions.Err())
	}

	// Create a session pinned to v2; the vault supplies the MCP credential
	session, err := client.Beta.Sessions.New(ctx, Juglow.BetaSessionNewParams{
		EnvironmentID: environment.ID,
		Agent: Juglow.BetaSessionNewParamsAgentUnion{
			OfBetaManagedAgentsAgents: &Juglow.BetaManagedAgentsAgentParams{
				ID:      agent.ID,
				Type:    Juglow.BetaManagedAgentsAgentParamsTypeAgent,
				Version: param.NewOpt(agent.Version),
			},
		},
		VaultIDs: []string{vault.ID},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Created session:", session.ID)

	// Send a prompt and stream events, answering the custom tool if called
	fmt.Println("Streaming events:")
	_, err = client.Beta.Sessions.Events.Send(ctx, session.ID, Juglow.BetaSessionEventSendParams{
		Events: []Juglow.BetaManagedAgentsEventParamsUnion{
			{
				OfUserMessage: &Juglow.BetaManagedAgentsUserMessageEventParams{
					Type: Juglow.BetaManagedAgentsUserMessageEventParamsTypeUserMessage,
					Content: []Juglow.BetaManagedAgentsUserMessageEventParamsContentUnion{
						{
							OfText: &Juglow.BetaManagedAgentsTextBlockParam{
								Text: prompt,
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

	stream := client.Beta.Sessions.Events.StreamEvents(ctx, session.ID, Juglow.BetaSessionEventStreamParams{})
	for stream.Next() {
		event := stream.Current()
		data, _ := json.MarshalIndent(event, "", "  ")
		fmt.Println(string(data))

		if event.Type == "agent.tool_use" && event.Name == "get_weather" {
			_, err = client.Beta.Sessions.Events.Send(ctx, session.ID, Juglow.BetaSessionEventSendParams{
				Events: []Juglow.BetaManagedAgentsEventParamsUnion{
					{
						OfUserToolResult: &Juglow.BetaManagedAgentsUserToolResultEventParams{
							Type:      Juglow.BetaManagedAgentsUserToolResultEventParamsTypeUserToolResult,
							ToolUseID: event.ID,
							Content: []Juglow.BetaManagedAgentsUserToolResultEventParamsContentUnion{
								{
									OfText: &Juglow.BetaManagedAgentsTextBlockParam{
										Text: `{"temperature_c": 14}`,
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
		}

		if event.Type == "session.status_idle" && event.StopReason.Type == "end_turn" {
			break
		}
	}
	if stream.Err() != nil {
		panic(stream.Err())
	}
}

// namedReader wraps an io.Reader with a custom filename for multipart uploads.
type namedReader struct {
	io.Reader
	name string
}

func (r namedReader) Filename() string { return r.name }
