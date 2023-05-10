// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceCreate struct {
	// The values required to configure the source. The schema for this must match the schema return by source_definition_specifications/get for the source.
	ConnectionConfiguration interface{} `json:"connectionConfiguration"`
	Name                    string      `json:"name"`
	SourceDefinitionID      string      `json:"sourceDefinitionId"`
	WorkspaceID             string      `json:"workspaceId"`
}
