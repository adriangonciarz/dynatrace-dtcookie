package infrastructurehosts

import "github.com/dtcookie/hcl"

type NetworkDroppedPacketsDetection struct {
	CustomThresholds *NetworkDroppedPacketsDetectionThresholds `json:"customThresholds"` // Alert if the dropped packet percentage is higher than the specified threshold **and** the total packets rate is higher than the defined threshold for the defined amount of samples
	DetectionMode    DetectionMode                             `json:"detectionMode"`    // Detection mode for high number of dropped packets
	Enabled          bool                                      `json:"enabled"`          // Detect high number of dropped packets
}

func (me *NetworkDroppedPacketsDetection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"custom_thresholds": {
			Type:        hcl.TypeList,
			Description: "Alert if the dropped packet percentage is higher than the specified threshold **and** the total packets rate is higher than the defined threshold for the defined amount of samples",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(NetworkDroppedPacketsDetectionThresholds).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"detection_mode": {
			Type:        hcl.TypeString,
			Description: "Detection mode for high number of dropped packets",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect high number of dropped packets",
			Optional:    true,
		},
	}
}

func (me *NetworkDroppedPacketsDetection) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"custom_thresholds": me.CustomThresholds,
		"detection_mode":    me.DetectionMode,
		"enabled":           me.Enabled,
	})
}

func (me *NetworkDroppedPacketsDetection) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"custom_thresholds": &me.CustomThresholds,
		"detection_mode":    &me.DetectionMode,
		"enabled":           &me.Enabled,
	})
}
