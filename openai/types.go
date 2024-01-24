package openai

import "github.com/lomsa-dev/gonull"

type ChatRequestImage struct {
	Messages  []ImageMessage       `json:"messages"`
	Model     string               `json:"model"`
	MaxTokens gonull.Nullable[int] `json:"max_tokens"`
}

type ImageMessage struct {
	Content any                     `json:"content"`
	Role    string                  `json:"role"`
	Name    gonull.Nullable[string] `json:"name"`
}

type TextContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type ImageContent struct {
	Type     string   `json:"type"`
	ImageURL ImageURL `json:"image_url"`
}

type ImageURL struct {
	URL string `json:"url"`
}

type ChatRequestDefault struct {
	Messages  []DefaultMessage `json:"messages"`
	Model     string           `json:"model"`
	MaxTokens gonull.Nullable[int]
}

type DefaultMessage struct {
	Content string                  `json:"content"`
	Role    string                  `json:"role"`
	Name    gonull.Nullable[string] `json:"name"`
}

type ChatCompletion struct {
	ID                string   `json:"id"`
	Choices           []Choice `json:"choices"`
	Created           int      `json:"created"`
	Model             string   `json:"model"`
	SystemFingerprint string   `json:"system_fingerprint"`
	Object            string   `json:"object"`
	Usage             Usage    `json:"usage"`
}

type Choice struct {
	FinishReason string                    `json:"finish_reason"`
	Index        int                       `json:"index"`
	Message      ChoiceMessage             `json:"message"`
	Logprobs     gonull.Nullable[LogProbs] `json:"logprobs"`
}

type ChoiceMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type LogProbs struct {
	Content gonull.Nullable[LogProbsContent] `json:"content"`
}

type LogProbsContent struct {
	Token        string                 `json:"token"`
	LogProbValue int                    `json:"logprob"`
	Bytes        gonull.Nullable[[]int] `json:"bytes"`
}

type Usage struct {
	CompletionTokens int `json:"completion_tokens"`
	PromptTokens     int `json:"prompt_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
