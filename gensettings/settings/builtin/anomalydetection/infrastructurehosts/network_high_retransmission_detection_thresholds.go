package infrastructurehosts

import "github.com/dtcookie/hcl"

type NetworkHighRetransmissionDetectionThresholds struct {
	EventThresholds                     *EventThresholds `json:"eventThresholds"`
	RetransmissionRatePercentage        int              `json:"retransmissionRatePercentage"`        // Retransmission rate threshold
	RetransmittedPacketsNumberPerMinute int              `json:"retransmittedPacketsNumberPerMinute"` // Number of retransmitted packets threshold
}

func (me *NetworkHighRetransmissionDetectionThresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"event_thresholds": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(EventThresholds).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"retransmission_rate_percentage": {
			Type:        hcl.TypeInt,
			Description: "Retransmission rate threshold",
			Required:    true,
		},
		"retransmitted_packets_number_per_minute": {
			Type:        hcl.TypeInt,
			Description: "Number of retransmitted packets threshold",
			Required:    true,
		},
	}
}

func (me *NetworkHighRetransmissionDetectionThresholds) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"event_thresholds":                        me.EventThresholds,
		"retransmission_rate_percentage":          me.RetransmissionRatePercentage,
		"retransmitted_packets_number_per_minute": me.RetransmittedPacketsNumberPerMinute,
	})
}

func (me *NetworkHighRetransmissionDetectionThresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"event_thresholds":                        &me.EventThresholds,
		"retransmission_rate_percentage":          &me.RetransmissionRatePercentage,
		"retransmitted_packets_number_per_minute": &me.RetransmittedPacketsNumberPerMinute,
	})
}
