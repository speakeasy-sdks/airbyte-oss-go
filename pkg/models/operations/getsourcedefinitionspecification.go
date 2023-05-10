// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte-test/pkg/models/shared"
	"net/http"
)

type GetSourceDefinitionSpecificationResponse struct {
	ContentType string
	// Input failed validation
	InvalidInputExceptionInfo *shared.InvalidInputExceptionInfo
	// Object with given id was not found.
	NotFoundKnownExceptionInfo *shared.NotFoundKnownExceptionInfo
	// Successful operation
	SourceDefinitionSpecificationRead *shared.SourceDefinitionSpecificationRead
	StatusCode                        int
	RawResponse                       *http.Response
}
