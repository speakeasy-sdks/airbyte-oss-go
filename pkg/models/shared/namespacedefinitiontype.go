// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// NamespaceDefinitionType - Method used for computing final namespace in destination
type NamespaceDefinitionType string

const (
	NamespaceDefinitionTypeSource       NamespaceDefinitionType = "source"
	NamespaceDefinitionTypeDestination  NamespaceDefinitionType = "destination"
	NamespaceDefinitionTypeCustomformat NamespaceDefinitionType = "customformat"
)

func (e NamespaceDefinitionType) ToPointer() *NamespaceDefinitionType {
	return &e
}

func (e *NamespaceDefinitionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "source":
		fallthrough
	case "destination":
		fallthrough
	case "customformat":
		*e = NamespaceDefinitionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NamespaceDefinitionType: %v", v)
	}
}
