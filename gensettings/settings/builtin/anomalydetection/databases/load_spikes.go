package databases

import "github.com/dtcookie/hcl"

type LoadSpikes struct {
	Enabled              bool    `json:"enabled"`              // Detect service load spikes
	LoadSpikePercent     float64 `json:"loadSpikePercent"`     // Threshold
	MinutesAbnormalState int     `json:"minutesAbnormalState"` // Time span
}

func (me *LoadSpikes) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect service load spikes",
			Optional:    true,
		},
		"load_spike_percent": {
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

func (me *LoadSpikes) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enabled":                me.Enabled,
		"load_spike_percent":     me.LoadSpikePercent,
		"minutes_abnormal_state": me.MinutesAbnormalState,
	})
}

func (me *LoadSpikes) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":                &me.Enabled,
		"load_spike_percent":     &me.LoadSpikePercent,
		"minutes_abnormal_state": &me.MinutesAbnormalState,
	})
}
