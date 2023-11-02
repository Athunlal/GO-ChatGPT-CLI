package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	apitKey := viper.GetString("API_Key")
	if apitKey == "" {
		log.Fatalln("Missing API key")
	}

	ctx := context.Background()
	clinet := gpt3.NewClient(apitKey)

	const intputFile = "./input_with_code.txt"
	fileBytes, err := os.ReadFile(intputFile)
	if err != nil {
		log.Fatalln(err)
	}
	msgPrefix := "give me a short list of libraries that are used in the code \n```python\n"
	msgSuffix := "\n```"
	msg := msgPrefix + string(fileBytes) + msgSuffix
	outputBuilder := strings.Builder{}
	err = clinet.CompletionStreamWithEngine(ctx, gpt3.TextDavinci001Engine, gpt3.CompletionRequest{
		Prompt: []string{
			msg,
		},
		MaxTokens:   gpt3.IntPtr(3000),
		Temperature: gpt3.Float32Ptr(0),
	}, func(res *gpt3.CompletionResponse) {
		outputBuilder.WriteString(res.Choices[0].Text)
	})

	if err != nil {
		log.Fatalln(err)
	}
	output := strings.TrimSpace(outputBuilder.String())
	const outputFile = "./output.txt"
	err = os.WriteFile(outputFile, []byte(output), os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}

}
