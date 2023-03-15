package operations

import (
	"github.com/honeyhive-ai/honeyhive-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateDatasetResponse struct {
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	CreateDatasetResponse *shared.CreateDatasetResponse
	StatusCode            int
	RawResponse           *http.Response
}
