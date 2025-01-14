package notifications

import (
	"github.com/dtcookie/hcl"
)

type Slack struct {
	Enabled   bool   `json:"-"`
	Name      string `json:"-"`
	ProfileID string `json:"-"`

	URL     string `json:"url"`     // Set up an incoming WebHook integration within your Slack account. Copy and paste the generated WebHook URL into the field above
	Channel string `json:"channel"` // The channel (for example, `#general`) or the user (for example, `@john.smith`) to send the message to
	Message string `json:"message"` // The content of the message. Type '{' for placeholder suggestions
}

func (me *Slack) Schema() map[string]*hcl.Schema {
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

		"message": {
			Type:        hcl.TypeString,
			Description: "The content of the message.  You can use the following placeholders:  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{PID}`: The ID of the reported problem.  * `{ProblemDetailsText}`: All problem event details, including root cause, as a text-formatted string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`.  * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas",
			Required:    true,
		},
		"url": {
			Type:        hcl.TypeString,
			Sensitive:   true,
			Description: "The URL of the Slack WebHook. This is confidential information, therefore GET requests return this field with the `null` value, and it is optional for PUT requests",
			Required:    true,
		},
		"channel": {
			Type:        hcl.TypeString,
			Description: "The channel (for example, `#general`) or the user (for example, `@john.smith`) to send the message to",
			Required:    true,
		},
	}
}

func (me *Slack) MarshalHCL() (map[string]interface{}, error) {
	// The url field MUST NOT get serialized into HCL here
	// The Dynatrace Settings 2.0 API delivers a scrambled version of any previously stored url here
	// Evaluation at this point would lead to that scrambled version to make it into the Terraform State
	// As a result any plans would be non-empty
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"name":    me.Name,
		"active":  me.Enabled,
		"profile": me.ProfileID,

		"message": me.Message,
		"channel": me.Channel,
		"url":     usecret(me.URL),
	})
}

func (me *Slack) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"name":    &me.Name,
		"active":  &me.Enabled,
		"profile": &me.ProfileID,

		"message": &me.Message,
		"channel": &me.Channel,
		"url":     &me.URL,
	})
}
