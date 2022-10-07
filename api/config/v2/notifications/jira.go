package notifications

import (
	"github.com/dtcookie/hcl"
)

type Jira struct {
	Enabled   bool   `json:"-"`
	Name      string `json:"-"`
	ProfileID string `json:"-"`

	URL         string `json:"url"`         // Jira endpoint URL
	Username    string `json:"username"`    // The username of the Jira profile
	APIToken    string `json:"apiToken"`    // The API token for the Jira profile. Using password authentication [was deprecated by Jira](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-basic-auth-and-cookie-based-auth/)
	ProjectKey  string `json:"projectKey"`  // The project key of the Jira issue to be created by this notification
	IssueType   string `json:"issueType"`   // The type of the Jira issue to be created by this notification.\n\nTo find all available issue types, or to create your own issue type, within JIRA go to Options > Issues
	Summary     string `json:"summary"`     // The summary of the Jira issue to be created by this notification. Type '{' for placeholder suggestions
	Description string `json:"description"` // The description of the Jira issue to be created by this notification. Type '{' for placeholder suggestions
}

func (me *Jira) Schema() map[string]*hcl.Schema {
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
			Description: "The URL of the Jira API endpoint",
			Required:    true,
		},
		"username": {
			Type:        hcl.TypeString,
			Description: "The username of the Jira profile",
			Required:    true,
		},
		"api_token": {
			Type:        hcl.TypeString,
			Sensitive:   true,
			Description: "The API token for the Jira profile. Using password authentication [was deprecated by Jira](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-basic-auth-and-cookie-based-auth/)",
			Optional:    true,
		},
		"project_key": {
			Type:        hcl.TypeString,
			Description: "The project key of the Jira issue to be created by this notification",
			Required:    true,
		},
		"issue_type": {
			Type:        hcl.TypeString,
			Description: "The type of the Jira issue to be created by this notification",
			Required:    true,
		},
		"summary": {
			Type:        hcl.TypeString,
			Description: "The summary of the Jira issue to be created by this notification.  You can use the following placeholders:  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{PID}`: The ID of the reported problem.  * `{ProblemDetailsText}`: All problem event details, including root cause, as a text-formatted string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`.  * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas",
			Required:    true,
		},
		"description": {
			Type:        hcl.TypeString,
			Description: "The description of the Jira issue to be created by this notification.   You can use same placeholders as in issue summary",
			Required:    true,
		},
	}
}

func (me *Jira) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	// The api_token field MUST NOT get serialized into HCL here
	// The Dynatrace Settings 2.0 API delivers a scrambled version of any previously stored api_token here
	// Evaluation at this point would lead to that scrambled version to make it into the Terraform State
	// As a result any plans would be non-empty
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"name":    me.Name,
		"active":  me.Enabled,
		"profile": me.ProfileID,

		"url":         me.URL,
		"username":    me.Username,
		"project_key": me.ProjectKey,
		"issue_type":  me.IssueType,
		"summary":     me.Summary,
		"description": me.Description,
		// "api_token":        me.APIToken,
	})
}

func (me *Jira) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"name":    &me.Name,
		"active":  &me.Enabled,
		"profile": &me.ProfileID,

		"url":         &me.URL,
		"username":    &me.Username,
		"api_token":   &me.APIToken,
		"project_key": &me.ProjectKey,
		"issue_type":  &me.IssueType,
		"summary":     &me.Summary,
		"description": &me.Description,
	})
}
