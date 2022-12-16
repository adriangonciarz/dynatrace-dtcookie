package notifications

type NotificationType string

var NotificationTypes = struct {
	Email        NotificationType
	Slack        NotificationType
	Ansibletower NotificationType
	OpsGenie     NotificationType
	Victorops    NotificationType
	Xmatters     NotificationType
	Trello       NotificationType
	PagerDuty    NotificationType
	Webhook      NotificationType
	Jira         NotificationType
	ServiceNow   NotificationType
}{
	NotificationType("EMAIL"),
	NotificationType("SLACK"),
	NotificationType("ANSIBLETOWER"),
	NotificationType("OPS_GENIE"),
	NotificationType("VICTOROPS"),
	NotificationType("XMATTERS"),
	NotificationType("TRELLO"),
	NotificationType("PAGER_DUTY"),
	NotificationType("WEBHOOK"),
	NotificationType("JIRA"),
	NotificationType("SERVICE_NOW"),
}
