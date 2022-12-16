package golang

import "github.com/dtcookie/hcl"

type Settings struct {
	Enabled                   bool `json:"enabled"`                   // Monitor Go
	EnabledGoStaticMonitoring bool `json:"enabledGoStaticMonitoring"` // Learn more about the [known limitations for Go static monitoring](https://www.dynatrace.com/support/help/technology-support/application-software/go/support/go-known-limitations#limitations)
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Monitor Go",
			Optional:    true,
		},
		"enabled_go_static_monitoring": {
			Type:        hcl.TypeBool,
			Description: "Learn more about the [known limitations for Go static monitoring](https://www.dynatrace.com/support/help/technology-support/application-software/go/support/go-known-limitations#limitations)",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enabled":                      me.Enabled,
		"enabled_go_static_monitoring": me.EnabledGoStaticMonitoring,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":                      &me.Enabled,
		"enabled_go_static_monitoring": &me.EnabledGoStaticMonitoring,
	})
}
