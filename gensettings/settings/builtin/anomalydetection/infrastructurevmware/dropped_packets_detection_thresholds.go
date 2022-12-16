package infrastructurevmware

import "github.com/dtcookie/hcl"

// DroppedPacketsDetectionThresholds. Alert if the condition is met in 3 out of 5 samples
type DroppedPacketsDetectionThresholds struct {
	DroppedPacketsPerSecond int `json:"droppedPacketsPerSecond"` // Receive/transmit dropped packets rate on NIC is higher than
}

func (me *DroppedPacketsDetectionThresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"dropped_packets_per_second": {
			Type:        hcl.TypeInt,
			Description: "Receive/transmit dropped packets rate on NIC is higher than",
			Required:    true,
		},
	}
}

func (me *DroppedPacketsDetectionThresholds) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"dropped_packets_per_second": me.DroppedPacketsPerSecond,
	})
}

func (me *DroppedPacketsDetectionThresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"dropped_packets_per_second": &me.DroppedPacketsPerSecond,
	})
}
