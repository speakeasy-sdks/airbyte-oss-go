// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte-test/pkg/models/shared"
	"net/http"
)

type SetInstancewideSourceOauthParamsResponse struct {
	ContentType string
	// Exception occurred; see message for details.
	KnownExceptionInfo *shared.KnownExceptionInfo
	// Object with given id was not found.
	NotFoundKnownExceptionInfo *shared.NotFoundKnownExceptionInfo
	StatusCode                 int
	RawResponse                *http.Response
}
