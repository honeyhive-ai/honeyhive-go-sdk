package operations

import (
	"github.com/honeyhive-ai/honeyhive-go-sdk/pkg/models/shared"
	"net/http"
)

type IngestGenerationsRequest struct {
	Request shared.Generations `request:"mediaType=application/json"`
}

type IngestGenerationsResponse struct {
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	Generations           *shared.Generations
	StatusCode            int
	RawResponse           *http.Response
}
