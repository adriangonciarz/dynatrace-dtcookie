package rummobile

import "github.com/dtcookie/hcl"

type SlowUserActionsManualSlowest struct {
	DurationThreshold float64 `json:"durationThreshold"` // Absolute threshold
}

func (me *SlowUserActionsManualSlowest) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"duration_threshold": {
			Type:        hcl.TypeFloat,
			Description: "Absolute threshold",
			Required:    true,
		},
	}
}

func (me *SlowUserActionsManualSlowest) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"duration_threshold": me.DurationThreshold,
	})
}

func (me *SlowUserActionsManualSlowest) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"duration_threshold": &me.DurationThreshold,
	})
}
