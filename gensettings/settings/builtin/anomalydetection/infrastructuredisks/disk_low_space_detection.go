package infrastructuredisks

import "github.com/dtcookie/hcl"

type DiskLowSpaceDetection struct {
	CustomThresholds *DiskLowSpaceDetectionThresholds `json:"customThresholds"`
	DetectionMode    DetectionMode                    `json:"detectionMode"` // Detection mode for low disk space
	Enabled          bool                             `json:"enabled"`       // Detect low disk space
}

func (me *DiskLowSpaceDetection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"custom_thresholds": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(DiskLowSpaceDetectionThresholds).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"detection_mode": {
			Type:        hcl.TypeString,
			Description: "Detection mode for low disk space",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect low disk space",
			Optional:    true,
		},
	}
}

func (me *DiskLowSpaceDetection) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"custom_thresholds": me.CustomThresholds,
		"detection_mode":    me.DetectionMode,
		"enabled":           me.Enabled,
	})
}

func (me *DiskLowSpaceDetection) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"custom_thresholds": &me.CustomThresholds,
		"detection_mode":    &me.DetectionMode,
		"enabled":           &me.Enabled,
	})
}
