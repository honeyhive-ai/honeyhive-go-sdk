package operations

import (
	"github.com/honeyhive-ai/honeyhive-go-sdk/pkg/models/shared"
)

type TaskCreateGenerationRequest struct {
	Request shared.TaskGenerationInput `request:"mediaType=application/json"`
}

type TaskCreateGenerationResponse struct {
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	StatusCode            int
	TaskGenerationOutput  *shared.TaskGenerationOutput
}
