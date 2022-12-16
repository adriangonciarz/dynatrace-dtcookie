package infrastructurehosts

import "github.com/dtcookie/hcl"

type HighMemoryDetection struct {
	CustomThresholds *HighMemoryDetectionThresholds `json:"customThresholds"` // Alert if **both** the memory usage and the memory page fault rate thresholds are exceeded on Windows or on Unix systems
	DetectionMode    DetectionMode                  `json:"detectionMode"`    // Detection mode for high memory usage
	Enabled          bool                           `json:"enabled"`          // Detect high memory usage on host
}

func (me *HighMemoryDetection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"custom_thresholds": {
			Type:        hcl.TypeList,
			Description: "Alert if **both** the memory usage and the memory page fault rate thresholds are exceeded on Windows or on Unix systems",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(HighMemoryDetectionThresholds).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"detection_mode": {
			Type:        hcl.TypeString,
			Description: "Detection mode for high memory usage",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect high memory usage on host",
			Optional:    true,
		},
	}
}

func (me *HighMemoryDetection) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"custom_thresholds": me.CustomThresholds,
		"detection_mode":    me.DetectionMode,
		"enabled":           me.Enabled,
	})
}

func (me *HighMemoryDetection) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"custom_thresholds": &me.CustomThresholds,
		"detection_mode":    &me.DetectionMode,
		"enabled":           &me.Enabled,
	})
}
