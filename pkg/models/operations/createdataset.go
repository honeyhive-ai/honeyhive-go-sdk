package operations

import (
	"github.com/honeyhive-ai/honeyhive-go-sdk/pkg/models/shared"
)

type CreateDatasetRequest struct {
	Request shared.Dataset `request:"mediaType=application/json"`
}

type CreateDatasetResponse struct {
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	CreateDatasetResponse *shared.CreateDatasetResponse
	StatusCode            int
}
