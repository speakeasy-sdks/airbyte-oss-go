// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Notification struct {
	CustomerioConfiguration map[string]interface{}          `json:"customerioConfiguration,omitempty"`
	NotificationType        NotificationTypeEnum            `json:"notificationType"`
	SendOnFailure           bool                            `json:"sendOnFailure"`
	SendOnSuccess           bool                            `json:"sendOnSuccess"`
	SlackConfiguration      *SlackNotificationConfiguration `json:"slackConfiguration,omitempty"`
}
