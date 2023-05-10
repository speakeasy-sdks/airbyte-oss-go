// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionScheduleTimeUnitEnum string

const (
	ConnectionScheduleTimeUnitEnumMinutes ConnectionScheduleTimeUnitEnum = "minutes"
	ConnectionScheduleTimeUnitEnumHours   ConnectionScheduleTimeUnitEnum = "hours"
	ConnectionScheduleTimeUnitEnumDays    ConnectionScheduleTimeUnitEnum = "days"
	ConnectionScheduleTimeUnitEnumWeeks   ConnectionScheduleTimeUnitEnum = "weeks"
	ConnectionScheduleTimeUnitEnumMonths  ConnectionScheduleTimeUnitEnum = "months"
)

func (e ConnectionScheduleTimeUnitEnum) ToPointer() *ConnectionScheduleTimeUnitEnum {
	return &e
}

func (e *ConnectionScheduleTimeUnitEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "minutes":
		fallthrough
	case "hours":
		fallthrough
	case "days":
		fallthrough
	case "weeks":
		fallthrough
	case "months":
		*e = ConnectionScheduleTimeUnitEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionScheduleTimeUnitEnum: %v", v)
	}
}

// ConnectionSchedule - if null, then no schedule is set.
type ConnectionSchedule struct {
	TimeUnit ConnectionScheduleTimeUnitEnum `json:"timeUnit"`
	Units    int64                          `json:"units"`
}
