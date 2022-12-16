package notifications

import "github.com/dtcookie/hcl"

type VictorOpsNotification struct {
	ApiKey     string `json:"apiKey"`     // The API key for the target VictorOps account.\n\nReceive your VictorOps API key by navigating to: Settings -> Integrations -> Rest Endpoint -> Dynatrace within your VictorOps account.
	Message    string `json:"message"`    // The content of the message. Type '{' for placeholder suggestions.. #### Available placeholders\n**{ImpactedEntity}**: A short description of the problem and impacted entity (or multiple impacted entities).\n\n**{ImpactedEntityNames}**: The entity impacted by the problem (or multiple impacted entities).\n\n**{ProblemDetailsText}**: All problem event details including root cause as a text-formatted string.\n\n**{ProblemID}**: Display number of the reported problem.\n\n**{ProblemImpact}**: Impact level of the problem. Possible values are APPLICATION, SERVICE, or INFRASTRUCTURE.\n\n**{ProblemSeverity}**: Severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, or CUSTOM_ALERT.\n\n**{ProblemTitle}**: Short description of the problem.\n\n**{ProblemURL}**: URL of the problem within Dynatrace.\n\n**{State}**: Problem state. Possible values are OPEN or RESOLVED.
	RoutingKey string `json:"routingKey"` // The routing key, defining the group to be notified.
}

func (me *VictorOpsNotification) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"api_key": {
			Type:        hcl.TypeString,
			Description: "The API key for the target VictorOps account.\n\nReceive your VictorOps API key by navigating to: Settings -> Integrations -> Rest Endpoint -> Dynatrace within your VictorOps account.",
			Required:    true,
		},
		"message": {
			Type:        hcl.TypeString,
			Description: "The content of the message. Type '{' for placeholder suggestions.. #### Available placeholders\n**{ImpactedEntity}**: A short description of the problem and impacted entity (or multiple impacted entities).\n\n**{ImpactedEntityNames}**: The entity impacted by the problem (or multiple impacted entities).\n\n**{ProblemDetailsText}**: All problem event details including root cause as a text-formatted string.\n\n**{ProblemID}**: Display number of the reported problem.\n\n**{ProblemImpact}**: Impact level of the problem. Possible values are APPLICATION, SERVICE, or INFRASTRUCTURE.\n\n**{ProblemSeverity}**: Severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, or CUSTOM_ALERT.\n\n**{ProblemTitle}**: Short description of the problem.\n\n**{ProblemURL}**: URL of the problem within Dynatrace.\n\n**{State}**: Problem state. Possible values are OPEN or RESOLVED.",
			Required:    true,
		},
		"routing_key": {
			Type:        hcl.TypeString,
			Description: "The routing key, defining the group to be notified.",
			Required:    true,
		},
	}
}

func (me *VictorOpsNotification) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"api_key":     me.ApiKey,
		"message":     me.Message,
		"routing_key": me.RoutingKey,
	})
}

func (me *VictorOpsNotification) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"api_key":     &me.ApiKey,
		"message":     &me.Message,
		"routing_key": &me.RoutingKey,
	})
}
