package notifications

import (
	"errors"
	"fmt"

	"github.com/dtcookie/hcl"
)

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

func (me *Notification) MarshalHCL() (map[string]interface{}, error) {
	m := map[Type]hcl.Marshaler{
		Types.AnsibleTower: me.AnsibleTower,
		Types.Email:        me.Email,
		Types.Slack:        me.Slack,
		Types.Jira:         me.Jira,
		Types.OpsGenie:     me.OpsGenie,
		Types.PagerDuty:    me.PagerDuty,
		Types.VictorOps:    me.VictorOps,
		Types.WebHook:      me.WebHook,
		Types.XMatters:     me.XMatters,
		Types.Trello:       me.Trello,
		Types.ServiceNow:   me.ServiceNow,
	}
	if len(me.Type) == 0 {
		return nil, errors.New("no notification type set")
	}
	if config, ok := m[me.Type]; ok {
		if config == nil {
			return nil, fmt.Errorf("notification type is `%v` but the corresponding configuration is missing", me.Type)
		}
		result, err := config.MarshalHCL()
		if err != nil {
			return nil, err
		}
		if me.Enabled {
			result["enabled"] = me.Enabled
		}
		result["name"] = me.Name
		result["profile"] = me.ProfileID
		result["enabled"] = me.Enabled
		return result, nil
	}
	return nil, fmt.Errorf("notification type `%v` not supported", me.Type)
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
