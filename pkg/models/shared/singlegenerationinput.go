package shared

type SingleGenerationInput struct {
	Generation      *string                `json:"generation,omitempty"`
	Hyperparameters *HyperParameters       `json:"hyperparameters,omitempty"`
	Inputs          map[string]interface{} `json:"inputs,omitempty"`
	Latency         *float64               `json:"latency,omitempty"`
	Model           *string                `json:"model,omitempty"`
	PromptTemplate          *string                `json:"prompt_template,omitempty"`
	Usage        map[string]int `json:"usage,omitempty"`
	Source          *string                `json:"source,omitempty"`
	Task            *string                `json:"task,omitempty"`
}
