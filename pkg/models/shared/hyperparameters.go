package shared

type HyperParameters struct {
	FrequencyPenalty *float32 `json:"frequency_penalty,omitempty"`
	MaxTokens        *int   `json:"max_tokens,omitempty"`
	PresencePenalty  *float32 `json:"presence_penalty,omitempty"`
	Stop             []string `json:"stop,omitempty"`
	Temperature      *float32 `json:"temperature,omitempty"`
	TopP             *float32 `json:"top_p,omitempty"`
}
