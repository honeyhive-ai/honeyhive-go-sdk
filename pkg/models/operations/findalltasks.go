package operations

import (
	"github.com/honeyhive-ai/honeyhive-go-sdk/pkg/models/shared"
	"net/http"
)

type FindAllTasksResponse struct {
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	StatusCode            int
	RawResponse           *http.Response
	Tasks                 *shared.Tasks
}
