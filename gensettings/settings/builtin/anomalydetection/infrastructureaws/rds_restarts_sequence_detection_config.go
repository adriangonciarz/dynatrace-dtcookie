package infrastructureaws

import "github.com/dtcookie/hcl"

type RdsRestartsSequenceDetectionConfig struct {
	CustomThresholds *RdsRestartsSequenceDetectionThresholds `json:"customThresholds"` // Alert if the condition is met in 2 out of 20 samples
	DetectionMode    DetectionMode                           `json:"detectionMode"`    // Detection mode
	Enabled          bool                                    `json:"enabled"`          // Detect restarts sequence on RDS
}

func (me *RdsRestartsSequenceDetectionConfig) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"custom_thresholds": {
			Type:        hcl.TypeList,
			Description: "Alert if the condition is met in 2 out of 20 samples",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(RdsRestartsSequenceDetectionThresholds).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"detection_mode": {
			Type:        hcl.TypeString,
			Description: "Detection mode",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect restarts sequence on RDS",
			Optional:    true,
		},
	}
}

func (me *RdsRestartsSequenceDetectionConfig) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"custom_thresholds": me.CustomThresholds,
		"detection_mode":    me.DetectionMode,
		"enabled":           me.Enabled,
	})
}

func (me *RdsRestartsSequenceDetectionConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"custom_thresholds": &me.CustomThresholds,
		"detection_mode":    &me.DetectionMode,
		"enabled":           &me.Enabled,
	})
}
