package perdiskoverride

import "github.com/dtcookie/hcl"

type DiskLowInodesDetection struct {
	CustomThresholds *DiskLowInodesDetectionThresholds `json:"customThresholds"`
	DetectionMode    DetectionMode                     `json:"detectionMode"` // Detection mode for low inodes number available
	Enabled          bool                              `json:"enabled"`       // Detect low inodes number available
}

func (me *DiskLowInodesDetection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"custom_thresholds": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(DiskLowInodesDetectionThresholds).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"detection_mode": {
			Type:        hcl.TypeString,
			Description: "Detection mode for low inodes number available",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect low inodes number available",
			Optional:    true,
		},
	}
}

func (me *DiskLowInodesDetection) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"custom_thresholds": me.CustomThresholds,
		"detection_mode":    me.DetectionMode,
		"enabled":           me.Enabled,
	})
}

func (me *DiskLowInodesDetection) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"custom_thresholds": &me.CustomThresholds,
		"detection_mode":    &me.DetectionMode,
		"enabled":           &me.Enabled,
	})
}
