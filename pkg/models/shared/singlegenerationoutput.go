package shared

type SingleGenerationOutput struct {
	CompletionTokens *int64   `json:"completion_tokens,omitempty"`
	Cost             *float64 `json:"cost,omitempty"`
	Generation       *string  `json:"generation,omitempty"`
	GenerationID     *string  `json:"generation_id,omitempty"`
	Latency          *float64 `json:"latency,omitempty"`
	TotalTokens      *int64   `json:"total_tokens,omitempty"`
	Version          *string  `json:"version,omitempty"`
}
