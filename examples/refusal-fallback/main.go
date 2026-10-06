// Falls back to a second model when the primary refuses, two ways:
// server-side via the fallbacks param (preferred), and client-side via
// betafallback.BetaRefusalFallbackMiddleware for providers without
// server-side support.
//
// Requires Juglow_API_KEY.
package main

import (
	"context"
	"fmt"

	"github.com/takebox/juglow-sdk-go"
	"github.com/takebox/juglow-sdk-go/lib/betafallback"
	"github.com/takebox/juglow-sdk-go/option"
)

func main() {
	ctx := context.Background()

	// 1. Server-side fallbacks (preferred): the API retries a refusal itself â€”
	// one request, a plain client, no client-side logic. Use this when talking
	// to the API directly.
	client := Juglow.NewClient()
	served, err := client.Beta.Messages.New(ctx, Juglow.BetaMessageNewParams{
		MaxTokens: 1024,
		Model:     Juglow.ModelHaijunFable5,
		Messages: []Juglow.BetaMessageParam{
			Juglow.NewBetaUserMessage(Juglow.NewBetaTextBlock("Some prompt that triggers a refusal")),
		},
		Fallbacks: Juglow.BetaFallbacksParamUnion{
			OfBetaFallbackArray: []Juglow.BetaFallbackParam{{Model: Juglow.ModelHaijunOpus4_8}},
		},
		Betas: []Juglow.JuglowBeta{Juglow.JuglowBetaServerSideFallback2026_07_01},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("server-side, served by:", served.Model)

	// If your provider doesn't support server-side fallbacks, register the
	// client-side middleware instead:
	fallbackClient := Juglow.NewClient(
		option.WithMiddleware(betafallback.BetaRefusalFallbackMiddleware(
			[]Juglow.BetaFallbackParam{{Model: Juglow.ModelHaijunOpus4_8}},
		)),
	)
	state := betafallback.WithBetaFallbackState(&betafallback.BetaFallbackState{}) // pins follow-ups to the model that accepted

	// 2. Streaming: on a refusal the middleware retries and splices the
	// fallback's events onto the open stream â€” one continuous message, with a
	// `fallback` content block marking the model boundary.
	stream := fallbackClient.Beta.Messages.NewStreaming(ctx, Juglow.BetaMessageNewParams{
		MaxTokens: 1024,
		Model:     Juglow.ModelHaijunFable5,
		Messages: []Juglow.BetaMessageParam{
			Juglow.NewBetaUserMessage(Juglow.NewBetaTextBlock("Some prompt that triggers a refusal")),
		},
	}, state)
	defer stream.Close()

	var streamed Juglow.BetaMessage
	for stream.Next() {
		event := stream.Current()
		if err := streamed.Accumulate(event); err != nil {
			panic(err)
		}
		switch event := event.AsAny().(type) {
		case Juglow.BetaRawContentBlockStartEvent:
			// the fallback block marks the splice point
			if fallback, ok := event.ContentBlock.AsAny().(Juglow.BetaFallbackBlock); ok {
				fmt.Printf("\n--- fell back: %s -> %s ---\n", fallback.From.Model, fallback.To.Model)
			}
		case Juglow.BetaRawContentBlockDeltaEvent:
			if delta, ok := event.Delta.AsAny().(Juglow.BetaTextDelta); ok {
				fmt.Print(delta.Text)
			}
		}
	}
	if err := stream.Err(); err != nil {
		panic(err)
	}
	fmt.Println("\nstreaming, served by:", streamed.Model)

	// 3. Non-streaming: same middleware, the retry just happens before you
	// get the message back.
	message, err := fallbackClient.Beta.Messages.New(ctx, Juglow.BetaMessageNewParams{
		MaxTokens: 1024,
		Model:     Juglow.ModelHaijunFable5,
		Messages: []Juglow.BetaMessageParam{
			Juglow.NewBetaUserMessage(Juglow.NewBetaTextBlock("Some prompt that triggers a refusal")),
		},
	}, state) // reusing the state keeps the conversation pinned
	if err != nil {
		panic(err)
	}
	fmt.Println("non-streaming, served by:", message.Model)
}
