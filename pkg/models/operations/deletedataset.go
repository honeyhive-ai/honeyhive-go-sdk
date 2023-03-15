package operations

import (
	"github.com/honeyhive-ai/honeyhive-go-sdk/pkg/models/shared"
	"net/http"
)

type DeleteDatasetRequest struct {
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

type DeleteDatasetResponse struct {
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	DeleteDatasetResponse *shared.DeleteDatasetResponse
	StatusCode            int
	RawResponse           *http.Response
}
