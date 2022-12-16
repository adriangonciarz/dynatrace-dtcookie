package rummobile

import "github.com/dtcookie/hcl"

type UnexpectedLowLoad struct {
	Enabled             bool    `json:"enabled"`             // Detect unexpected low load
	ThresholdPercentage float64 `json:"thresholdPercentage"` // Dynatrace learns your typical application traffic over an observation period of one week. Depending on this expected value Dynatrace detects abnormal traffic drops within your application.
}

func (me *UnexpectedLowLoad) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect unexpected low load",
			Optional:    true,
		},
		"threshold_percentage": {
			Type:        hcl.TypeFloat,
			Description: "Dynatrace learns your typical application traffic over an observation period of one week. Depending on this expected value Dynatrace detects abnormal traffic drops within your application.",
			Required:    true,
		},
	}
}

func (me *UnexpectedLowLoad) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enabled":              me.Enabled,
		"threshold_percentage": me.ThresholdPercentage,
	})
}

func (me *UnexpectedLowLoad) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":              &me.Enabled,
		"threshold_percentage": &me.ThresholdPercentage,
	})
}
