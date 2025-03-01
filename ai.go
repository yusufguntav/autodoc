package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func sendMessageAI(changes string) string {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text("Please create general documentation and documentation for frontend developers with the following information: key commit names and value changes: "+changes))
	if err != nil {
		log.Fatal(err)
	}

	return returnResponse(resp)

}

func returnResponse(resp *genai.GenerateContentResponse) string {
	response := ""
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				response += fmt.Sprintf("%s\n", part)
			}
		}
	}
	return response
}
