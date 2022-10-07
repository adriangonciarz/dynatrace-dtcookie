package notifications

type Notification struct {
	Enabled   bool   `json:"enabled"`
	Name      string `json:"displayName"`
	ProfileID string `json:"alertingProfile"`
	Type      Type   `json:"type"`

	Email        *Email        `json:"emailNotification,omitempty"`
	Slack        *Slack        `json:"slackNotification,omitempty"`
	Jira         *Jira         `json:"jiraNotification,omitempty"`
	AnsibleTower *AnsibleTower `json:"ansibleTowerNotification,omitempty"`
	OpsGenie     *OpsGenie     `json:"opsGenieNotification,omitempty"`
	PagerDuty    *PagerDuty    `json:"pagerDutyNotification,omitempty"`
	VictorOps    *VictorOps    `json:"victorOpsNotification,omitempty"`
	WebHook      *WebHook      `json:"webHookNotification,omitempty"`
	XMatters     *XMatters     `json:"xMattersNotification,omitempty"`
	Trello       *Trello       `json:"trelloNotification,omitempty"`
	ServiceNow   *ServiceNow   `json:"serviceNowNotification,omitempty"`
}

type Type string

var Types = struct {
	Email        Type
	Slack        Type
	Jira         Type
	AnsibleTower Type
	OpsGenie     Type
	PagerDuty    Type
	VictorOps    Type
	WebHook      Type
	XMatters     Type
	Trello       Type
	ServiceNow   Type
}{
	Type("EMAIL"),
	Type("SLACK"),
	Type("JIRA"),
	Type("ANSIBLETOWER"),
	Type("OPS_GENIE"),
	Type("PAGER_DUTY"),
	Type("VICTOROPS"),
	Type("WEBHOOK"),
	Type("XMATTERS"),
	Type("TRELLO"),
	Type("SERVICE_NOW"),
}
