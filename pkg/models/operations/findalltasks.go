package operations

import (
	"github.com/honeyhive-ai/honeyhive-go-sdk/pkg/models/shared"
	"net/http"
)

type FindAllTasksRequest struct {
	Request *string `request:"mediaType=application/json"`
}

type FindAllTasksResponse struct {
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	StatusCode            int
	RawResponse           *http.Response
	Tasks                 *shared.Tasks
}
