package infrastructureaws

import "github.com/dtcookie/hcl"

type RdsHighCpuDetectionConfig struct {
	CustomThresholds *RdsHighCpuDetectionThresholds `json:"customThresholds"` // Alert if the condition is met in 3 out of 5 samples
	DetectionMode    DetectionMode                  `json:"detectionMode"`    // Detection mode
	Enabled          bool                           `json:"enabled"`          // Detect high CPU utilization on RDS
}

func (me *RdsHighCpuDetectionConfig) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"custom_thresholds": {
			Type:        hcl.TypeList,
			Description: "Alert if the condition is met in 3 out of 5 samples",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(RdsHighCpuDetectionThresholds).Schema()},
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
			Description: "Detect high CPU utilization on RDS",
			Optional:    true,
		},
	}
}

func (me *RdsHighCpuDetectionConfig) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"custom_thresholds": me.CustomThresholds,
		"detection_mode":    me.DetectionMode,
		"enabled":           me.Enabled,
	})
}

func (me *RdsHighCpuDetectionConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"custom_thresholds": &me.CustomThresholds,
		"detection_mode":    &me.DetectionMode,
		"enabled":           &me.Enabled,
	})
}
