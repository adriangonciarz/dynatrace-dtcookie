package infrastructurevmware

import "github.com/dtcookie/hcl"

type EsxiHighCpuDetectionConfig struct {
	CustomThresholds *EsxiHighCpuDetectionThresholds `json:"customThresholds"` // Alert if **all three** conditions are met in 3 out of 5 samples
	DetectionMode    DetectionMode                   `json:"detectionMode"`    // Detection mode
	Enabled          bool                            `json:"enabled"`          // Detect high CPU saturation on ESXi host
}

func (me *EsxiHighCpuDetectionConfig) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"custom_thresholds": {
			Type:        hcl.TypeList,
			Description: "Alert if **all three** conditions are met in 3 out of 5 samples",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(EsxiHighCpuDetectionThresholds).Schema()},
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
			Description: "Detect high CPU saturation on ESXi host",
			Optional:    true,
		},
	}
}

func (me *EsxiHighCpuDetectionConfig) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"custom_thresholds": me.CustomThresholds,
		"detection_mode":    me.DetectionMode,
		"enabled":           me.Enabled,
	})
}

func (me *EsxiHighCpuDetectionConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"custom_thresholds": &me.CustomThresholds,
		"detection_mode":    &me.DetectionMode,
		"enabled":           &me.Enabled,
	})
}
