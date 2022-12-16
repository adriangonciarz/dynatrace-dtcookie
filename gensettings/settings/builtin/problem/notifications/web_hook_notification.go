package notifications

import "github.com/dtcookie/hcl"

type WebHookNotification struct {
	AcceptAnyCertificate     bool                       `json:"acceptAnyCertificate"`     // Accept any SSL certificate (including self-signed and invalid certificates)
	Headers                  WebHookNotificationHeaders `json:"headers"`                  // A list of the additional HTTP headers.
	NotifyClosedProblems     bool                       `json:"notifyClosedProblems"`     // Call webhook if problem is closed
	NotifyEventMergesEnabled bool                       `json:"notifyEventMergesEnabled"` // Call webhook if new events merge into existing problems
	Payload                  string                     `json:"payload"`                  // The content of the notification message. Type '{' for placeholder suggestions.. #### Available placeholders\n**{ImpactedEntities}**: Details about the entities impacted by the problem in form of a json array.\n\n**{ImpactedEntity}**: A short description of the problem and impacted entity (or multiple impacted entities).\n\n**{ImpactedEntityNames}**: The entity impacted by the problem (or multiple impacted entities).\n\n**{PID}**: Unique system identifier of the reported problem.\n\n**{ProblemDetailsHTML}**: All problem event details including root cause as an HTML-formatted string.\n\n**{ProblemDetailsJSONv2}**: Problem as json object following the structure from the [Dynatrace Problems V2 API](https://dt-url.net/7a03ti2). The optional fields evidenceDetails and impactAnalysis are included, but recentComments is not.\n\n**{ProblemDetailsJSON}**: Problem as json object following the structure from the [Dynatrace Problems V1 API](https://dt-url.net/qn23tk2).\n\n**{ProblemDetailsMarkdown}**: All problem event details including root cause as a Markdown-formatted string.\n\n**{ProblemDetailsText}**: All problem event details including root cause as a text-formatted string.\n\n**{ProblemID}**: Display number of the reported problem.\n\n**{ProblemImpact}**: Impact level of the problem. Possible values are APPLICATION, SERVICE, or INFRASTRUCTURE.\n\n**{ProblemSeverity}**: Severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, or CUSTOM_ALERT.\n\n**{ProblemTitle}**: Short description of the problem.\n\n**{ProblemURL}**: URL of the problem within Dynatrace.\n\n**{State}**: Problem state. Possible values are OPEN or RESOLVED.\n\n**{Tags}**: Comma separated list of tags that are defined for all impacted entities. To refer to the value of a specific tag, specify the tag's key in square brackets: **{Tags[key]}**. If the tag does not have any assigned value, the placeholder will be replaced by an empty string. The placeholder will not be replaced if the tag key does not exist.
	Url                      string                     `json:"url"`                      // The URL of the WebHook endpoint.
}

func (me *WebHookNotification) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"accept_any_certificate": {
			Type:        hcl.TypeBool,
			Description: "Accept any SSL certificate (including self-signed and invalid certificates)",
			Optional:    true,
		},
		"headers": {
			Type:        hcl.TypeList,
			Description: "A list of the additional HTTP headers.",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(WebHookNotificationHeaders).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"notify_closed_problems": {
			Type:        hcl.TypeBool,
			Description: "Call webhook if problem is closed",
			Optional:    true,
		},
		"notify_event_merges_enabled": {
			Type:        hcl.TypeBool,
			Description: "Call webhook if new events merge into existing problems",
			Optional:    true,
		},
		"payload": {
			Type:        hcl.TypeString,
			Description: "The content of the notification message. Type '{' for placeholder suggestions.. #### Available placeholders\n**{ImpactedEntities}**: Details about the entities impacted by the problem in form of a json array.\n\n**{ImpactedEntity}**: A short description of the problem and impacted entity (or multiple impacted entities).\n\n**{ImpactedEntityNames}**: The entity impacted by the problem (or multiple impacted entities).\n\n**{PID}**: Unique system identifier of the reported problem.\n\n**{ProblemDetailsHTML}**: All problem event details including root cause as an HTML-formatted string.\n\n**{ProblemDetailsJSONv2}**: Problem as json object following the structure from the [Dynatrace Problems V2 API](https://dt-url.net/7a03ti2). The optional fields evidenceDetails and impactAnalysis are included, but recentComments is not.\n\n**{ProblemDetailsJSON}**: Problem as json object following the structure from the [Dynatrace Problems V1 API](https://dt-url.net/qn23tk2).\n\n**{ProblemDetailsMarkdown}**: All problem event details including root cause as a Markdown-formatted string.\n\n**{ProblemDetailsText}**: All problem event details including root cause as a text-formatted string.\n\n**{ProblemID}**: Display number of the reported problem.\n\n**{ProblemImpact}**: Impact level of the problem. Possible values are APPLICATION, SERVICE, or INFRASTRUCTURE.\n\n**{ProblemSeverity}**: Severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, or CUSTOM_ALERT.\n\n**{ProblemTitle}**: Short description of the problem.\n\n**{ProblemURL}**: URL of the problem within Dynatrace.\n\n**{State}**: Problem state. Possible values are OPEN or RESOLVED.\n\n**{Tags}**: Comma separated list of tags that are defined for all impacted entities. To refer to the value of a specific tag, specify the tag's key in square brackets: **{Tags[key]}**. If the tag does not have any assigned value, the placeholder will be replaced by an empty string. The placeholder will not be replaced if the tag key does not exist.",
			Required:    true,
		},
		"url": {
			Type:        hcl.TypeString,
			Description: "The URL of the WebHook endpoint.",
			Required:    true,
		},
	}
}

func (me *WebHookNotification) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"accept_any_certificate":      me.AcceptAnyCertificate,
		"headers":                     me.Headers,
		"notify_closed_problems":      me.NotifyClosedProblems,
		"notify_event_merges_enabled": me.NotifyEventMergesEnabled,
		"payload":                     me.Payload,
		"url":                         me.Url,
	})
}

func (me *WebHookNotification) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"accept_any_certificate":      &me.AcceptAnyCertificate,
		"headers":                     &me.Headers,
		"notify_closed_problems":      &me.NotifyClosedProblems,
		"notify_event_merges_enabled": &me.NotifyEventMergesEnabled,
		"payload":                     &me.Payload,
		"url":                         &me.Url,
	})
}
