package notifications

import "github.com/dtcookie/hcl"

type XMatters struct {
	Enabled   bool   `json:"-"`
	Name      string `json:"-"`
	ProfileID string `json:"-"`

	URL      string      `json:"url"`                  // The URL of the xMatters webhook
	Insecure bool        `json:"acceptAnyCertificate"` // Accept any SSL certificate (including self-signed and invalid certificates)
	Headers  HTTPHeaders `json:"headers,omitempty"`    // Additional HTTP headers
	Payload  string      `json:"payload"`              // The content of the notification message. Type '{' for placeholder suggestions
}

func (me *XMatters) Schema() map[string]*hcl.Schema {
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

		"url": {
			Type:        hcl.TypeString,
			Description: "The URL of the WebHook endpoint",
			Required:    true,
		},
		"insecure": {
			Type:        hcl.TypeBool,
			Description: "Accept any, including self-signed and invalid, SSL certificate (`true`) or only trusted (`false`) certificates",
			Optional:    true,
		},
		"headers": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			MaxItems:    1,
			Description: "A list of the additional HTTP headers",
			Elem:        &hcl.Resource{Schema: new(HTTPHeaders).Schema()},
		},
		"payload": {
			Type:        hcl.TypeString,
			Description: "The content of the notification message. You can use the following placeholders:  * `{ImpactedEntities}`: Details about the entities impacted by the problem in form of a JSON array.  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{PID}`: The ID of the reported problem.  * `{ProblemDetailsHTML}`: All problem event details, including root cause, as an HTML-formatted string.  * `{ProblemDetailsJSON}`: All problem event details, including root cause, as a JSON object.  * `{ProblemDetailsMarkdown}`: All problem event details, including root cause, as a [Markdown-formatted](https://www.markdownguide.org/cheat-sheet/) string.  * `{ProblemDetailsText}`: All problem event details, including root cause, as a text-formatted string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`.  * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas",
			Required:    true,
		},
	}
}

func (me *XMatters) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"name":    me.Name,
		"active":  me.Enabled,
		"profile": me.ProfileID,

		"url":      me.URL,
		"insecure": me.Insecure,
		"headers":  me.Headers,
		"payload":  me.Payload,
	})
}

func (me *XMatters) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"name":    &me.Name,
		"active":  &me.Enabled,
		"profile": &me.ProfileID,

		"url":      &me.URL,
		"insecure": &me.Insecure,
		"headers":  &me.Headers,
		"payload":  &me.Payload,
	})
}
