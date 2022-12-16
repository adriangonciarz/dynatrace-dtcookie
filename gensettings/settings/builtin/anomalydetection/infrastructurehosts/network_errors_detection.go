package infrastructurehosts

import "github.com/dtcookie/hcl"

type NetworkErrorsDetection struct {
	CustomThresholds *NetworkErrorsDetectionThresholds `json:"customThresholds"` // Alert if the receive/transmit error packet percentage is higher than the specified threshold **and** the total packets rate is higher than the defined threshold for the defined amount of samples
	DetectionMode    DetectionMode                     `json:"detectionMode"`    // Detection mode for high number of network errors
	Enabled          bool                              `json:"enabled"`          // Detect high number of network errors
}

func (me *NetworkErrorsDetection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"custom_thresholds": {
			Type:        hcl.TypeList,
			Description: "Alert if the receive/transmit error packet percentage is higher than the specified threshold **and** the total packets rate is higher than the defined threshold for the defined amount of samples",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(NetworkErrorsDetectionThresholds).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"detection_mode": {
			Type:        hcl.TypeString,
			Description: "Detection mode for high number of network errors",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect high number of network errors",
			Optional:    true,
		},
	}
}

func (me *NetworkErrorsDetection) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"custom_thresholds": me.CustomThresholds,
		"detection_mode":    me.DetectionMode,
		"enabled":           me.Enabled,
	})
}

func (me *NetworkErrorsDetection) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"custom_thresholds": &me.CustomThresholds,
		"detection_mode":    &me.DetectionMode,
		"enabled":           &me.Enabled,
	})
}
