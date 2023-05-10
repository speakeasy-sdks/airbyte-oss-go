// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte-test/pkg/models/shared"
	"net/http"
)

type CreateOperationResponse struct {
	ContentType string
	// Input failed validation
	InvalidInputExceptionInfo *shared.InvalidInputExceptionInfo
	// Successful operation
	OperationRead *shared.OperationRead
	StatusCode    int
	RawResponse   *http.Response
}
