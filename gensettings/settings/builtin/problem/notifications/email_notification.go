package notifications

import "github.com/dtcookie/hcl"

type EmailNotification struct {
	BccRecipients        []string `json:"bccRecipients"`        // BCC
	Body                 string   `json:"body"`                 // The template of the email notifications. Type '{' for placeholder suggestions.. #### Available placeholders\n**{ImpactedEntities}**: Details about the entities impacted by the problem in form of a json array.\n\n**{ImpactedEntity}**: A short description of the problem and impacted entity (or multiple impacted entities).\n\n**{ImpactedEntityNames}**: The entity impacted by the problem (or multiple impacted entities).\n\n**{PID}**: Unique system identifier of the reported problem.\n\n**{ProblemDetailsHTML}**: All problem event details including root cause as an HTML-formatted string.\n\n**{ProblemDetailsJSONv2}**: Problem as json object following the structure from the [Dynatrace Problems V2 API](https://dt-url.net/7a03ti2). The optional fields evidenceDetails and impactAnalysis are included, but recentComments is not.\n\n**{ProblemDetailsJSON}**: Problem as json object following the structure from the [Dynatrace Problems V1 API](https://dt-url.net/qn23tk2).\n\n**{ProblemDetailsMarkdown}**: All problem event details including root cause as a Markdown-formatted string.\n\n**{ProblemDetailsText}**: All problem event details including root cause as a text-formatted string.\n\n**{ProblemID}**: Display number of the reported problem.\n\n**{ProblemImpact}**: Impact level of the problem. Possible values are APPLICATION, SERVICE, or INFRASTRUCTURE.\n\n**{ProblemSeverity}**: Severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, or CUSTOM_ALERT.\n\n**{ProblemTitle}**: Short description of the problem.\n\n**{ProblemURL}**: URL of the problem within Dynatrace.\n\n**{State}**: Problem state. Possible values are OPEN or RESOLVED.\n\n**{Tags}**: Comma separated list of tags that are defined for all impacted entities. To refer to the value of a specific tag, specify the tag's key in square brackets: **{Tags[key]}**. If the tag does not have any assigned value, the placeholder will be replaced by an empty string. The placeholder will not be replaced if the tag key does not exist.
	CcRecipients         []string `json:"ccRecipients"`         // CC
	NotifyClosedProblems bool     `json:"notifyClosedProblems"` // Send email if problem is closed
	Recipients           []string `json:"recipients"`           // To
	Subject              string   `json:"subject"`              // The subject of the email notifications. Type '{' for placeholder suggestions.. #### Available placeholders\n**{ImpactedEntity}**: A short description of the problem and impacted entity (or multiple impacted entities).\n\n**{ImpactedEntityNames}**: The entity impacted by the problem (or multiple impacted entities).\n\n**{PID}**: Unique system identifier of the reported problem.\n\n**{ProblemID}**: Display number of the reported problem.\n\n**{ProblemImpact}**: Impact level of the problem. Possible values are APPLICATION, SERVICE, or INFRASTRUCTURE.\n\n**{ProblemSeverity}**: Severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, or CUSTOM_ALERT.\n\n**{ProblemTitle}**: Short description of the problem.\n\n**{ProblemURL}**: URL of the problem within Dynatrace.\n\n**{State}**: Problem state. Possible values are OPEN or RESOLVED.\n\n**{Tags}**: Comma separated list of tags that are defined for all impacted entities. To refer to the value of a specific tag, specify the tag's key in square brackets: **{Tags[key]}**. If the tag does not have any assigned value, the placeholder will be replaced by an empty string. The placeholder will not be replaced if the tag key does not exist.
}

