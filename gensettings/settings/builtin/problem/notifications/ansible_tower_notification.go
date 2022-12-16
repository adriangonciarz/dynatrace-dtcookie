package notifications

import "github.com/dtcookie/hcl"

type AnsibleTowerNotification struct {
	AcceptAnyCertificate bool   `json:"acceptAnyCertificate"` // Accept any SSL certificate (including self-signed and invalid certificates)
	CustomMessage        string `json:"customMessage"`        // This message will be displayed in the Extra Variables **Message** field of your job template. Type '{' for placeholder suggestions.. #### Available placeholders\n**{ImpactedEntities}**: Details about the entities impacted by the problem in form of a json array.\n\n**{ImpactedEntity}**: A short description of the problem and impacted entity (or multiple impacted entities).\n\n**{ImpactedEntityNames}**: The entity impacted by the problem (or multiple impacted entities).\n\n**{PID}**: Unique system identifier of the reported problem.\n\n**{ProblemDetailsText}**: All problem event details including root cause as a text-formatted string.\n\n**{ProblemID}**: Display number of the reported problem.\n\n**{ProblemImpact}**: Impact level of the problem. Possible values are APPLICATION, SERVICE, or INFRASTRUCTURE.\n\n**{ProblemSeverity}**: Severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, or CUSTOM_ALERT.\n\n**{ProblemTitle}**: Short description of the problem.\n\n**{ProblemURL}**: URL of the problem within Dynatrace.\n\n**{State}**: Problem state. Possible values are OPEN or RESOLVED.\n\n**{Tags}**: Comma separated list of tags that are defined for all impacted entities. To refer to the value of a specific tag, specify the tag's key in square brackets: **{Tags[key]}**. If the tag does not have any assigned value, the placeholder will be replaced by an empty string. The placeholder will not be replaced if the tag key does not exist.
	JobTemplateURL       string `json:"jobTemplateURL"`       // The URL of the target job template.\n\nFor example, https://<Ansible server name>/#/templates/job_template/<JobTemplateID>\n\n**Note:** Be sure to select the **Prompt on Launch** option in the Extra Variables section of your job template configuration.
	Password             string `json:"password"`             // Account password.
	Username             string `json:"username"`             // Account username.
}

func (me *AnsibleTowerNotification) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"accept_any_certificate": {
			Type:        hcl.TypeBool,
			Description: "Accept any SSL certificate (including self-signed and invalid certificates)",
			Optional:    true,
		},
		"custom_message": {
			Type:        hcl.TypeString,
			Description: "This message will be displayed in the Extra Variables **Message** field of your job template. Type '{' for placeholder suggestions.. #### Available placeholders\n**{ImpactedEntities}**: Details about the entities impacted by the problem in form of a json array.\n\n**{ImpactedEntity}**: A short description of the problem and impacted entity (or multiple impacted entities).\n\n**{ImpactedEntityNames}**: The entity impacted by the problem (or multiple impacted entities).\n\n**{PID}**: Unique system identifier of the reported problem.\n\n**{ProblemDetailsText}**: All problem event details including root cause as a text-formatted string.\n\n**{ProblemID}**: Display number of the reported problem.\n\n**{ProblemImpact}**: Impact level of the problem. Possible values are APPLICATION, SERVICE, or INFRASTRUCTURE.\n\n**{ProblemSeverity}**: Severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, or CUSTOM_ALERT.\n\n**{ProblemTitle}**: Short description of the problem.\n\n**{ProblemURL}**: URL of the problem within Dynatrace.\n\n**{State}**: Problem state. Possible values are OPEN or RESOLVED.\n\n**{Tags}**: Comma separated list of tags that are defined for all impacted entities. To refer to the value of a specific tag, specify the tag's key in square brackets: **{Tags[key]}**. If the tag does not have any assigned value, the placeholder will be replaced by an empty string. The placeholder will not be replaced if the tag key does not exist.",
			Required:    true,
		},
		"job_template_url": {
			Type:        hcl.TypeString,
			Description: "The URL of the target job template.\n\nFor example, https://<Ansible server name>/#/templates/job_template/<JobTemplateID>\n\n**Note:** Be sure to select the **Prompt on Launch** option in the Extra Variables section of your job template configuration.",
			Required:    true,
		},
		"password": {
			Type:        hcl.TypeString,
			Description: "Account password.",
			Required:    true,
		},
		"username": {
			Type:        hcl.TypeString,
			Description: "Account username.",
			Required:    true,
		},
	}
}

func (me *AnsibleTowerNotification) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"accept_any_certificate": me.AcceptAnyCertificate,
		"custom_message":         me.CustomMessage,
		"job_template_url":       me.JobTemplateURL,
		"password":               me.Password,
		"username":               me.Username,
	})
}

func (me *AnsibleTowerNotification) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"accept_any_certificate": &me.AcceptAnyCertificate,
		"custom_message":         &me.CustomMessage,
		"job_template_url":       &me.JobTemplateURL,
		"password":               &me.Password,
		"username":               &me.Username,
	})
}
