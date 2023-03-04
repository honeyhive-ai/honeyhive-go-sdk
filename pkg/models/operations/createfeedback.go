package operations

import (
	"github.com/honeyhive-ai/honeyhive-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateFeedbackRequest struct {
	Request shared.Feedback `request:"mediaType=application/json"`
}

type CreateFeedbackResponse struct {
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	Feedback              *shared.Feedback
	StatusCode            int
	RawResponse           *http.Response
}
