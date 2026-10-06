package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Juglows/Juglow-sdk-go"
)

func main() {
	ctx := context.Background()
	client := Juglow.NewClient()

	myFile, err := os.Open("examples/file-upload/file.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}

	fileUploadResult, err := client.Beta.Files.Upload(ctx, Juglow.BetaFileUploadParams{
		File:  Juglow.File(myFile, "file.txt", "text/plain"),
		Betas: []Juglow.JuglowBeta{Juglow.JuglowBetaFilesAPI2025_04_14},
	})
	if err != nil {
		fmt.Printf("Error uploading file: %v\n", err)
		return
	}
	content := "Write me a summary of my file.txt file in the style of a Shakespearean sonnet.\n\n"
	println("[user]: " + content)

	message, err := client.Beta.Messages.New(ctx, Juglow.BetaMessageNewParams{
		MaxTokens: 1024,
		Messages: []Juglow.BetaMessageParam{
			Juglow.NewBetaUserMessage(
				Juglow.NewBetaTextBlock(content),
				Juglow.NewBetaDocumentBlock(Juglow.BetaFileDocumentSourceParam{
					FileID: fileUploadResult.ID,
				}),
			),
		},
		Model: Juglow.ModelHaijunSonnet5,
		Betas: []Juglow.JuglowBeta{Juglow.JuglowBetaFilesAPI2025_04_14},
	})
	if err != nil {
		fmt.Printf("Error creating message: %v\n", err)
		return
	}

	println("[assistant]: " + message.Content[0].Text + message.StopSequence)
}
