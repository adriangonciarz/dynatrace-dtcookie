package notifications

import "github.com/dtcookie/hcl"

type OpsGenieNotification struct {
	ApiKey  string `json:"apiKey"`  // The API key to access OpsGenie.\n\nGo to OpsGenie-Integrations and create a new Dynatrace integration. Copy the newly created API key.
	Domain  string `json:"domain"`  // The region domain of the OpsGenie.\n\nFor example, **api.opsgenie.com** for US or **api.eu.opsgenie.com** for EU.
	Message string `json:"message"` // The content of the message. Type '{' for placeholder suggestions.. #### Available placeholders\n**{ProblemID}**: Display number of the reported problem.\n\n**{ProblemImpact}**: Impact level of the problem. Possible values are APPLICATION, SERVICE, or INFRASTRUCTURE.\n\n**{ProblemSeverity}**: Severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, or CUSTOM_ALERT.\n\n**{ProblemTitle}**: Short description of the problem.\n\n**{ImpactedEntityNames}**: The entity impacted by the problem (or multiple impacted entities).
}

func (me *OpsGenieNotification) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"api_key": {
			Type:        hcl.TypeString,
			Description: "The API key to access OpsGenie.\n\nGo to OpsGenie-Integrations and create a new Dynatrace integration. Copy the newly created API key.",
			Required:    true,
		},
		"domain": {
			Type:        hcl.TypeString,
			Description: "The region domain of the OpsGenie.\n\nFor example, **api.opsgenie.com** for US or **api.eu.opsgenie.com** for EU.",
			Required:    true,
		},
		"message": {
			Type:        hcl.TypeString,
			Description: "The content of the message. Type '{' for placeholder suggestions.. #### Available placeholders\n**{ProblemID}**: Display number of the reported problem.\n\n**{ProblemImpact}**: Impact level of the problem. Possible values are APPLICATION, SERVICE, or INFRASTRUCTURE.\n\n**{ProblemSeverity}**: Severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, or CUSTOM_ALERT.\n\n**{ProblemTitle}**: Short description of the problem.\n\n**{ImpactedEntityNames}**: The entity impacted by the problem (or multiple impacted entities).",
			Required:    true,
		},
	}
}

func (me *OpsGenieNotification) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"api_key": me.ApiKey,
		"domain":  me.Domain,
		"message": me.Message,
	})
}

func (me *OpsGenieNotification) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"api_key": &me.ApiKey,
		"domain":  &me.Domain,
		"message": &me.Message,
	})
}
