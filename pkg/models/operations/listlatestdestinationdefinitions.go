// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte-test/pkg/models/shared"
	"net/http"
)

type ListLatestDestinationDefinitionsResponse struct {
	ContentType string
	// Successful operation
	DestinationDefinitionReadList *shared.DestinationDefinitionReadList
	StatusCode                    int
	RawResponse                   *http.Response
}
