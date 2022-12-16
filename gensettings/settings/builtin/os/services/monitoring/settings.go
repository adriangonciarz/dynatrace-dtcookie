package monitoring

import "github.com/dtcookie/hcl"

type Settings struct {
	DisplayName string `json:"displayName"` // Display name
	Enabled     bool   `json:"enabled"`     // Enabled
	ServiceName string `json:"serviceName"` // Service name
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"display_name": {
			Type:        hcl.TypeString,
			Description: "Display name",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enabled",
			Optional:    true,
		},
		"service_name": {
			Type:        hcl.TypeString,
			Description: "Service name",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"display_name": me.DisplayName,
		"enabled":      me.Enabled,
		"service_name": me.ServiceName,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"display_name": &me.DisplayName,
		"enabled":      &me.Enabled,
		"service_name": &me.ServiceName,
	})
}
