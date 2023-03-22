// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/honeyhive-ai/honeyhive-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateMetricResponse struct {
	// Bad Request error
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	// Successful response
	Metric      *shared.Metric
	StatusCode  int
	RawResponse *http.Response
}
