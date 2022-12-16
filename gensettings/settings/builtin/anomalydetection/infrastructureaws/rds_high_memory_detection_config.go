package infrastructureaws

import "github.com/dtcookie/hcl"

type RdsHighMemoryDetectionConfig struct {
	CustomThresholds *RdsHighMemoryDetectionThresholds `json:"customThresholds"` // Alert if **both** conditions is met in 3 out of 5 samples
	DetectionMode    DetectionMode                     `json:"detectionMode"`    // Detection mode
	Enabled          bool                              `json:"enabled"`          // Detect RDS running out of memory
}

func (me *RdsHighMemoryDetectionConfig) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"custom_thresholds": {
			Type:        hcl.TypeList,
			Description: "Alert if **both** conditions is met in 3 out of 5 samples",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(RdsHighMemoryDetectionThresholds).Schema()},
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
			Description: "Detect RDS running out of memory",
			Optional:    true,
		},
	}
}

func (me *RdsHighMemoryDetectionConfig) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"custom_thresholds": me.CustomThresholds,
		"detection_mode":    me.DetectionMode,
		"enabled":           me.Enabled,
	})
}

func (me *RdsHighMemoryDetectionConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"custom_thresholds": &me.CustomThresholds,
		"detection_mode":    &me.DetectionMode,
		"enabled":           &me.Enabled,
	})
}
