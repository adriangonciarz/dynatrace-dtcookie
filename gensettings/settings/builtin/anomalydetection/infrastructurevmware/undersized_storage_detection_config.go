package infrastructurevmware

import "github.com/dtcookie/hcl"

type UndersizedStorageDetectionConfig struct {
	CustomThresholds *UndersizedStorageDetectionThresholds `json:"customThresholds"` // Alert if **any** condition is met in 3 out of 5 samples
	DetectionMode    DetectionMode                         `json:"detectionMode"`    // Detection mode
	Enabled          bool                                  `json:"enabled"`          // Detect undersized storage device
}

func (me *UndersizedStorageDetectionConfig) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"custom_thresholds": {
			Type:        hcl.TypeList,
			Description: "Alert if **any** condition is met in 3 out of 5 samples",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(UndersizedStorageDetectionThresholds).Schema()},
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
			Description: "Detect undersized storage device",
			Optional:    true,
		},
	}
}

func (me *UndersizedStorageDetectionConfig) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"custom_thresholds": me.CustomThresholds,
		"detection_mode":    me.DetectionMode,
		"enabled":           me.Enabled,
	})
}

func (me *UndersizedStorageDetectionConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"custom_thresholds": &me.CustomThresholds,
		"detection_mode":    &me.DetectionMode,
		"enabled":           &me.Enabled,
	})
}
