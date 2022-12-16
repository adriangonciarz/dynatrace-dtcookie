package infrastructureaws

import "github.com/dtcookie/hcl"

// RdsHighWriteReadLatencyDetectionThresholds. Alert if the condition is met in 3 out of 5 samples
type RdsHighWriteReadLatencyDetectionThresholds struct {
	ReadWriteLatency int `json:"readWriteLatency"` // Read/write latency is higher than
}

func (me *RdsHighWriteReadLatencyDetectionThresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"read_write_latency": {
			Type:        hcl.TypeInt,
			Description: "Read/write latency is higher than",
			Required:    true,
		},
	}
}

func (me *RdsHighWriteReadLatencyDetectionThresholds) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"read_write_latency": me.ReadWriteLatency,
	})
}

func (me *RdsHighWriteReadLatencyDetectionThresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"read_write_latency": &me.ReadWriteLatency,
	})
}
