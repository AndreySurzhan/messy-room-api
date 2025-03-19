package openai

import (
	"context"
	"fmt"
	"github.com/AndreySurzhan/messy-room-api/internal/config"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/shared"
)

type OpenAI struct {
	client *openai.Client
}

func New(cfg *config.Config) *OpenAI {
	client := openai.NewClient(option.WithAPIKey(cfg.GetString(config.OpenAIAPIKey)))

	return &OpenAI{
		client: client,
	}
}

func (o *OpenAI) GetRoomCleanlinessStatus(ctx context.Context, image string) (string, error) {
	params := openai.ChatCompletionNewParams{
		Model: openai.F(shared.ChatModelGPT4oMini),
		Store: openai.Bool(true),
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.ChatCompletionMessage{
				Role: openai.ChatCompletionMessageRole(openai.ChatCompletionMessageParamRoleUser),
				Content: fmt.Sprintf(`
					[
						{
							"type": "text",
							"text": "It is a kid's bedroom. 
							  I want to determine if it is messy or not.
							  For that I need to have a list of things that should be put away.
							  Please return a JSON object with the following properties:
							  - isMessy: boolean
							  - items: string[]"
						},
						{
							"type": "image_url",
							"image_url": "data:image/jpeg;base64,%s"
						}
					]`, image),
			},
		}),
	}

	res, err := o.client.Chat.Completions.New(ctx, params)
	if err != nil {
		return "", err
	}

	return res.Choices[0].Message.Content, nil
}
