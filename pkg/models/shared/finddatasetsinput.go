// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// FindDatasetsInput - Find datasets by task name.
type FindDatasetsInput struct {
	// id of the dataset
	DatasetID *string `json:"dataset_id,omitempty"`
	// prompt of the dataset
	Prompt  *string `json:"prompt,omitempty"`
	Purpose *string `json:"purpose,omitempty"`
	// task of the dataset
	Task *string `json:"task,omitempty"`
}
