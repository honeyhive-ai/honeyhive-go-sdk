package operations

import (
	"github.com/honeyhive-ai/honeyhive-go-sdk/pkg/models/shared"
	"net/http"
)

type IngestSingleGenerationResponse struct {
	APIResponseBadRequest  *shared.APIResponseBadRequest
	ContentType            string
	SingleGenerationOutput *shared.SingleGenerationOutput
	StatusCode             int
	RawResponse            *http.Response
}
