package shared

type Generation struct {
	Completion *string `json:"completion,omitempty"`
	CreatedAt  *string `json:"created_at,omitempty"`
	ID         *string `json:"id,omitempty"`
	Model      *string `json:"model,omitempty"`
	Prompt     *string `json:"prompt,omitempty"`
	TaskID     *string `json:"task_id,omitempty"`
	UpdatedAt  *string `json:"updated_at,omitempty"`
}
