package databases

import "github.com/dtcookie/hcl"

type LoadDrops struct {
	Enabled              bool    `json:"enabled"`              // Detect service load drops
	LoadDropPercent      float64 `json:"loadDropPercent"`      // Threshold
	MinutesAbnormalState int     `json:"minutesAbnormalState"` // Time span
}

func (me *LoadDrops) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect service load drops",
			Optional:    true,
		},
		"load_drop_percent": {
			Type:        hcl.TypeFloat,
			Description: "Threshold",
			Required:    true,
		},
		"minutes_abnormal_state": {
			Type:        hcl.TypeInt,
			Description: "Time span",
			Required:    true,
		},
	}
}

func (me *LoadDrops) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enabled":                me.Enabled,
		"load_drop_percent":      me.LoadDropPercent,
		"minutes_abnormal_state": me.MinutesAbnormalState,
	})
}

func (me *LoadDrops) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":                &me.Enabled,
		"load_drop_percent":      &me.LoadDropPercent,
		"minutes_abnormal_state": &me.MinutesAbnormalState,
	})
}
