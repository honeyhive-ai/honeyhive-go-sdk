package operations

import (
	"github.com/honeyhive-ai/honeyhive-go-sdk/pkg/models/shared"
)

type FindDatasetsRequest struct {
	Request *shared.FindDatasetsInput `request:"mediaType=application/json"`
}

type FindDatasetsResponse struct {
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	FindDatasetsResponse  *shared.FindDatasetsResponse
	StatusCode            int
}
