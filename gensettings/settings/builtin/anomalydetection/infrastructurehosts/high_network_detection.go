package infrastructurehosts

import "github.com/dtcookie/hcl"

type HighNetworkDetection struct {
	CustomThresholds *HighNetworkDetectionThresholds `json:"customThresholds"`
	DetectionMode    DetectionMode                   `json:"detectionMode"` // Detection mode for high network utilization
	Enabled          bool                            `json:"enabled"`       // Detect high network utilization
}

func (me *HighNetworkDetection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"custom_thresholds": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(HighNetworkDetectionThresholds).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"detection_mode": {
			Type:        hcl.TypeString,
			Description: "Detection mode for high network utilization",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect high network utilization",
			Optional:    true,
		},
	}
}

func (me *HighNetworkDetection) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"custom_thresholds": me.CustomThresholds,
		"detection_mode":    me.DetectionMode,
		"enabled":           me.Enabled,
	})
}

func (me *HighNetworkDetection) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"custom_thresholds": &me.CustomThresholds,
		"detection_mode":    &me.DetectionMode,
		"enabled":           &me.Enabled,
	})
}
