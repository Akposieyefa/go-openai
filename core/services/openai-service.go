package services

import (
	"context"

	"github.com/akposieyefa/open-ai/config"
	openai "github.com/sashabaranov/go-openai"
)

var client = openai.NewClient(config.Envs.API_KEY)

func ChatCompletion(content string) (*openai.ChatCompletionMessage, error) {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: content,
				},
			},
		},
	)

	if err != nil {
		return nil, err
	}

	return &resp.Choices[0].Message, nil
}
