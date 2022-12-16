package rumweb

import "github.com/dtcookie/hcl"

type ErrorRate struct {
	Enabled                bool            `json:"enabled"`                // Detect increases in JavaScript errors
	ErrorRateAuto          *ErrorRateAuto  `json:"errorRateAuto"`          // Alert if the percentage of failing user actions increases by **both** the absolute and relative thresholds:
	ErrorRateDetectionMode DetectionMode   `json:"errorRateDetectionMode"` // Detection strategy for increases in JavaScript errors
	ErrorRateFixed         *ErrorRateFixed `json:"errorRateFixed"`
}

func (me *ErrorRate) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect increases in JavaScript errors",
			Optional:    true,
		},
		"error_rate_auto": {
			Type:        hcl.TypeList,
			Description: "Alert if the percentage of failing user actions increases by **both** the absolute and relative thresholds:",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ErrorRateAuto).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"error_rate_detection_mode": {
			Type:        hcl.TypeString,
			Description: "Detection strategy for increases in JavaScript errors",
			Required:    true,
		},
		"error_rate_fixed": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ErrorRateFixed).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *ErrorRate) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enabled":                   me.Enabled,
		"error_rate_auto":           me.ErrorRateAuto,
		"error_rate_detection_mode": me.ErrorRateDetectionMode,
		"error_rate_fixed":          me.ErrorRateFixed,
	})
}

func (me *ErrorRate) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":                   &me.Enabled,
		"error_rate_auto":           &me.ErrorRateAuto,
		"error_rate_detection_mode": &me.ErrorRateDetectionMode,
		"error_rate_fixed":          &me.ErrorRateFixed,
	})
}
