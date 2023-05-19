// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ConnectionState - Contains the state for a connection. The stateType field identifies what type of state it is. Only the field corresponding to that type will be set, the rest will be null. If stateType=not_set, then none of the fields will be set.
type ConnectionState struct {
	ConnectionID string                 `json:"connectionId"`
	GlobalState  *GlobalState           `json:"globalState,omitempty"`
	State        map[string]interface{} `json:"state,omitempty"`
	StateType    ConnectionStateType    `json:"stateType"`
	StreamState  []StreamState          `json:"streamState,omitempty"`
}
