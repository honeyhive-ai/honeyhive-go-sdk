package operations

import (
	"github.com/honeyhive-ai/honeyhive-go-sdk/pkg/models/shared"
)

type FindMetricsRequest struct {
	Request *shared.FindMetricsInput `request:"mediaType=application/json"`
}

type FindMetricsResponse struct {
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	FindMetricsResponse   []shared.MetricsResponse
	StatusCode            int
}
