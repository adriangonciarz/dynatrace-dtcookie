package infrastructurehosts

import "github.com/dtcookie/hcl"

type NetworkErrorsDetectionThresholds struct {
	ErrorsPercentage int              `json:"errorsPercentage"` // Receive/transmit error packet percentage threshold
	EventThresholds  *EventThresholds `json:"eventThresholds"`
	TotalPacketsRate int              `json:"totalPacketsRate"` // Total packets rate threshold
}

func (me *NetworkErrorsDetectionThresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"errors_percentage": {
			Type:        hcl.TypeInt,
			Description: "Receive/transmit error packet percentage threshold",
			Required:    true,
		},
		"event_thresholds": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(EventThresholds).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"total_packets_rate": {
			Type:        hcl.TypeInt,
			Description: "Total packets rate threshold",
			Required:    true,
		},
	}
}

func (me *NetworkErrorsDetectionThresholds) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"errors_percentage":  me.ErrorsPercentage,
		"event_thresholds":   me.EventThresholds,
		"total_packets_rate": me.TotalPacketsRate,
	})
}

func (me *NetworkErrorsDetectionThresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"errors_percentage":  &me.ErrorsPercentage,
		"event_thresholds":   &me.EventThresholds,
		"total_packets_rate": &me.TotalPacketsRate,
	})
}
