package infrastructureaws

import "github.com/dtcookie/hcl"

// RdsHighMemoryDetectionThresholds. Alert if **both** conditions is met in 3 out of 5 samples
type RdsHighMemoryDetectionThresholds struct {
	FreeMemory float64 `json:"freeMemory"` // Freeable memory is lower than
	SwapUsage  float64 `json:"swapUsage"`  // Swap usage is higher than
}

func (me *RdsHighMemoryDetectionThresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"free_memory": {
			Type:        hcl.TypeFloat,
			Description: "Freeable memory is lower than",
			Required:    true,
		},
		"swap_usage": {
			Type:        hcl.TypeFloat,
			Description: "Swap usage is higher than",
			Required:    true,
		},
	}
}

func (me *RdsHighMemoryDetectionThresholds) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"free_memory": me.FreeMemory,
		"swap_usage":  me.SwapUsage,
	})
}

func (me *RdsHighMemoryDetectionThresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"free_memory": &me.FreeMemory,
		"swap_usage":  &me.SwapUsage,
	})
}
