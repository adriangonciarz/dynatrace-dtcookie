package notifications

import (
	"github.com/dtcookie/hcl"
)

type AnsibleTower struct {
	Enabled   bool   `json:"-" hcl:"active,omitempty"` // The notification is active (`true`) or inactive (`false`). Default is `false`
	Name      string `json:"-" hcl:"name"`             // The display name within the Dynatrace WebUI
	ProfileID string `json:"-" hcl:"alerting_profile"` // The ID of the associated alerting profile

	JobTemplateURL string `json:"jobTemplateURL" hcl:"job_template_url"`             // The URL of the target Ansible Tower job template.\n\nFor example, https://<Ansible Tower server name>/#/templates/job_template/<JobTemplateID>\n\n**Note:** Be sure to select the **Prompt on Launch** option in the Extra Variables section of your job template configuration
	Insecure       bool   `json:"acceptAnyCertificate" hcl:"accept_any_certificate"` // Accept any SSL certificate (including self-signed and invalid certificates)
	Username       string `json:"username" hcl:"username"`                           // The username of the Ansible Tower account
	Password       string `json:"password" hcl:"password,secret"`                    // The password for the Ansible Tower account
	CustomMessage  string `json:"customMessage" hcl:"custom_message"`                // This message will be displayed in the Extra Variables **Message** field of your job template. Type '{' for placeholder suggestions
}

func (me *AnsibleTower) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "The display name within the Dynatrace WebUI.",
			Required:    true,
		},
		"active": {
			Type:        hcl.TypeBool,
			Description: "The notification is active (`true`) or inactive (`false`). Default is `false`.",
			Optional:    true,
		},
		"profile": {
			Type:        hcl.TypeString,
			Description: "The ID of the associated alerting profile.",
			Required:    true,
		},

		"insecure": {
			Type:        hcl.TypeBool,
			Description: "Accept any, including self-signed and invalid, SSL certificate (`true`) or only trusted (`false`) certificates. Default is `false`.",
			Optional:    true,
		},
		"custom_message": {
			Type:        hcl.TypeString,
			Description: "The custom message of the notification. This message will be displayed in the extra variables **Message** field of your job template. You can use the following placeholders:  * `{ImpactedEntities}`: Details about the entities impacted by the problem in form of a JSON array.  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{PID}`: The ID of the reported problem.  * `{ProblemDetailsText}`: All problem event details, including root cause, as a text-formatted string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`.  * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas",
			Required:    true,
		},
		"job_template_url": {
			Type:        hcl.TypeString,
			Description: "The URL of the target Ansible Tower job template",
			Required:    true,
		},
		"password": {
			Type:        hcl.TypeString,
			Sensitive:   true,
			Description: "The password for the Ansible Tower account",
			Optional:    true,
		},
		"username": {
			Type:        hcl.TypeString,
			Description: "The username of the Ansible Tower account",
			Required:    true,
		},
	}
}

func (me *AnsibleTower) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	// The password field MUST NOT get serialized into HCL here
	// The Dynatrace Settings 2.0 API delivers a scrambled version of any previously stored password here
	// Evaluation at this point would lead to that scrambled version to make it into the Terraform State
	// As a result any plans would be non-empty
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"name":    me.Name,
		"active":  me.Enabled,
		"profile": me.ProfileID,

		"insecure":         me.Insecure,
		"custom_message":   me.CustomMessage,
		"job_template_url": me.JobTemplateURL,
		// "password":               me.Password,
		"username": me.Username,
	})
}

func (me *AnsibleTower) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"name":    &me.Name,
		"active":  &me.Enabled,
		"profile": &me.ProfileID,

		"insecure":         &me.Insecure,
		"custom_message":   &me.CustomMessage,
		"job_template_url": &me.JobTemplateURL,
		"password":         &me.Password,
		"username":         &me.Username,
	})
}
