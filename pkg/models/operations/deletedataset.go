// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/honeyhive-ai/honeyhive-go-sdk/pkg/models/shared"
	"net/http"
)

type DeleteDatasetRequest struct {
	// The name of the dataset
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

type DeleteDatasetResponse struct {
	// Bad Request error
	APIResponseBadRequest *shared.APIResponseBadRequest
	ContentType           string
	// Successful response
	DeleteDatasetResponse *shared.DeleteDatasetResponse
	StatusCode            int
	RawResponse           *http.Response
}
