// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type AttemptStatusEnum string

const (
	AttemptStatusEnumRunning   AttemptStatusEnum = "running"
	AttemptStatusEnumFailed    AttemptStatusEnum = "failed"
	AttemptStatusEnumSucceeded AttemptStatusEnum = "succeeded"
)

func (e AttemptStatusEnum) ToPointer() *AttemptStatusEnum {
	return &e
}

func (e *AttemptStatusEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "running":
		fallthrough
	case "failed":
		fallthrough
	case "succeeded":
		*e = AttemptStatusEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AttemptStatusEnum: %v", v)
	}
}
