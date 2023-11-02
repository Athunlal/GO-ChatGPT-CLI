package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func GetResponse(client gpt3.Client, ctx context.Context, question string) {
	err := client.CompletionStreamWithEngine(ctx, gpt3.TextDavinci001Engine, gpt3.CompletionRequest{
		Prompt: []string{
			question,
		},
		MaxTokens:   gpt3.IntPtr(3000),
		Temperature: gpt3.Float32Ptr(0),
	}, func(res *gpt3.CompletionResponse) {
		fmt.Println(res.Choices[0].Text)
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(13)
	}

}

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	apiKey := viper.GetString("API_Key")
	if apiKey == "" {
		log.Fatalln("Missing api key")
	}
	ctx := context.Background()
	client := gpt3.NewClient(apiKey)
	rootCmd := &cobra.Command{
		Use:   "chatgpt",
		Short: "chat with ChatGPT in consol",
		Run: func(cmd *cobra.Command, args []string) {
			scanner := bufio.NewScanner(os.Stdin)
			quit := false

			for !quit {
				fmt.Println("Say somthing ('quit' to end):")
				if !scanner.Scan() {
					break
				}
				question := scanner.Text()
				switch question {
				case "quit":
					quit = true
				default:
					GetResponse(client, ctx, question)
				}
			}
		},
	}

	rootCmd.Execute()
}
