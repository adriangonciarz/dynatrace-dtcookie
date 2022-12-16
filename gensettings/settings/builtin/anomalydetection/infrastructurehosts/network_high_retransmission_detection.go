package infrastructurehosts

import "github.com/dtcookie/hcl"

type NetworkHighRetransmissionDetection struct {
	CustomThresholds *NetworkHighRetransmissionDetectionThresholds `json:"customThresholds"` // Alert if the retransmission rate is higher than the specified threshold **and** the number of retransmitted packets is higher than the defined threshold for the defined amount of samples
	DetectionMode    DetectionMode                                 `json:"detectionMode"`    // Detection mode for high retransmission rate
	Enabled          bool                                          `json:"enabled"`          // Detect high retransmission rate
}

func (me *NetworkHighRetransmissionDetection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"custom_thresholds": {
			Type:        hcl.TypeList,
			Description: "Alert if the retransmission rate is higher than the specified threshold **and** the number of retransmitted packets is higher than the defined threshold for the defined amount of samples",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(NetworkHighRetransmissionDetectionThresholds).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"detection_mode": {
			Type:        hcl.TypeString,
			Description: "Detection mode for high retransmission rate",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect high retransmission rate",
			Optional:    true,
		},
	}
}

func (me *NetworkHighRetransmissionDetection) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"custom_thresholds": me.CustomThresholds,
		"detection_mode":    me.DetectionMode,
		"enabled":           me.Enabled,
	})
}

func (me *NetworkHighRetransmissionDetection) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"custom_thresholds": &me.CustomThresholds,
		"detection_mode":    &me.DetectionMode,
		"enabled":           &me.Enabled,
	})
}
