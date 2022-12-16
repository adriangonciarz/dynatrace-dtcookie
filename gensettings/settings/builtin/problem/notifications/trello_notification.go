package notifications

import "github.com/dtcookie/hcl"

type TrelloNotification struct {
	ApplicationKey     string `json:"applicationKey"`     // The application key for the Trello account.\n\nYou must be logged into Trello to have Trello automatically generate an application key for you. [Get application key](https://trello.com/app-key)
	AuthorizationToken string `json:"authorizationToken"` // The authorization token for the Trello account.
	BoardID            string `json:"boardId"`            // Trello board ID problem cards should be assigned to
	Description        string `json:"description"`        // The description of the Trello card. Type '{' for placeholder suggestions.. #### Available placeholders\n**{ImpactedEntity}**: A short description of the problem and impacted entity (or multiple impacted entities).\n\n**{ImpactedEntityNames}**: The entity impacted by the problem (or multiple impacted entities).\n\n**{PID}**: Unique system identifier of the reported problem.\n\n**{ProblemDetailsMarkdown}**: All problem event details including root cause as a Markdown-formatted string.\n\n**{ProblemID}**: Display number of the reported problem.\n\n**{ProblemImpact}**: Impact level of the problem. Possible values are APPLICATION, SERVICE, or INFRASTRUCTURE.\n\n**{ProblemSeverity}**: Severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, or CUSTOM_ALERT.\n\n**{ProblemTitle}**: Short description of the problem.\n\n**{ProblemURL}**: URL of the problem within Dynatrace.\n\n**{State}**: Problem state. Possible values are OPEN or RESOLVED.\n\n**{Tags}**: Comma separated list of tags that are defined for all impacted entities. To refer to the value of a specific tag, specify the tag's key in square brackets: **{Tags[key]}**. If the tag does not have any assigned value, the placeholder will be replaced by an empty string. The placeholder will not be replaced if the tag key does not exist.
	ListID             string `json:"listId"`             // Trello list ID new problem cards should be assigned to
	ResolvedListID     string `json:"resolvedListId"`     // Trello list ID resolved problem cards should be assigned to
	Text               string `json:"text"`               // The card text and problem placeholders to appear on new problem cards. Type '{' for placeholder suggestions.. #### Available placeholders\n**{ImpactedEntity}**: A short description of the problem and impacted entity (or multiple impacted entities).\n\n**{ImpactedEntityNames}**: The entity impacted by the problem (or multiple impacted entities).\n\n**{PID}**: Unique system identifier of the reported problem.\n\n**{ProblemID}**: Display number of the reported problem.\n\n**{ProblemImpact}**: Impact level of the problem. Possible values are APPLICATION, SERVICE, or INFRASTRUCTURE.\n\n**{ProblemSeverity}**: Severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, or CUSTOM_ALERT.\n\n**{ProblemTitle}**: Short description of the problem.\n\n**{ProblemURL}**: URL of the problem within Dynatrace.\n\n**{State}**: Problem state. Possible values are OPEN or RESOLVED.\n\n**{Tags}**: Comma separated list of tags that are defined for all impacted entities. To refer to the value of a specific tag, specify the tag's key in square brackets: **{Tags[key]}**. If the tag does not have any assigned value, the placeholder will be replaced by an empty string. The placeholder will not be replaced if the tag key does not exist.
}

func (me *TrelloNotification) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"application_key": {
			Type:        hcl.TypeString,
			Description: "The application key for the Trello account.\n\nYou must be logged into Trello to have Trello automatically generate an application key for you. [Get application key](https://trello.com/app-key)",
			Required:    true,
		},
		"authorization_token": {
			Type:        hcl.TypeString,
			Description: "The authorization token for the Trello account.",
			Required:    true,
		},
		"board_id": {
			Type:        hcl.TypeString,
			Description: "Trello board ID problem cards should be assigned to",
			Required:    true,
		},
		"description": {
			Type:        hcl.TypeString,
			Description: "The description of the Trello card. Type '{' for placeholder suggestions.. #### Available placeholders\n**{ImpactedEntity}**: A short description of the problem and impacted entity (or multiple impacted entities).\n\n**{ImpactedEntityNames}**: The entity impacted by the problem (or multiple impacted entities).\n\n**{PID}**: Unique system identifier of the reported problem.\n\n**{ProblemDetailsMarkdown}**: All problem event details including root cause as a Markdown-formatted string.\n\n**{ProblemID}**: Display number of the reported problem.\n\n**{ProblemImpact}**: Impact level of the problem. Possible values are APPLICATION, SERVICE, or INFRASTRUCTURE.\n\n**{ProblemSeverity}**: Severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, or CUSTOM_ALERT.\n\n**{ProblemTitle}**: Short description of the problem.\n\n**{ProblemURL}**: URL of the problem within Dynatrace.\n\n**{State}**: Problem state. Possible values are OPEN or RESOLVED.\n\n**{Tags}**: Comma separated list of tags that are defined for all impacted entities. To refer to the value of a specific tag, specify the tag's key in square brackets: **{Tags[key]}**. If the tag does not have any assigned value, the placeholder will be replaced by an empty string. The placeholder will not be replaced if the tag key does not exist.",
			Required:    true,
		},
		"list_id": {
			Type:        hcl.TypeString,
			Description: "Trello list ID new problem cards should be assigned to",
			Required:    true,
		},
		"resolved_list_id": {
			Type:        hcl.TypeString,
			Description: "Trello list ID resolved problem cards should be assigned to",
			Required:    true,
		},
		"text": {
			Type:        hcl.TypeString,
			Description: "The card text and problem placeholders to appear on new problem cards. Type '{' for placeholder suggestions.. #### Available placeholders\n**{ImpactedEntity}**: A short description of the problem and impacted entity (or multiple impacted entities).\n\n**{ImpactedEntityNames}**: The entity impacted by the problem (or multiple impacted entities).\n\n**{PID}**: Unique system identifier of the reported problem.\n\n**{ProblemID}**: Display number of the reported problem.\n\n**{ProblemImpact}**: Impact level of the problem. Possible values are APPLICATION, SERVICE, or INFRASTRUCTURE.\n\n**{ProblemSeverity}**: Severity level of the problem. Possible values are AVAILABILITY, ERROR, PERFORMANCE, RESOURCE_CONTENTION, or CUSTOM_ALERT.\n\n**{ProblemTitle}**: Short description of the problem.\n\n**{ProblemURL}**: URL of the problem within Dynatrace.\n\n**{State}**: Problem state. Possible values are OPEN or RESOLVED.\n\n**{Tags}**: Comma separated list of tags that are defined for all impacted entities. To refer to the value of a specific tag, specify the tag's key in square brackets: **{Tags[key]}**. If the tag does not have any assigned value, the placeholder will be replaced by an empty string. The placeholder will not be replaced if the tag key does not exist.",
			Required:    true,
		},
	}
}

func (me *TrelloNotification) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"application_key":     me.ApplicationKey,
		"authorization_token": me.AuthorizationToken,
		"board_id":            me.BoardID,
		"description":         me.Description,
		"list_id":             me.ListID,
		"resolved_list_id":    me.ResolvedListID,
		"text":                me.Text,
	})
}

func (me *TrelloNotification) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"application_key":     &me.ApplicationKey,
		"authorization_token": &me.AuthorizationToken,
		"board_id":            &me.BoardID,
		"description":         &me.Description,
		"list_id":             &me.ListID,
		"resolved_list_id":    &me.ResolvedListID,
		"text":                &me.Text,
	})
}
