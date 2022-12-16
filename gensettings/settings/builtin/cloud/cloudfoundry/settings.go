package cloudfoundry

import "github.com/dtcookie/hcl"

type Settings struct {
	ActiveGateGroup *string `json:"activeGateGroup"` // ActiveGate group
	ApiUrl          string  `json:"apiUrl"`          // Cloud Foundry API Target
	Enabled         bool    `json:"enabled"`         // Enabled
	Label           string  `json:"label"`           // Name this connection
	LoginUrl        string  `json:"loginUrl"`        // Cloud Foundry Authentication Endpoint
	Password        string  `json:"password"`        // Cloud Foundry Password
	Username        string  `json:"username"`        // Cloud Foundry Username
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"active_gate_group": {
			Type:        hcl.TypeString,
			Description: "ActiveGate group",
			Optional:    true,
		},
		"api_url": {
			Type:        hcl.TypeString,
			Description: "Cloud Foundry API Target",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enabled",
			Optional:    true,
		},
		"label": {
			Type:        hcl.TypeString,
			Description: "Name this connection",
			Required:    true,
		},
		"login_url": {
			Type:        hcl.TypeString,
			Description: "Cloud Foundry Authentication Endpoint",
			Required:    true,
		},
		"password": {
			Type:        hcl.TypeString,
			Description: "Cloud Foundry Password",
			Required:    true,
		},
		"username": {
			Type:        hcl.TypeString,
			Description: "Cloud Foundry Username",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"active_gate_group": me.ActiveGateGroup,
		"api_url":           me.ApiUrl,
		"enabled":           me.Enabled,
		"label":             me.Label,
		"login_url":         me.LoginUrl,
		"password":          me.Password,
		"username":          me.Username,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"active_gate_group": &me.ActiveGateGroup,
		"api_url":           &me.ApiUrl,
		"enabled":           &me.Enabled,
		"label":             &me.Label,
		"login_url":         &me.LoginUrl,
		"password":          &me.Password,
		"username":          &me.Username,
	})
}
