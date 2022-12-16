package infrastructureaws

import "github.com/dtcookie/hcl"

// RdsHighCpuDetectionThresholds. Alert if the condition is met in 3 out of 5 samples
type RdsHighCpuDetectionThresholds struct {
	CpuUsage float64 `json:"cpuUsage"` // CPU usage is higher than
}

func (me *RdsHighCpuDetectionThresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"cpu_usage": {
			Type:        hcl.TypeFloat,
			Description: "CPU usage is higher than",
			Required:    true,
		},
	}
}

func (me *RdsHighCpuDetectionThresholds) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"cpu_usage": me.CpuUsage,
	})
}

func (me *RdsHighCpuDetectionThresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"cpu_usage": &me.CpuUsage,
	})
}
