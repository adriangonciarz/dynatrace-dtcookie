package infrastructurehosts

import "github.com/dtcookie/hcl"

type HighCpuSaturationDetection struct {
	CustomThresholds *HighCpuSaturationDetectionThresholds `json:"customThresholds"`
	DetectionMode    DetectionMode                         `json:"detectionMode"` // Detection mode for CPU saturation
	Enabled          bool                                  `json:"enabled"`       // Detect CPU saturation on host
}

func (me *HighCpuSaturationDetection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"custom_thresholds": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(HighCpuSaturationDetectionThresholds).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"detection_mode": {
			Type:        hcl.TypeString,
			Description: "Detection mode for CPU saturation",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect CPU saturation on host",
			Optional:    true,
		},
	}
}

func (me *HighCpuSaturationDetection) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"custom_thresholds": me.CustomThresholds,
		"detection_mode":    me.DetectionMode,
		"enabled":           me.Enabled,
	})
}

func (me *HighCpuSaturationDetection) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"custom_thresholds": &me.CustomThresholds,
		"detection_mode":    &me.DetectionMode,
		"enabled":           &me.Enabled,
	})
}
