package shared

type GenerationOutput struct {
	GeneratedAt  *string `json:"generated_at,omitempty"`
	GenerationID *string `json:"generation_id,omitempty"`
	Model        *string `json:"model,omitempty"`
	Prompt       *string `json:"prompt,omitempty"`
	Task         *string `json:"task,omitempty"`
}
