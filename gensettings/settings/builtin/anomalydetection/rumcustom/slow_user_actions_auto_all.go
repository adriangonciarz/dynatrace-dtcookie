package rumcustom

import "github.com/dtcookie/hcl"

type SlowUserActionsAutoAll struct {
	DurationThreshold  float64 `json:"durationThreshold"`  // Absolute threshold
	SlowdownPercentage float64 `json:"slowdownPercentage"` // Relative threshold
}

func (me *SlowUserActionsAutoAll) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"duration_threshold": {
			Type:        hcl.TypeFloat,
			Description: "Absolute threshold",
			Required:    true,
		},
		"slowdown_percentage": {
			Type:        hcl.TypeFloat,
			Description: "Relative threshold",
			Required:    true,
		},
	}
}

func (me *SlowUserActionsAutoAll) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"duration_threshold":  me.DurationThreshold,
		"slowdown_percentage": me.SlowdownPercentage,
	})
}

func (me *SlowUserActionsAutoAll) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"duration_threshold":  &me.DurationThreshold,
		"slowdown_percentage": &me.SlowdownPercentage,
	})
}
