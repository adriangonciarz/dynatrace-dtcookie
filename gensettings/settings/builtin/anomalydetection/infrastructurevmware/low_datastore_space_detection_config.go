package infrastructurevmware

import "github.com/dtcookie/hcl"

type LowDatastoreSpaceDetectionConfig struct {
	CustomThresholds *LowDatastoreSpaceDetectionThresholds `json:"customThresholds"` // Alert if the condition is met in 1 out of 5 samples
	DetectionMode    DetectionMode                         `json:"detectionMode"`    // Detection mode
	Enabled          bool                                  `json:"enabled"`          // Detect low datastore space
}

func (me *LowDatastoreSpaceDetectionConfig) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"custom_thresholds": {
			Type:        hcl.TypeList,
			Description: "Alert if the condition is met in 1 out of 5 samples",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(LowDatastoreSpaceDetectionThresholds).Schema()},
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
			Description: "Detect low datastore space",
			Optional:    true,
		},
	}
}

func (me *LowDatastoreSpaceDetectionConfig) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"custom_thresholds": me.CustomThresholds,
		"detection_mode":    me.DetectionMode,
		"enabled":           me.Enabled,
	})
}

func (me *LowDatastoreSpaceDetectionConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"custom_thresholds": &me.CustomThresholds,
		"detection_mode":    &me.DetectionMode,
		"enabled":           &me.Enabled,
	})
}
