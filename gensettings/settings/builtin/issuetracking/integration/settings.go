package integration

import "github.com/dtcookie/hcl"

type Settings struct {
	Enabled            bool               `json:"enabled"`            // Enabled
	Issuelabel         string             `json:"issuelabel"`         // Set a label to identify these issues, for example, `release_blocker` or `non-critical`
	Issuequery         string             `json:"issuequery"`         // You can use the following placeholders to automatically insert values from the **Release monitoring** page in your query: `{NAME}`, `{VERSION}`, `{STAGE}`, `{PRODUCT}`.
	Issuetheme         IssueTheme         `json:"issuetheme"`         // Select the issue type to be displayed.
	Issuetrackersystem IssueTrackerSystem `json:"issuetrackersystem"` // Select the issue-tracking system you want to query.
	Password           *string            `json:"password"`           // Password
	Token              *string            `json:"token"`              // Token
	Url                string             `json:"url"`                // For Jira, use the base URL (for example, https://jira.yourcompany.com); for GitHub, use the repository URL (for example, https://github.com/org/repo); for GitLab, use the specific project API for a single project (for example, https://gitlab.com/api/v4/projects/:projectId), and the specific group API for a multiple projects (for example, https://gitlab.com/api/v4/groups/:groupId); for ServiceNow, use your company instance URL (for example, https://yourinstance.service-now.com/)
	Username           string             `json:"username"`           // Username
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enabled",
			Optional:    true,
		},
		"issuelabel": {
			Type:        hcl.TypeString,
			Description: "Set a label to identify these issues, for example, `release_blocker` or `non-critical`",
			Required:    true,
		},
		"issuequery": {
			Type:        hcl.TypeString,
			Description: "You can use the following placeholders to automatically insert values from the **Release monitoring** page in your query: `{NAME}`, `{VERSION}`, `{STAGE}`, `{PRODUCT}`.",
			Required:    true,
		},
		"issuetheme": {
			Type:        hcl.TypeString,
			Description: "Select the issue type to be displayed.",
			Required:    true,
		},
		"issuetrackersystem": {
			Type:        hcl.TypeString,
			Description: "Select the issue-tracking system you want to query.",
			Required:    true,
		},
		"password": {
			Type:        hcl.TypeString,
			Description: "Password",
			Optional:    true,
		},
		"token": {
			Type:        hcl.TypeString,
			Description: "Token",
			Optional:    true,
		},
		"url": {
			Type:        hcl.TypeString,
			Description: "For Jira, use the base URL (for example, https://jira.yourcompany.com); for GitHub, use the repository URL (for example, https://github.com/org/repo); for GitLab, use the specific project API for a single project (for example, https://gitlab.com/api/v4/projects/:projectId), and the specific group API for a multiple projects (for example, https://gitlab.com/api/v4/groups/:groupId); for ServiceNow, use your company instance URL (for example, https://yourinstance.service-now.com/)",
			Required:    true,
		},
		"username": {
			Type:        hcl.TypeString,
			Description: "Username",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enabled":            me.Enabled,
		"issuelabel":         me.Issuelabel,
		"issuequery":         me.Issuequery,
		"issuetheme":         me.Issuetheme,
		"issuetrackersystem": me.Issuetrackersystem,
		"password":           me.Password,
		"token":              me.Token,
		"url":                me.Url,
		"username":           me.Username,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":            &me.Enabled,
		"issuelabel":         &me.Issuelabel,
		"issuequery":         &me.Issuequery,
		"issuetheme":         &me.Issuetheme,
		"issuetrackersystem": &me.Issuetrackersystem,
		"password":           &me.Password,
		"token":              &me.Token,
		"url":                &me.Url,
		"username":           &me.Username,
	})
}
