package notifications

import (
	"github.com/dtcookie/hcl"
)

type OpsGenie struct {
	Enabled   bool   `json:"-"`
	Name      string `json:"-"`
	ProfileID string `json:"-"`

	APIKey  *string `json:"apiKey"`  // The API key to access OpsGenie.\n\nGo to OpsGenie-Integrations and create a new Dynatrace integration. Copy the newly created API key
	Domain  string  `json:"domain"`  // The region domain of the OpsGenie.\n\nFor example, **api.opsgenie.com** for US or **api.eu.opsgenie.com** for EU
	Message string  `json:"message"` // The content of the message. Type '{' for placeholder suggestions
}

func (me *OpsGenie) Schema() map[string]*hcl.Schema {
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
			Description: "The API key to access OpsGenie",
			Sensitive:   true,
			Optional:    true,
		},
		"domain": {
			Type:        hcl.TypeString,
			Description: "The region domain of the OpsGenie",
			Required:    true,
		},
		"message": {
			Type:        hcl.TypeString,
			Description: "The content of the message.  You can use the following placeholders:  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem",
			Required:    true,
		},
	}
}

func (me *OpsGenie) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	// The api_key field MUST NOT get serialized into HCL here
	// The Dynatrace Settings 2.0 API delivers a scrambled version of any previously stored api_key here
	// Evaluation at this point would lead to that scrambled version to make it into the Terraform State
	// As a result any plans would be non-empty
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"name":    me.Name,
		"active":  me.Enabled,
		"profile": me.ProfileID,

		"domain":  me.Domain,
		"message": me.Message,
		// "api_key":     me.APIKey,
	})
}

func (me *OpsGenie) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"name":    &me.Name,
		"active":  &me.Enabled,
		"profile": &me.ProfileID,

		"domain":  &me.Domain,
		"message": &me.Message,
		"api_key": &me.APIKey,
	})
}
