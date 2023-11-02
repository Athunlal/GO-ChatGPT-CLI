package main

import (
	"context"
	"fmt"
	"log"
	"os"

	gpt3 "github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatalln("Missing Api key")
	}
	ctx := context.Background()
	client := gpt3.NewClient(apiKey)

	res, err := client.Completion(ctx, gpt3.CompletionRequest{
		Prompt:    []string{"The first thing you should know about golang is"},
		MaxTokens: gpt3.IntPtr(30),
		Echo:      true,
		Stop:      []string{"."},
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res.Choices[0].Text)
}
