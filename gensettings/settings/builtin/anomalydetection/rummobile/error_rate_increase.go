package rummobile

import "github.com/dtcookie/hcl"

type ErrorRateIncrease struct {
	DetectionMode          DetectionMode           `json:"detectionMode"`          // Detection strategy for error rate increases
	Enabled                bool                    `json:"enabled"`                // Detect reported error rate increase
	ErrorRateIncreaseAuto  *ErrorRateIncreaseAuto  `json:"errorRateIncreaseAuto"`  // Alert if the percentage of user actions affected by reported errors exceeds **both** the absolute threshold and the relative threshold
	ErrorRateIncreaseFixed *ErrorRateIncreaseFixed `json:"errorRateIncreaseFixed"` // Alert if the custom reported error rate threshold is exceeded during any 5-minute period
}

func (me *ErrorRateIncrease) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"detection_mode": {
			Type:        hcl.TypeString,
			Description: "Detection strategy for error rate increases",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect reported error rate increase",
			Optional:    true,
		},
		"error_rate_increase_auto": {
			Type:        hcl.TypeList,
			Description: "Alert if the percentage of user actions affected by reported errors exceeds **both** the absolute threshold and the relative threshold",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ErrorRateIncreaseAuto).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"error_rate_increase_fixed": {
			Type:        hcl.TypeList,
			Description: "Alert if the custom reported error rate threshold is exceeded during any 5-minute period",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ErrorRateIncreaseFixed).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *ErrorRateIncrease) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"detection_mode":            me.DetectionMode,
		"enabled":                   me.Enabled,
		"error_rate_increase_auto":  me.ErrorRateIncreaseAuto,
		"error_rate_increase_fixed": me.ErrorRateIncreaseFixed,
	})
}

func (me *ErrorRateIncrease) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"detection_mode":            &me.DetectionMode,
		"enabled":                   &me.Enabled,
		"error_rate_increase_auto":  &me.ErrorRateIncreaseAuto,
		"error_rate_increase_fixed": &me.ErrorRateIncreaseFixed,
	})
}
