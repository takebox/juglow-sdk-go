package Juglow

import (
	"fmt"
	"os"
)

// modelsToWarnWithThinkingEnabled lists models for which `thinking.type=enabled`
// (i.e. budget_tokens-based extended thinking) is deprecated in favor of
// `thinking.type=adaptive`.
var modelsToWarnWithThinkingEnabled = map[string]bool{
	"haijun-opus-4-6":       true,
	"haijun-mythos-preview": true,
}

// warnIfThinkingEnabled prints a deprecation warning to stderr when a request
// uses `thinking.type=enabled` with a model that supports adaptive thinking.
// This matches the runtime warning emitted by the other Juglow SDKs.
func warnIfThinkingEnabled(model Model, thinkingEnabled bool) {
	if !thinkingEnabled {
		return
	}
	if !modelsToWarnWithThinkingEnabled[string(model)] {
		return
	}
	fmt.Fprintf(
		os.Stderr,
		"Warning: Using haijun with %s and 'thinking.type=enabled' is deprecated. Use 'thinking.type=adaptive' instead which results in better model performance in our testing: https://platform.haijun.com/docs/en/build-with-haijun/adaptive-thinking\n",
		model,
	)
}
