package shared

type ModelPromptGenerationOutput struct {
	CompletionTokens *int64           `json:"completion_tokens,omitempty"`
	Cost             *float64         `json:"cost,omitempty"`
	Generation       *string          `json:"generation,omitempty"`
	Hyperparameters  *HyperParameters `json:"hyperparameters,omitempty"`
	Model            *string          `json:"model,omitempty"`
	Prompt           *string          `json:"prompt,omitempty"`
	Source           *string          `json:"source,omitempty"`
	Task             *string          `json:"task,omitempty"`
	Timestamp        *string          `json:"timestamp,omitempty"`
	TotalTokens      *int64           `json:"total_tokens,omitempty"`
	Version          *string          `json:"version,omitempty"`
}
