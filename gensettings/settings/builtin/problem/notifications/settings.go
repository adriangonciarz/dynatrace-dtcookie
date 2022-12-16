package notifications

import "github.com/dtcookie/hcl"

type Settings struct {
	AlertingProfile          string                    `json:"alertingProfile"` // Select an [alerting profile](/ui/settings/builtin:alerting.profile) to control the delivery of problem notifications related to this integration.
	AnsibleTowerNotification *AnsibleTowerNotification `json:"ansibleTowerNotification"`
	DisplayName              string                    `json:"displayName"` // The name of the notification configuration.
	EmailNotification        *EmailNotification        `json:"emailNotification"`
	Enabled                  bool                      `json:"enabled"` // Enabled
	JiraNotification         *JiraNotification         `json:"jiraNotification"`
	OpsGenieNotification     *OpsGenieNotification     `json:"opsGenieNotification"`
	PagerDutyNotification    *PagerDutyNotification    `json:"pagerDutyNotification"`
	ServiceNowNotification   *ServiceNowNotification   `json:"serviceNowNotification"`
	SlackNotification        *SlackNotification        `json:"slackNotification"`
	TrelloNotification       *TrelloNotification       `json:"trelloNotification"`
	Type                     NotificationType          `json:"type"` // Notification type
	VictorOpsNotification    *VictorOpsNotification    `json:"victorOpsNotification"`
	WebHookNotification      *WebHookNotification      `json:"webHookNotification"`
	XMattersNotification     *XMattersNotification     `json:"xMattersNotification"`
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"alerting_profile": {
			Type:        hcl.TypeString,
			Description: "Select an [alerting profile](/ui/settings/builtin:alerting.profile) to control the delivery of problem notifications related to this integration.",
			Required:    true,
		},
		"ansible_tower_notification": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(AnsibleTowerNotification).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"display_name": {
			Type:        hcl.TypeString,
			Description: "The name of the notification configuration.",
			Required:    true,
		},
		"email_notification": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(EmailNotification).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enabled",
			Optional:    true,
		},
		"jira_notification": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(JiraNotification).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"ops_genie_notification": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(OpsGenieNotification).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"pager_duty_notification": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(PagerDutyNotification).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"service_now_notification": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ServiceNowNotification).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"slack_notification": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(SlackNotification).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"trello_notification": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(TrelloNotification).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"type": {
			Type:        hcl.TypeString,
			Description: "Notification type",
			Required:    true,
		},
		"victor_ops_notification": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(VictorOpsNotification).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"web_hook_notification": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(WebHookNotification).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"x_matters_notification": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(XMattersNotification).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"alerting_profile":           me.AlertingProfile,
		"ansible_tower_notification": me.AnsibleTowerNotification,
		"display_name":               me.DisplayName,
		"email_notification":         me.EmailNotification,
		"enabled":                    me.Enabled,
		"jira_notification":          me.JiraNotification,
		"ops_genie_notification":     me.OpsGenieNotification,
		"pager_duty_notification":    me.PagerDutyNotification,
		"service_now_notification":   me.ServiceNowNotification,
		"slack_notification":         me.SlackNotification,
		"trello_notification":        me.TrelloNotification,
		"type":                       me.Type,
		"victor_ops_notification":    me.VictorOpsNotification,
		"web_hook_notification":      me.WebHookNotification,
		"x_matters_notification":     me.XMattersNotification,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"alerting_profile":           &me.AlertingProfile,
		"ansible_tower_notification": &me.AnsibleTowerNotification,
		"display_name":               &me.DisplayName,
		"email_notification":         &me.EmailNotification,
		"enabled":                    &me.Enabled,
		"jira_notification":          &me.JiraNotification,
		"ops_genie_notification":     &me.OpsGenieNotification,
		"pager_duty_notification":    &me.PagerDutyNotification,
		"service_now_notification":   &me.ServiceNowNotification,
		"slack_notification":         &me.SlackNotification,
		"trello_notification":        &me.TrelloNotification,
		"type":                       &me.Type,
		"victor_ops_notification":    &me.VictorOpsNotification,
		"web_hook_notification":      &me.WebHookNotification,
		"x_matters_notification":     &me.XMattersNotification,
	})
}
