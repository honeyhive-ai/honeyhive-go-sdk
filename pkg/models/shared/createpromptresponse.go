package shared

type CreatePromptResponse struct {
	CreatedAt       *string          `json:"created_at,omitempty"`
	FewShotExamples []string         `json:"few_shot_examples,omitempty"`
	Hyperparameters *HyperParameters `json:"hyperparameters,omitempty"`
	InputVariables  []string         `json:"input_variables,omitempty"`
	Model           *string          `json:"model,omitempty"`
	Name            *string          `json:"name,omitempty"`
	Provider        *string          `json:"provider,omitempty"`
	Task            *string          `json:"task,omitempty"`
	Text            *string          `json:"text,omitempty"`
	UpdatedAt       *string          `json:"updated_at,omitempty"`
	Version         *string          `json:"version,omitempty"`
}
