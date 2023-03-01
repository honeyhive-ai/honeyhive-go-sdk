package operations

import (
	"github.com/honeyhive-ai/honeyhive-go-sdk/pkg/models/shared"
)

type IngestSingleGenerationRequest struct {
	Request shared.SingleGenerationInput `request:"mediaType=application/json"`
}

type IngestSingleGenerationResponse struct {
	APIResponseBadRequest  *shared.APIResponseBadRequest
	ContentType            string
	SingleGenerationOutput *shared.SingleGenerationOutput
	StatusCode             int
}
