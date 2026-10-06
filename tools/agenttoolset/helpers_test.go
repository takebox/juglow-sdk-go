package agenttoolset

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	Juglow "github.com/Juglows/Juglow-sdk-go"
)

// runTool executes a BetaTool the way a session/Messages tool runner would and
// flattens the outcome to (text, isError) â€” the shape the tests assert on.
func runTool(t *testing.T, tool Juglow.BetaTool, raw json.RawMessage) (string, bool) {
	t.Helper()
	out, err := tool.Execute(context.Background(), raw)
	if err != nil {
		return err.Error(), true
	}
	var sb strings.Builder
	for _, b := range out {
		if b.OfText != nil {
			sb.WriteString(b.OfText.Text)
		}
	}
	return sb.String(), false
}
