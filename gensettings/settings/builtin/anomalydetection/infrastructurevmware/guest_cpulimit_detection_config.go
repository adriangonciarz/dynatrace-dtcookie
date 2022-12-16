package infrastructurevmware

import "github.com/dtcookie/hcl"

type GuestCPULimitDetectionConfig struct {
	CustomThresholds *GuestCPULimitDetectionThresholds `json:"customThresholds"` // Alert if **all three** conditions are met in 3 out of 5 samples
	DetectionMode    DetectionMode                     `json:"detectionMode"`    // Detection mode
	Enabled          bool                              `json:"enabled"`          // Detect guest CPU limit reached
}

func (me *GuestCPULimitDetectionConfig) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"custom_thresholds": {
			Type:        hcl.TypeList,
			Description: "Alert if **all three** conditions are met in 3 out of 5 samples",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(GuestCPULimitDetectionThresholds).Schema()},
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
			Description: "Detect guest CPU limit reached",
			Optional:    true,
		},
	}
}

func (me *GuestCPULimitDetectionConfig) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"custom_thresholds": me.CustomThresholds,
		"detection_mode":    me.DetectionMode,
		"enabled":           me.Enabled,
	})
}

func (me *GuestCPULimitDetectionConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"custom_thresholds": &me.CustomThresholds,
		"detection_mode":    &me.DetectionMode,
		"enabled":           &me.Enabled,
	})
}
