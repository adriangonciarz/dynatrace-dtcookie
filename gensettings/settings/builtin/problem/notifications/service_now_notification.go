package notifications

import "github.com/dtcookie/hcl"

type ServiceNowNotification struct {
	InstanceName  string `json:"instanceName"`  // The ServiceNow instance identifier. It refers to the first part of your own ServiceNow URL. \n\n This field is mutually exclusive with the **url** field. You can only use one of them.
	Message       string `json:"message"`       // The content of the ServiceNow description. Type '{' for placeholder suggestions.. #### Available placeholders\n**{ImpactedEntity}**: A short description of the problem and impacted entity (or multiple impacted entities).\n\n**{ImpactedEntityNames}**: The entity impacted by the problem (or multiple impacted entities).\n\n**{PID}**: Unique system identifier of the reported problem.\n\n**{ProblemDetailsHTML}**: All problem event details including root cause as an HTML-formatted string.\n\n**{ProblemDetailsText}**: All problem event details including root cause as a text-formatted string.\n\n**{ProblemID}**: Display number of the reported problem.\n\n**{ProblemImpact}**: Impact level of the problem. Possible values are APPLICATION, SERVICE, or INFRASTRUCTURE.\n\n**{ProblemSeverity}**: Severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, or CUSTOM_ALERT.\n\n**{ProblemTitle}**: Short description of the problem.\n\n**{State}**: Problem state. Possible values are OPEN or RESOLVED.\n\n**{Tags}**: Comma separated list of tags that are defined for all impacted entities. To refer to the value of a specific tag, specify the tag's key in square brackets: **{Tags[key]}**. If the tag does not have any assigned value, the placeholder will be replaced by an empty string. The placeholder will not be replaced if the tag key does not exist. To refer to the value of a specific tag, specify the tag's key in square brackets: **{Tags[key]}**. If the tag does not have any assigned value, the placeholder will be replaced by an empty string. The placeholder will not be replaced if the tag key does not exist.
	Password      string `json:"password"`      // The password to the ServiceNow account.
	SendEvents    bool   `json:"sendEvents"`    // Send events into ServiceNow ITOM.
	SendIncidents bool   `json:"sendIncidents"` // Send incidents into ServiceNow ITSM.
	Url           string `json:"url"`           // The URL of the on-premise ServiceNow installation. \n\n This field is mutually exclusive with the **instanceName** field. You can only use one of them.
	Username      string `json:"username"`      // The username of the ServiceNow account. \n\n Make sure that your user account has the `web_service_admin` and `x_dynat_ruxit.Integration` roles.
}

func (me *ServiceNowNotification) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"instance_name": {
			Type:        hcl.TypeString,
			Description: "The ServiceNow instance identifier. It refers to the first part of your own ServiceNow URL. \n\n This field is mutually exclusive with the **url** field. You can only use one of them.",
			Required:    true,
		},
		"message": {
			Type:        hcl.TypeString,
			Description: "The content of the ServiceNow description. Type '{' for placeholder suggestions.. #### Available placeholders\n**{ImpactedEntity}**: A short description of the problem and impacted entity (or multiple impacted entities).\n\n**{ImpactedEntityNames}**: The entity impacted by the problem (or multiple impacted entities).\n\n**{PID}**: Unique system identifier of the reported problem.\n\n**{ProblemDetailsHTML}**: All problem event details including root cause as an HTML-formatted string.\n\n**{ProblemDetailsText}**: All problem event details including root cause as a text-formatted string.\n\n**{ProblemID}**: Display number of the reported problem.\n\n**{ProblemImpact}**: Impact level of the problem. Possible values are APPLICATION, SERVICE, or INFRASTRUCTURE.\n\n**{ProblemSeverity}**: Severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, or CUSTOM_ALERT.\n\n**{ProblemTitle}**: Short description of the problem.\n\n**{State}**: Problem state. Possible values are OPEN or RESOLVED.\n\n**{Tags}**: Comma separated list of tags that are defined for all impacted entities. To refer to the value of a specific tag, specify the tag's key in square brackets: **{Tags[key]}**. If the tag does not have any assigned value, the placeholder will be replaced by an empty string. The placeholder will not be replaced if the tag key does not exist. To refer to the value of a specific tag, specify the tag's key in square brackets: **{Tags[key]}**. If the tag does not have any assigned value, the placeholder will be replaced by an empty string. The placeholder will not be replaced if the tag key does not exist.",
			Required:    true,
		},
		"password": {
			Type:        hcl.TypeString,
			Description: "The password to the ServiceNow account.",
			Required:    true,
		},
		"send_events": {
			Type:        hcl.TypeBool,
			Description: "Send events into ServiceNow ITOM.",
			Optional:    true,
		},
		"send_incidents": {
			Type:        hcl.TypeBool,
			Description: "Send incidents into ServiceNow ITSM.",
			Optional:    true,
		},
		"url": {
			Type:        hcl.TypeString,
			Description: "The URL of the on-premise ServiceNow installation. \n\n This field is mutually exclusive with the **instanceName** field. You can only use one of them.",
			Required:    true,
		},
		"username": {
			Type:        hcl.TypeString,
			Description: "The username of the ServiceNow account. \n\n Make sure that your user account has the `web_service_admin` and `x_dynat_ruxit.Integration` roles.",
			Required:    true,
		},
	}
}

func (me *ServiceNowNotification) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"instance_name":  me.InstanceName,
		"message":        me.Message,
		"password":       me.Password,
		"send_events":    me.SendEvents,
		"send_incidents": me.SendIncidents,
		"url":            me.Url,
		"username":       me.Username,
	})
}

func (me *ServiceNowNotification) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"instance_name":  &me.InstanceName,
		"message":        &me.Message,
		"password":       &me.Password,
		"send_events":    &me.SendEvents,
		"send_incidents": &me.SendIncidents,
		"url":            &me.Url,
		"username":       &me.Username,
	})
}
