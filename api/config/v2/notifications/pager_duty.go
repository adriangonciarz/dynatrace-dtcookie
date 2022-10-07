package notifications

import "github.com/dtcookie/hcl"

type PagerDuty struct {
	Enabled   bool   `json:"-"`
	Name      string `json:"-"`
	ProfileID string `json:"-"`

	Account     string `json:"account"`       // The name of the PagerDuty account
	ServiceName string `json:"serviceName"`   // The name of the service
	APIKey      string `json:"serviceApiKey"` // The API key to access PagerDuty
}

func (me *PagerDuty) Schema() map[string]*hcl.Schema {
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

		"account": {
			Type:        hcl.TypeString,
			Description: "The name of the PagerDuty account",
			Required:    true,
		},
		"api_key": {
			Type:        hcl.TypeString,
			Sensitive:   true,
			Description: "The API key to access PagerDuty",
			Optional:    true,
		},
		"service": {
			Type:        hcl.TypeString,
			Description: "The name of the PagerDuty Service",
			Required:    true,
		},
	}
}

func (me *PagerDuty) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	// The api_key field MUST NOT get serialized into HCL here
	// The Dynatrace Settings 2.0 API delivers a scrambled version of any previously stored api_key here
	// Evaluation at this point would lead to that scrambled version to make it into the Terraform State
	// As a result any plans would be non-empty
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"name":    me.Name,
		"active":  me.Enabled,
		"profile": me.ProfileID,

		"account": me.Account,
		"service": me.ServiceName,
		// "api_key":     me.APIKey,
	})
}

func (me *PagerDuty) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"name":    &me.Name,
		"active":  &me.Enabled,
		"profile": &me.ProfileID,

		"account": &me.Account,
		"service": &me.ServiceName,
		"api_key": &me.APIKey,
	})
}
