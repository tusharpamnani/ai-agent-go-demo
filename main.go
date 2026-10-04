package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"

	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
)

func main() {
	apiKey := os.Getenv("OPENROUTER_API_KEY")

	if apiKey == "" {
		log.Fatal("OPENROUTER_API_KEY environment variable is not set")
	}

	url := "https://openrouter.ai/api/v1"

	client := openai.NewClient(
		option.WithBaseURL(url),
		option.WithAPIKey(apiKey),
	)

	ctx := context.Background()

	messages := []openai.ChatCompletionMessageParamUnion{}

	model := "google/gemma-4-26b-a4b-it:free"

	params := openai.ChatCompletionNewParams{
		Model:    model,
		Messages: messages,
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(color.WhiteString(
		fmt.Sprintf("model: %s", model),
	))

	for {
		fmt.Print(color.CyanString("\n> "))

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			continue
		}

		if input == "clear" {
			fmt.Print("\033[H\033[2J")
			continue
		}

		if input == "exit" {
			fmt.Println("Exiting...")
			break
		}

		params.Messages = append(params.Messages, openai.UserMessage(input))

		res, err := client.Chat.Completions.New(ctx, params)
		if err != nil {
			log.Fatal(err)
		}

		output := res.Choices[0].Message.Content

		fmt.Println(color.YellowString(output))

		params.Messages = append(params.Messages, openai.AssistantMessage(output))

		count := len(params.Messages)

		fmt.Print(color.MagentaString(fmt.Sprintf("count: %d", count)))

		if count >= 10 {
			params.Messages = params.Messages[count-4:]
			fmt.Println(color.YellowString("history refreshed!"))
		}

	}

}
