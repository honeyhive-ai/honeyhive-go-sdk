package shared

type Feedback struct {
	FeedbackJSON map[string]interface{} `json:"feedback_json,omitempty"`
	GenerationID *string               `json:"generation_id,omitempty"`
	Task         *string               `json:"task,omitempty"`
	Timestamp    *string               `json:"timestamp,omitempty"`
}
