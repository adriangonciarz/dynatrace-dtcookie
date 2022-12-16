package rumcustom

import "github.com/dtcookie/hcl"

type ErrorRateIncreaseAuto struct {
	ThresholdAbsolute float64 `json:"thresholdAbsolute"` // Absolute threshold
	ThresholdRelative float64 `json:"thresholdRelative"` // Relative threshold
}

func (me *ErrorRateIncreaseAuto) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"threshold_absolute": {
			Type:        hcl.TypeFloat,
			Description: "Absolute threshold",
			Required:    true,
		},
		"threshold_relative": {
			Type:        hcl.TypeFloat,
			Description: "Relative threshold",
			Required:    true,
		},
	}
}

func (me *ErrorRateIncreaseAuto) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"threshold_absolute": me.ThresholdAbsolute,
		"threshold_relative": me.ThresholdRelative,
	})
}

func (me *ErrorRateIncreaseAuto) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"threshold_absolute": &me.ThresholdAbsolute,
		"threshold_relative": &me.ThresholdRelative,
	})
}
