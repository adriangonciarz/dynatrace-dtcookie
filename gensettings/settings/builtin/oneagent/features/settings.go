package features

import "github.com/dtcookie/hcl"

type Settings struct {
	Enabled         bool   `json:"enabled"`         // Enabled
	Forcible        *bool  `json:"forcible"`        // Activate this feature also in OneAgents only fulfilling the minimum Opt-In version
	Instrumentation *bool  `json:"instrumentation"` // Instrumentation enabled (change needs a process restart)
	Key             string `json:"key"`             // Feature
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Enabled",
			Optional:    true,
		},
		"forcible": {
			Type:        hcl.TypeBool,
			Description: "Activate this feature also in OneAgents only fulfilling the minimum Opt-In version",
			Optional:    true,
		},
		"instrumentation": {
			Type:        hcl.TypeBool,
			Description: "Instrumentation enabled (change needs a process restart)",
			Optional:    true,
		},
		"key": {
			Type:        hcl.TypeString,
			Description: "Feature",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enabled":         me.Enabled,
		"forcible":        me.Forcible,
		"instrumentation": me.Instrumentation,
		"key":             me.Key,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":         &me.Enabled,
		"forcible":        &me.Forcible,
		"instrumentation": &me.Instrumentation,
		"key":             &me.Key,
	})
}
