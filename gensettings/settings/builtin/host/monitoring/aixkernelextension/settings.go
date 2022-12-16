package aixkernelextension

import "github.com/dtcookie/hcl"

type Settings struct {
	Enabled           bool `json:"enabled"`           // Allow AIX kernel extension
	UseGlobalSettings bool `json:"useGlobalSettings"` // Use global settings
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Allow AIX kernel extension",
			Optional:    true,
		},
		"use_global_settings": {
			Type:        hcl.TypeBool,
			Description: "Use global settings",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enabled":             me.Enabled,
		"use_global_settings": me.UseGlobalSettings,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":             &me.Enabled,
		"use_global_settings": &me.UseGlobalSettings,
	})
}
