// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte-test/pkg/models/shared"
	"net/http"
)

type WebBackendUpdateConnectionResponse struct {
	ContentType string
	// Input failed validation
	InvalidInputExceptionInfo *shared.InvalidInputExceptionInfo
	StatusCode                int
	RawResponse               *http.Response
	// Successful operation
	WebBackendConnectionRead *shared.WebBackendConnectionRead
}
