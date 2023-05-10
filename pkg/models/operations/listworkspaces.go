// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte-test/pkg/models/shared"
	"net/http"
)

type ListWorkspacesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful operation
	WorkspaceReadList *shared.WorkspaceReadList
}
