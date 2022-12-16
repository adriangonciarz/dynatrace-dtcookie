package infrastructurevmware

import "github.com/dtcookie/hcl"

type SlowPhysicalStorageDetectionConfig struct {
	CustomThresholds *SlowPhysicalStorageDetectionThresholds `json:"customThresholds"` // Alert if **any** condition is met in 4 out of 5 samples
	DetectionMode    DetectionMode                           `json:"detectionMode"`    // Detection mode
	Enabled          bool                                    `json:"enabled"`          // Detect physical storage device running slow
}

func (me *SlowPhysicalStorageDetectionConfig) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"custom_thresholds": {
			Type:        hcl.TypeList,
			Description: "Alert if **any** condition is met in 4 out of 5 samples",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(SlowPhysicalStorageDetectionThresholds).Schema()},
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
			Description: "Detect physical storage device running slow",
			Optional:    true,
		},
	}
}

func (me *SlowPhysicalStorageDetectionConfig) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"custom_thresholds": me.CustomThresholds,
		"detection_mode":    me.DetectionMode,
		"enabled":           me.Enabled,
	})
}

func (me *SlowPhysicalStorageDetectionConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"custom_thresholds": &me.CustomThresholds,
		"detection_mode":    &me.DetectionMode,
		"enabled":           &me.Enabled,
	})
}
