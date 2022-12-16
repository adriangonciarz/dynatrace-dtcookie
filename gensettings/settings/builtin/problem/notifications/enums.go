package notifications

type NotificationType string

var NotificationTypes = struct {
	Ansibletower NotificationType
	Email        NotificationType
	Jira         NotificationType
	OpsGenie     NotificationType
	PagerDuty    NotificationType
	ServiceNow   NotificationType
	Slack        NotificationType
	Trello       NotificationType
	Victorops    NotificationType
	Webhook      NotificationType
	Xmatters     NotificationType
}{
	NotificationType("ANSIBLETOWER"),
	NotificationType("EMAIL"),
	NotificationType("JIRA"),
	NotificationType("OPS_GENIE"),
	NotificationType("PAGER_DUTY"),
	NotificationType("SERVICE_NOW"),
	NotificationType("SLACK"),
	NotificationType("TRELLO"),
	NotificationType("VICTOROPS"),
	NotificationType("WEBHOOK"),
	NotificationType("XMATTERS"),
}
