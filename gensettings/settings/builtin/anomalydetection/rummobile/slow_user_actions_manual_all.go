package rummobile

import "github.com/dtcookie/hcl"

type SlowUserActionsManualAll struct {
	DurationThreshold float64 `json:"durationThreshold"` // Absolute threshold
}

func (me *SlowUserActionsManualAll) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"duration_threshold": {
			Type:        hcl.TypeFloat,
			Description: "Absolute threshold",
			Required:    true,
		},
	}
}

func (me *SlowUserActionsManualAll) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"duration_threshold": me.DurationThreshold,
	})
}

func (me *SlowUserActionsManualAll) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"duration_threshold": &me.DurationThreshold,
	})
}
