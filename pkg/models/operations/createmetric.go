package operations

import (
	"github.com/honeyhive-ai/honeyhive-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateMetricResponse struct {
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	Metric                *shared.Metric
	StatusCode            int
	RawResponse           *http.Response
}
