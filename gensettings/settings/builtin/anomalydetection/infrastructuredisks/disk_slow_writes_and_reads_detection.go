package infrastructuredisks

import "github.com/dtcookie/hcl"

type DiskSlowWritesAndReadsDetection struct {
	CustomThresholds *DiskSlowWritesAndReadsDetectionThresholds `json:"customThresholds"`
	DetectionMode    DetectionMode                              `json:"detectionMode"` // Detection mode for slow running disks
	Enabled          bool                                       `json:"enabled"`       // Detect slow-running disks
}

func (me *DiskSlowWritesAndReadsDetection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"custom_thresholds": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(DiskSlowWritesAndReadsDetectionThresholds).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"detection_mode": {
			Type:        hcl.TypeString,
			Description: "Detection mode for slow running disks",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect slow-running disks",
			Optional:    true,
		},
	}
}

func (me *DiskSlowWritesAndReadsDetection) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"custom_thresholds": me.CustomThresholds,
		"detection_mode":    me.DetectionMode,
		"enabled":           me.Enabled,
	})
}

func (me *DiskSlowWritesAndReadsDetection) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"custom_thresholds": &me.CustomThresholds,
		"detection_mode":    &me.DetectionMode,
		"enabled":           &me.Enabled,
	})
}
