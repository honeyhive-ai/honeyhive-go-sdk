package operations

import (
	"github.com/honeyhive-ai/honeyhive-go-sdk/pkg/models/shared"
)

type ModelPromptCreateGenerationRequest struct {
	Request shared.ModelPromptGenerationInput `request:"mediaType=application/json"`
}

type ModelPromptCreateGenerationResponse struct {
	APIResponseBadRequest       *shared.APIResponseBadRequest
	ContentType                 string
	ModelPromptGenerationOutput *shared.ModelPromptGenerationOutput
	StatusCode                  int
}