func (me *EmailNotification) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"bcc_recipients": {
			Type:        hcl.TypeSet,
			Description: "BCC",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"body": {
			Type:        hcl.TypeString,
			Description: "The template of the email notifications. Type '{' for placeholder suggestions.. #### Available placeholders\n**{ImpactedEntities}**: Details about the entities impacted by the problem in form of a json array.\n\n**{ImpactedEntity}**: A short description of the problem and impacted entity (or multiple impacted entities).\n\n**{ImpactedEntityNames}**: The entity impacted by the problem (or multiple impacted entities).\n\n**{PID}**: Unique system identifier of the reported problem.\n\n**{ProblemDetailsHTML}**: All problem event details including root cause as an HTML-formatted string.\n\n**{ProblemDetailsJSONv2}**: Problem as json object following the structure from the [Dynatrace Problems V2 API](https://dt-url.net/7a03ti2). The optional fields evidenceDetails and impactAnalysis are included, but recentComments is not.\n\n**{ProblemDetailsJSON}**: Problem as json object following the structure from the [Dynatrace Problems V1 API](https://dt-url.net/qn23tk2).\n\n**{ProblemDetailsMarkdown}**: All problem event details including root cause as a Markdown-formatted string.\n\n**{ProblemDetailsText}**: All problem event details including root cause as a text-formatted string.\n\n**{ProblemID}**: Display number of the reported problem.\n\n**{ProblemImpact}**: Impact level of the problem. Possible values are APPLICATION, SERVICE, or INFRASTRUCTURE.\n\n**{ProblemSeverity}**: Severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, or CUSTOM_ALERT.\n\n**{ProblemTitle}**: Short description of the problem.\n\n**{ProblemURL}**: URL of the problem within Dynatrace.\n\n**{State}**: Problem state. Possible values are OPEN or RESOLVED.\n\n**{Tags}**: Comma separated list of tags that are defined for all impacted entities. To refer to the value of a specific tag, specify the tag's key in square brackets: **{Tags[key]}**. If the tag does not have any assigned value, the placeholder will be replaced by an empty string. The placeholder will not be replaced if the tag key does not exist.",
			Required:    true,
		},
		"cc_recipients": {
			Type:        hcl.TypeSet,
			Description: "CC",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"notify_closed_problems": {
			Type:        hcl.TypeBool,
			Description: "Send email if problem is closed",
			Optional:    true,
		},
		"recipients": {
			Type:        hcl.TypeSet,
			Description: "To",
			Required:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"subject": {
			Type:        hcl.TypeString,
			Description: "The subject of the email notifications. Type '{' for placeholder suggestions.. #### Available placeholders\n**{ImpactedEntity}**: A short description of the problem and impacted entity (or multiple impacted entities).\n\n**{ImpactedEntityNames}**: The entity impacted by the problem (or multiple impacted entities).\n\n**{PID}**: Unique system identifier of the reported problem.\n\n**{ProblemID}**: Display number of the reported problem.\n\n**{ProblemImpact}**: Impact level of the problem. Possible values are APPLICATION, SERVICE, or INFRASTRUCTURE.\n\n**{ProblemSeverity}**: Severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, or CUSTOM_ALERT.\n\n**{ProblemTitle}**: Short description of the problem.\n\n**{ProblemURL}**: URL of the problem within Dynatrace.\n\n**{State}**: Problem state. Possible values are OPEN or RESOLVED.\n\n**{Tags}**: Comma separated list of tags that are defined for all impacted entities. To refer to the value of a specific tag, specify the tag's key in square brackets: **{Tags[key]}**. If the tag does not have any assigned value, the placeholder will be replaced by an empty string. The placeholder will not be replaced if the tag key does not exist.",
			Required:    true,
		},
	}
}

func (me *EmailNotification) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"bcc_recipients":         me.BccRecipients,
		"body":                   me.Body,
		"cc_recipients":          me.CcRecipients,
		"notify_closed_problems": me.NotifyClosedProblems,
		"recipients":             me.Recipients,
		"subject":                me.Subject,
	})
}

func (me *EmailNotification) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"bcc_recipients":         &me.BccRecipients,
		"body":                   &me.Body,
		"cc_recipients":          &me.CcRecipients,
		"notify_closed_problems": &me.NotifyClosedProblems,
		"recipients":             &me.Recipients,
		"subject":                &me.Subject,
	})
}
