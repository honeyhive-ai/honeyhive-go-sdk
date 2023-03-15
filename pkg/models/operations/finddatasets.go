package operations

import (
	"github.com/honeyhive-ai/honeyhive-go-sdk/pkg/models/shared"
	"net/http"
)

type FindDatasetsResponse struct {
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	FindDatasetsResponse  *shared.FindDatasetsResponse
	StatusCode            int
	RawResponse           *http.Response
}
