// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AttemptFailureSummary struct {
	Failures []AttemptFailureReason `json:"failures"`
	// True if the number of committed records for this attempt was greater than 0. False if 0 records were committed. If not set, the number of committed records is unknown.
	PartialSuccess *bool `json:"partialSuccess,omitempty"`
}
