package infrastructurehosts

import "github.com/dtcookie/hcl"

type HighGcActivityDetection struct {
	CustomThresholds *HighGcActivityDetectionThresholds `json:"customThresholds"` // Alert if the GC time **or** the GC suspension is exceeded
	DetectionMode    DetectionMode                      `json:"detectionMode"`    // Detection mode for high GC activity
	Enabled          bool                               `json:"enabled"`          // You may also configure high GC activity alerting for .NET processes on [extensions events page](/#settings/anomalydetection/extensionevents).
}

func (me *HighGcActivityDetection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"custom_thresholds": {
			Type:        hcl.TypeList,
			Description: "Alert if the GC time **or** the GC suspension is exceeded",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(HighGcActivityDetectionThresholds).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"detection_mode": {
			Type:        hcl.TypeString,
			Description: "Detection mode for high GC activity",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "You may also configure high GC activity alerting for .NET processes on [extensions events page](/#settings/anomalydetection/extensionevents).",
			Optional:    true,
		},
	}
}

func (me *HighGcActivityDetection) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"custom_thresholds": me.CustomThresholds,
		"detection_mode":    me.DetectionMode,
		"enabled":           me.Enabled,
	})
}

func (me *HighGcActivityDetection) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"custom_thresholds": &me.CustomThresholds,
		"detection_mode":    &me.DetectionMode,
		"enabled":           &me.Enabled,
	})
}
