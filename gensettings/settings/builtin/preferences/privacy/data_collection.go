package privacy

import "github.com/dtcookie/hcl"

type DataCollection struct {
	OptInModeEnabled bool `json:"optInModeEnabled"` // With [Data-collection and opt-in mode](https://dt-url.net/7l3p0p3h) enabled, Real User Monitoring data isn't captured until dtrum.enable() is called for specific user sessions.
}

func (me *DataCollection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"opt_in_mode_enabled": {
			Type:        hcl.TypeBool,
			Description: "With [Data-collection and opt-in mode](https://dt-url.net/7l3p0p3h) enabled, Real User Monitoring data isn't captured until dtrum.enable() is called for specific user sessions.",
			Optional:    true,
		},
	}
}

func (me *DataCollection) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"opt_in_mode_enabled": me.OptInModeEnabled,
	})
}

func (me *DataCollection) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"opt_in_mode_enabled": &me.OptInModeEnabled,
	})
}
