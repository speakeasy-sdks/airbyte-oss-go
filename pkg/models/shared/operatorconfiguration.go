// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type OperatorConfiguration struct {
	Dbt           *OperatorDbt           `json:"dbt,omitempty"`
	Normalization *OperatorNormalization `json:"normalization,omitempty"`
	OperatorType  OperatorTypeEnum       `json:"operatorType"`
	Webhook       *OperatorWebhook       `json:"webhook,omitempty"`
}
