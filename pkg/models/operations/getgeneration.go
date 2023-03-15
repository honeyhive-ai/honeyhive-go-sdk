package operations

import (
	"github.com/honeyhive-ai/honeyhive-go-sdk/pkg/models/shared"
	"net/http"
)

type GetGenerationResponse struct {
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	GenerationOutputs     *shared.GenerationOutputs
	StatusCode            int
	RawResponse           *http.Response
}
