package operations

import (
	"github.com/honeyhive-ai/honeyhive-go-sdk/pkg/models/shared"
)

type FindPromptsRequest struct {
	Request *shared.FindPromptsInput `request:"mediaType=application/json"`
}

type FindPromptsResponse struct {
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	Prompts               *shared.Prompts
	StatusCode            int
}
