package operations

import (
	"github.com/honeyhive-ai/honeyhive-go-sdk/pkg/models/shared"
	"net/http"
)

type CreatePromptResponse struct {
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	CreatePromptResponse  *shared.CreatePromptResponse
	StatusCode            int
	RawResponse           *http.Response
}
