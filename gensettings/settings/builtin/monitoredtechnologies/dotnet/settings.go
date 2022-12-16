package dotnet

import "github.com/dtcookie/hcl"

type Settings struct {
	Enabled           bool `json:"enabled"`           // Requires Dynatrace OneAgent version 1.75 or later on Windows
	EnabledDotNetCore bool `json:"enabledDotNetCore"` // Requires Dynatrace OneAgent version 1.117 or later on Windows and version 1.127 or later on Linux and .NET monitoring enabled
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Requires Dynatrace OneAgent version 1.75 or later on Windows",
			Optional:    true,
		},
		"enabled_dot_net_core": {
			Type:        hcl.TypeBool,
			Description: "Requires Dynatrace OneAgent version 1.117 or later on Windows and version 1.127 or later on Linux and .NET monitoring enabled",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enabled":              me.Enabled,
		"enabled_dot_net_core": me.EnabledDotNetCore,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":              &me.Enabled,
		"enabled_dot_net_core": &me.EnabledDotNetCore,
	})
}
