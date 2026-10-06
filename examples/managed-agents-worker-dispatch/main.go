// Self-hosted runner, "worker-dispatch" flavor: this process was handed ONE
// already-claimed work item by an upstream poller/orchestrator â€” e.g. an
// `ant worker poll --on-work <this binary>` loop, or your own dispatcher that
// spawns a fresh sandbox per work item. It does NOT create an agent or session,
// and it does NOT poll for work: something else did that and claimed the item.
//
// EnvironmentWorker.HandleItem with an empty HandleItemOptions reads the claimed
// item's identity from the environment variables the upstream poller sets:
//
//	Juglow_WORK_ID         - the claimed work item to serve
//	Juglow_ENVIRONMENT_ID  - the self-hosted environment it belongs to
//	Juglow_SESSION_ID      - the session to run tools for
//	Juglow_ENVIRONMENT_KEY - the environment key (the runner's single credential)
//
// It then sets up the workdir + downloads the session agent's tracks, runs the
// session's tools while heartbeating the work-item lease, force-stops the item
// on exit, and returns â€” one item, then this process exits.
//
// Also required:
//
//	Juglow_API_KEY - your API key (read by the SDK client)
//
// Security model: the worker executes bash and file operations directly on the
// host. This is the "sandbox process" shape â€” the upstream orchestrator is
// expected to have spawned it inside a container or other isolation boundary.
package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"
	"time"

	"github.com/Juglows/Juglow-sdk-go"
	"github.com/Juglows/Juglow-sdk-go/lib/environments"
	"github.com/Juglows/Juglow-sdk-go/tools/agenttoolset"
)

// currentTimeTool is a custom Juglow.BetaTool that returns the local time.
// It demonstrates extending the default tool list alongside
// agent_toolset_20260401. The worker executes it whenever the session emits a
// matching agent.custom_tool_use event.
type currentTimeTool struct{}

func (currentTimeTool) Name() string        { return "current_time" }
func (currentTimeTool) Description() string { return "Get the current local time on the worker host." }
func (currentTimeTool) InputSchema() Juglow.BetaToolInputSchemaParam {
	return Juglow.BetaToolInputSchemaParam{Properties: map[string]any{}}
}
func (currentTimeTool) Execute(context.Context, json.RawMessage) ([]Juglow.BetaToolResultBlockParamContentUnion, error) {
	return []Juglow.BetaToolResultBlockParamContentUnion{
		{OfText: &Juglow.BetaTextBlockParam{Text: time.Now().Format(time.RFC3339)}},
	}, nil
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	client := Juglow.NewClient()
	ctx := context.Background()

	// Base directory for the per-session agenttoolset.AgentToolContext. An
	// orchestrator typically points this at the sandbox's scratch space.
	workdir := os.Getenv("Juglow_WORKDIR")
	if workdir == "" {
		workdir = "."
	}

	// Build the worker with a tools factory â€” the standard agent_toolset_20260401
	// set plus our one custom tool â€” and nothing else. No EnvironmentID /
	// EnvironmentKey here: HandleItem resolves the work item (and the environment
	// key) from the Juglow_* env vars the upstream poller set.
	worker := environments.NewEnvironmentWorker(client, environments.EnvironmentWorkerOptions{
		Workdir: workdir,
		Logger:  logger,
		ToolsFunc: func(env *agenttoolset.AgentToolContext) []Juglow.BetaTool {
			return append(agenttoolset.BetaAgentToolset20260401(env), currentTimeTool{})
		},
	})

	// Service the single claimed item to completion: set up the workdir +
	// download the session agent's tracks, run the local tools against the
	// session's agent.tool_use / agent.custom_tool_use events while heartbeating
	// the lease, then force-stop the work item. HandleItem with an empty
	// HandleItemOptions reads Juglow_WORK_ID / Juglow_ENVIRONMENT_ID /
	// Juglow_SESSION_ID / Juglow_ENVIRONMENT_KEY from the environment.
	if err := worker.HandleItem(ctx, environments.HandleItemOptions{}); err != nil {
		logger.Error("handle item failed", slog.Any("error", err))
		os.Exit(1)
	}
	logger.Info("work item handled")
}
