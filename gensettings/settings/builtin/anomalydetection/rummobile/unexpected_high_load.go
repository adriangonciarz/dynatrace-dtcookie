package rummobile

import "github.com/dtcookie/hcl"

type UnexpectedHighLoad struct {
	Enabled             bool    `json:"enabled"`             // Detect unexpected high load
	ThresholdPercentage float64 `json:"thresholdPercentage"` // Dynatrace learns your typical application traffic over an observation period of one week. Depending on this expected value Dynatrace detects abnormal traffic spikes within your application.
}

func (me *UnexpectedHighLoad) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect unexpected high load",
			Optional:    true,
		},
		"threshold_percentage": {
			Type:        hcl.TypeFloat,
			Description: "Dynatrace learns your typical application traffic over an observation period of one week. Depending on this expected value Dynatrace detects abnormal traffic spikes within your application.",
			Required:    true,
		},
	}
}

func (me *UnexpectedHighLoad) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enabled":              me.Enabled,
		"threshold_percentage": me.ThresholdPercentage,
	})
}

func (me *UnexpectedHighLoad) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled":              &me.Enabled,
		"threshold_percentage": &me.ThresholdPercentage,
	})
}
