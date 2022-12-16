package php

import "github.com/dtcookie/hcl"

type Settings struct {
	Enabled        bool `json:"enabled"`        // Monitor PHP
	EnabledFastCGI bool `json:"enabledFastCGI"` // Requires PHP monitoring enabled and from Dynatrace OneAgent version 1.191 it's ignored and permanently enabled
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Monitor PHP",
			Optional:    true,
		},
		"enabled_fast_cgi": {
			Type:        hcl.TypeBool,
			Description: "Requires PHP monitoring enabled and from Dynatrace OneAgent version 1.191 it's ignored and permanently enabled",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enabled":          me.Enabled,
		"enabled_fast_cgi": me.EnabledFastCGI,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":          &me.Enabled,
		"enabled_fast_cgi": &me.EnabledFastCGI,
	})
}
