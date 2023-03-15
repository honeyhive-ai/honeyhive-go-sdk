package operations

import (
	"github.com/honeyhive-ai/honeyhive-go-sdk/pkg/models/shared"
	"net/http"
)

type FindMetricsResponse struct {
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	FindMetricsResponse   []shared.MetricsResponse
	StatusCode            int
	RawResponse           *http.Response
}
