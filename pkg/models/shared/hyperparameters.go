package shared

type HyperParameters struct {
	FrequencyPenalty *float64 `json:"frequency_penalty,omitempty"`
	MaxTokens        *int64   `json:"max_tokens,omitempty"`
	PresencePenalty  *float64 `json:"presence_penalty,omitempty"`
	Stop             []string `json:"stop,omitempty"`
	Temperature      *float64 `json:"temperature,omitempty"`
	TopP             *float64 `json:"top_p,omitempty"`
}
