package operations

import (
	"github.com/honeyhive-ai/honeyhive-go-sdk/pkg/models/shared"
	"net/http"
)

type FindFineTunedModelsResponse struct {
	APIResponseBadRequest   *shared.APIResponseBadRequest
	ContentType             string
	FineTunedModelsResponse *shared.FineTunedModelsResponse
	StatusCode              int
	RawResponse             *http.Response
}
