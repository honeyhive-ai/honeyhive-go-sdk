package shared

type SingleGenerationInput struct {
    Generation      *string                `json:"generation,omitempty"`
	Hyperparameters *HyperParameters       `json:"hyperparameters,omitempty"`
	Inputs          map[string]interface{} `json:"inputs,omitempty"`
	Latency         *float64               `json:"latency,omitempty"`
	Model           *string                `json:"model,omitempty"`
	Prompt          *string                `json:"prompt,omitempty"`
	Response        map[string]interface{} `json:"response,omitempty"`
	Source          *string                `json:"source,omitempty"`
	Task            *string                `json:"task,omitempty"`
}
