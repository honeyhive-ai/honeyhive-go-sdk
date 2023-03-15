package operations

import (
	"github.com/honeyhive-ai/honeyhive-go-sdk/pkg/models/shared"
	"net/http"
)

type ModelPromptCreateGenerationResponse struct {
	APIResponseBadRequest       *shared.APIResponseBadRequest
	ContentType                 string
	ModelPromptGenerationOutput *shared.ModelPromptGenerationOutput
	StatusCode                  int
	RawResponse                 *http.Response
}
