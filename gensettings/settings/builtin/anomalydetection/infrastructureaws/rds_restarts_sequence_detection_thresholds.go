package infrastructureaws

import "github.com/dtcookie/hcl"

// RdsRestartsSequenceDetectionThresholds. Alert if the condition is met in 2 out of 20 samples
type RdsRestartsSequenceDetectionThresholds struct {
	RestartsPerMinute int `json:"restartsPerMinute"` // Number of restarts per minute is equal or higher than
}

func (me *RdsRestartsSequenceDetectionThresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"restarts_per_minute": {
			Type:        hcl.TypeInt,
			Description: "Number of restarts per minute is equal or higher than",
			Required:    true,
		},
	}
}

func (me *RdsRestartsSequenceDetectionThresholds) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"restarts_per_minute": me.RestartsPerMinute,
	})
}

func (me *RdsRestartsSequenceDetectionThresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"restarts_per_minute": &me.RestartsPerMinute,
	})
}
