package infrastructurevmware

import "github.com/dtcookie/hcl"

// SlowPhysicalStorageDetectionThresholds. Alert if **any** condition is met in 4 out of 5 samples
type SlowPhysicalStorageDetectionThresholds struct {
	AvgReadWriteLatency  int `json:"avgReadWriteLatency"`  // Read/write latency is higher than
	PeakReadWriteLatency int `json:"peakReadWriteLatency"` // Peak value for read/write latency is higher than
}

func (me *SlowPhysicalStorageDetectionThresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"avg_read_write_latency": {
			Type:        hcl.TypeInt,
			Description: "Read/write latency is higher than",
			Required:    true,
		},
		"peak_read_write_latency": {
			Type:        hcl.TypeInt,
			Description: "Peak value for read/write latency is higher than",
			Required:    true,
		},
	}
}

func (me *SlowPhysicalStorageDetectionThresholds) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"avg_read_write_latency":  me.AvgReadWriteLatency,
		"peak_read_write_latency": me.PeakReadWriteLatency,
	})
}

func (me *SlowPhysicalStorageDetectionThresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"avg_read_write_latency":  &me.AvgReadWriteLatency,
		"peak_read_write_latency": &me.PeakReadWriteLatency,
	})
}
