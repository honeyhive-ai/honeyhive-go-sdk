package shared

type TaskTypeEnum string

const (
	TaskTypeEnumTextGeneration TaskTypeEnum = "Text Generation"
	TaskTypeEnumCodeGeneration TaskTypeEnum = "Code Generation"
	TaskTypeEnumClassification TaskTypeEnum = "Classification"
	TaskTypeEnumSummarization  TaskTypeEnum = "Summarization"
	TaskTypeEnumOther          TaskTypeEnum = "Other"
)

type Task struct {
	CreatedAt       *string       `json:"created_at,omitempty"`
	Datasets        []Dataset     `json:"datasets,omitempty"`
	Description     *string       `json:"description,omitempty"`
	FineTunedModels []string      `json:"fine_tuned_models,omitempty"`
	ID              *string       `json:"id,omitempty"`
	Metrics         []Metric      `json:"metrics,omitempty"`
	Name            *string       `json:"name,omitempty"`
	Prompts         []Prompt      `json:"prompts,omitempty"`
	Type            *TaskTypeEnum `json:"type,omitempty"`
	UpdatedAt       *string       `json:"updated_at,omitempty"`
}
