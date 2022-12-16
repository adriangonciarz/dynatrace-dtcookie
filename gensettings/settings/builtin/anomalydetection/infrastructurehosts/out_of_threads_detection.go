package infrastructurehosts

import "github.com/dtcookie/hcl"

type OutOfThreadsDetection struct {
	CustomThresholds *OutOfThreadsDetectionThresholds `json:"customThresholds"`
	DetectionMode    DetectionMode                    `json:"detectionMode"` // Detection mode for Java out of threads problem
	Enabled          bool                             `json:"enabled"`       // Detect Java out of threads problem
}

func (me *OutOfThreadsDetection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"custom_thresholds": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(OutOfThreadsDetectionThresholds).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"detection_mode": {
			Type:        hcl.TypeString,
			Description: "Detection mode for Java out of threads problem",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect Java out of threads problem",
			Optional:    true,
		},
	}
}

func (me *OutOfThreadsDetection) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"custom_thresholds": me.CustomThresholds,
		"detection_mode":    me.DetectionMode,
		"enabled":           me.Enabled,
	})
}

func (me *OutOfThreadsDetection) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"custom_thresholds": &me.CustomThresholds,
		"detection_mode":    &me.DetectionMode,
		"enabled":           &me.Enabled,
	})
}
