package infrastructurehosts

import "github.com/dtcookie/hcl"

type OutOfMemoryDetection struct {
	CustomThresholds *OutOfMemoryDetectionThresholds `json:"customThresholds"`
	DetectionMode    DetectionMode                   `json:"detectionMode"` // Detection mode for Java out of memory problem
	Enabled          bool                            `json:"enabled"`       // Detect Java out of memory problem
}

func (me *OutOfMemoryDetection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"custom_thresholds": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(OutOfMemoryDetectionThresholds).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"detection_mode": {
			Type:        hcl.TypeString,
			Description: "Detection mode for Java out of memory problem",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect Java out of memory problem",
			Optional:    true,
		},
	}
}

func (me *OutOfMemoryDetection) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"custom_thresholds": me.CustomThresholds,
		"detection_mode":    me.DetectionMode,
		"enabled":           me.Enabled,
	})
}

func (me *OutOfMemoryDetection) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"custom_thresholds": &me.CustomThresholds,
		"detection_mode":    &me.DetectionMode,
		"enabled":           &me.Enabled,
	})
}
