package infrastructurevmware

import "github.com/dtcookie/hcl"

// UndersizedStorageDetectionThresholds. Alert if **any** condition is met in 3 out of 5 samples
type UndersizedStorageDetectionThresholds struct {
	AverageQueueCommandLatency int `json:"averageQueueCommandLatency"` // Average queue command latency is higher than
	PeakQueueCommandLatency    int `json:"peakQueueCommandLatency"`    // Peak queue command latency is higher than
}

func (me *UndersizedStorageDetectionThresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"average_queue_command_latency": {
			Type:        hcl.TypeInt,
			Description: "Average queue command latency is higher than",
			Required:    true,
		},
		"peak_queue_command_latency": {
			Type:        hcl.TypeInt,
			Description: "Peak queue command latency is higher than",
			Required:    true,
		},
	}
}

func (me *UndersizedStorageDetectionThresholds) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"average_queue_command_latency": me.AverageQueueCommandLatency,
		"peak_queue_command_latency":    me.PeakQueueCommandLatency,
	})
}

func (me *UndersizedStorageDetectionThresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"average_queue_command_latency": &me.AverageQueueCommandLatency,
		"peak_queue_command_latency":    &me.PeakQueueCommandLatency,
	})
}
