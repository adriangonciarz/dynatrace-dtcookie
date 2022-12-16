package privacy

import "github.com/dtcookie/hcl"

type Settings struct {
	OptInModeEnabled bool `json:"optInModeEnabled"` // Enable user opt-in mode
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"opt_in_mode_enabled": {
			Type:        hcl.TypeBool,
			Description: "Enable user opt-in mode",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"opt_in_mode_enabled": me.OptInModeEnabled,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"opt_in_mode_enabled": &me.OptInModeEnabled,
	})
}
