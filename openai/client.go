package openai

import (
	"context"
	"fmt"

	"github.com/carlmjohnson/requests"
	"github.com/lomsa-dev/gonull"
)

type Client struct {
	apiKey string
}

func NewClient(apiKey string) Client {
	return Client{apiKey: apiKey}
}

func (c Client) LookAtImage(ctx context.Context, imgURL, prompt string) (string, error) {

	var response ChatCompletion

	request := prepareImageRequest(prompt, imgURL)

	err := requests.URL("/v1/chat/completions").
		Transport(requests.Record(nil, "response")).
		Host("api.openai.com").
		Header("Authorization", fmt.Sprintf("Bearer %s", c.apiKey)).
		BodyJSON(&request).
		ToJSON(&response).
		Fetch(context.Background())

	if err != nil {
		return "", fmt.Errorf("could not get image look completion: %w", err)
	}

	res := response.Choices[len(response.Choices)-1].Message.Content

	return res, nil

}

func prepareImageRequest(textPrompt, imageUrl string) ChatRequestImage {
	return ChatRequestImage{
		Messages: []ImageMessage{
			{
				Role: "user",
				Content: []any{
					TextContent{
						Type: "text",
						Text: textPrompt,
					},
					ImageContent{
						Type: "image_url",
						ImageURL: ImageURL{
							URL: imageUrl,
						},
					},
				},
			},
		},
		Model:     "gpt-4-vision-preview",
		MaxTokens: gonull.NewNullable[int](300),
	}
}
