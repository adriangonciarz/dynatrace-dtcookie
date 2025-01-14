package notifications

import "github.com/dtcookie/hcl"

type ServiceNow struct {
	Enabled   bool   `json:"-"`
	Name      string `json:"-"`
	ProfileID string `json:"-"`

	InstanceName  *string `json:"instanceName"`  // The ServiceNow instance identifier. It refers to the first part of your own ServiceNow URL. \n\n This field is mutually exclusive with the **url** field. You can only use one of them
	URL           *string `json:"url"`           // The URL of the on-premise ServiceNow installation. \n\n This field is mutually exclusive with the **instanceName** field. You can only use one of them
	Username      string  `json:"username"`      // The username of the ServiceNow account. \n\n Make sure that your user account has the `web_service_admin` and `x_dynat_ruxit.Integration` roles
	Password      string  `json:"password"`      // The password to the ServiceNow account
	Message       string  `json:"message"`       // The content of the ServiceNow description. Type '{' for placeholder suggestions
	SendIncidents bool    `json:"sendIncidents"` // Send incidents into ServiceNow ITSM
	SendEvents    bool    `json:"sendEvents"`    // Send events into ServiceNow ITOM
}

func (me *ServiceNow) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "The name of the notification configuration",
			Required:    true,
		},
		"active": {
			Type:        hcl.TypeBool,
			Description: "The configuration is enabled (`true`) or disabled (`false`)",
			Required:    true,
		},
		"profile": {
			Type:        hcl.TypeString,
			Description: "The ID of the associated alerting profile",
			Required:    true,
		},

		"events": {
			Type:        hcl.TypeBool,
			Description: "Send events into ServiceNow ITOM",
			Optional:    true,
		},
		"incidents": {
			Type:        hcl.TypeBool,
			Description: "Send incidents into ServiceNow ITSM",
			Required:    true,
		},
		"url": {
			Type:        hcl.TypeString,
			Description: "The URL of the on-premise ServiceNow installation. This field is mutually exclusive with the **instance** field. You can only use one of them",
			Optional:    true,
		},
		"username": {
			Type:        hcl.TypeString,
			Description: "The username of the ServiceNow account.   Make sure that your user account has the `rest_service`, `web_request_admin`, and `x_dynat_ruxit.Integration` roles",
			Required:    true,
		},
		"instance": {
			Type:        hcl.TypeString,
			Description: "The ServiceNow instance identifier. It refers to the first part of your own ServiceNow URL. This field is mutually exclusive with the **url** field. You can only use one of them",
			Optional:    true,
		},
		"message": {
			Type:        hcl.TypeString,
			Description: "The content of the ServiceNow description. You can use the following placeholders:  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{PID}`: The ID of the reported problem.  * `{ProblemDetailsHTML}`: All problem event details, including root cause, as an HTML-formatted string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`.  * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas",
			Required:    true,
		},
		"password": {
			Type:        hcl.TypeString,
			Description: "The password to the ServiceNow account",
			Sensitive:   true,
			Optional:    true,
		},
	}
}

func (me *ServiceNow) MarshalHCL() (map[string]interface{}, error) {
	// The password field MUST NOT get serialized into HCL here
	// The Dynatrace Settings 2.0 API delivers a scrambled version of any previously stored password here
	// Evaluation at this point would lead to that scrambled version to make it into the Terraform State
	// As a result any plans would be non-empty
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"name":    me.Name,
		"active":  me.Enabled,
		"profile": me.ProfileID,

		"events":    me.SendEvents,
		"incidents": me.SendIncidents,
		"url":       me.URL,
		"username":  me.Username,
		"instance":  me.InstanceName,
		"message":   me.Message,
		"password":  secret(me.Password),
	})
}

func (me *ServiceNow) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"name":    &me.Name,
		"active":  &me.Enabled,
		"profile": &me.ProfileID,

		"events":    &me.SendEvents,
		"incidents": &me.SendIncidents,
		"url":       &me.URL,
		"username":  &me.Username,
		"instance":  &me.InstanceName,
		"message":   &me.Message,
		"password":  &me.Password,
	})
}
