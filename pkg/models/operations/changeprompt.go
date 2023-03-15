package operations

import (
	"github.com/honeyhive-ai/honeyhive-go-sdk/pkg/models/shared"
	"net/http"
)

type ChangePromptRequest struct {
	Prompt shared.Prompt `request:"mediaType=application/json"`
	ID     string        `pathParam:"style=simple,explode=false,name=id"`
}

type ChangePromptResponse struct {
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	Prompt                *shared.Prompt
	StatusCode            int
	RawResponse           *http.Response
}
