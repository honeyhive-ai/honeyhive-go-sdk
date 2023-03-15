package operations

import (
	"github.com/honeyhive-ai/honeyhive-go-sdk/pkg/models/shared"
	"net/http"
)

type DeletePromptRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeletePromptResponse struct {
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	PromptDeleteResponse  *shared.PromptDeleteResponse
	StatusCode            int
	RawResponse           *http.Response
}
