package infrastructurehosts

import "github.com/dtcookie/hcl"

type NetworkDroppedPacketsDetectionThresholds struct {
	DroppedPacketsPercentage int              `json:"droppedPacketsPercentage"` // Receive/transmit dropped packet percentage threshold
	EventThresholds          *EventThresholds `json:"eventThresholds"`
	TotalPacketsRate         int              `json:"totalPacketsRate"` // Total packets rate threshold
}

func (me *NetworkDroppedPacketsDetectionThresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"dropped_packets_percentage": {
			Type:        hcl.TypeInt,
			Description: "Receive/transmit dropped packet percentage threshold",
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

func (me *NetworkDroppedPacketsDetectionThresholds) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"dropped_packets_percentage": me.DroppedPacketsPercentage,
		"event_thresholds":           me.EventThresholds,
		"total_packets_rate":         me.TotalPacketsRate,
	})
}

func (me *NetworkDroppedPacketsDetectionThresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"dropped_packets_percentage": &me.DroppedPacketsPercentage,
		"event_thresholds":           &me.EventThresholds,
		"total_packets_rate":         &me.TotalPacketsRate,
	})
}
