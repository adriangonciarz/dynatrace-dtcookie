package notifications

import "github.com/dtcookie/hcl"

type VictorOps struct {
	Enabled   bool   `json:"-"`
	Name      string `json:"-"`
	ProfileID string `json:"-"`

	APIKey     string `json:"apiKey"`     // The API key for the target VictorOps account.\n\nReceive your VictorOps API key by navigating to: Settings -> Integrations -> Rest Endpoint -> Dynatrace within your VictorOps account
	RoutingKey string `json:"routingKey"` // The routing key, defining the group to be notified
	Message    string `json:"message"`    // The content of the message. Type '{' for placeholder suggestions
}

func (me *VictorOps) Schema() map[string]*hcl.Schema {
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

		"api_key": {
			Type:        hcl.TypeString,
			Description: "The API key for the target VictorOps account",
			Sensitive:   true,
			Optional:    true,
		},
		"message": {
			Type:        hcl.TypeString,
			Description: "The content of the message.  You can use the following placeholders:  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{ProblemDetailsText}`: All problem event details, including root cause, as a text-formatted string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`",
			Required:    true,
		},
		"routing_key": {
			Type:        hcl.TypeString,
			Description: "The routing key, defining the group to be notified",
			Required:    true,
		},
	}
}

func (me *VictorOps) MarshalHCL() (map[string]interface{}, error) {
	// The api_key field MUST NOT get serialized into HCL here
	// The Dynatrace Settings 2.0 API delivers a scrambled version of any previously stored api_key here
	// Evaluation at this point would lead to that scrambled version to make it into the Terraform State
	// As a result any plans would be non-empty
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"name":    me.Name,
		"active":  me.Enabled,
		"profile": me.ProfileID,

		"message":     me.Message,
		"routing_key": me.RoutingKey,
		// "api_key":     me.APIKey,
	})
}

func (me *VictorOps) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"name":    &me.Name,
		"active":  &me.Enabled,
		"profile": &me.ProfileID,

		"message":     &me.Message,
		"routing_key": &me.RoutingKey,
		"api_key":     &me.APIKey,
	})
}
