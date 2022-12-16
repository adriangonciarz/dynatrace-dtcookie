package infrastructureaws

import "github.com/dtcookie/hcl"

type RdsHighWriteReadLatencyDetectionConfig struct {
	CustomThresholds *RdsHighWriteReadLatencyDetectionThresholds `json:"customThresholds"` // Alert if the condition is met in 3 out of 5 samples
	DetectionMode    DetectionMode                               `json:"detectionMode"`    // Detection mode
	Enabled          bool                                        `json:"enabled"`          // Detect high RDS write/read latency
}

func (me *RdsHighWriteReadLatencyDetectionConfig) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"custom_thresholds": {
			Type:        hcl.TypeList,
			Description: "Alert if the condition is met in 3 out of 5 samples",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(RdsHighWriteReadLatencyDetectionThresholds).Schema()},
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
			Description: "Detect high RDS write/read latency",
			Optional:    true,
		},
	}
}

func (me *RdsHighWriteReadLatencyDetectionConfig) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"custom_thresholds": me.CustomThresholds,
		"detection_mode":    me.DetectionMode,
		"enabled":           me.Enabled,
	})
}

func (me *RdsHighWriteReadLatencyDetectionConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"custom_thresholds": &me.CustomThresholds,
		"detection_mode":    &me.DetectionMode,
		"enabled":           &me.Enabled,
	})
}
