package rumcustom

import "github.com/dtcookie/hcl"

type ErrorRateIncreaseFixed struct {
	Sensitivity       Sensitivity `json:"sensitivity"`       // Detection sensitivity
	ThresholdAbsolute float64     `json:"thresholdAbsolute"` // Absolute threshold
}

func (me *ErrorRateIncreaseFixed) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"sensitivity": {
			Type:        hcl.TypeString,
			Description: "Detection sensitivity",
			Required:    true,
		},
		"threshold_absolute": {
			Type:        hcl.TypeFloat,
			Description: "Absolute threshold",
			Required:    true,
		},
	}
}

func (me *ErrorRateIncreaseFixed) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"sensitivity":        me.Sensitivity,
		"threshold_absolute": me.ThresholdAbsolute,
	})
}

func (me *ErrorRateIncreaseFixed) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"sensitivity":        &me.Sensitivity,
		"threshold_absolute": &me.ThresholdAbsolute,
	})
}
