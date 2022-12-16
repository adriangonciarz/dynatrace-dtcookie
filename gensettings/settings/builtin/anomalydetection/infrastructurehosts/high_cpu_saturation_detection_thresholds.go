package infrastructurehosts

import "github.com/dtcookie/hcl"

type HighCpuSaturationDetectionThresholds struct {
	CpuSaturation   int              `json:"cpuSaturation"` // Alert if the CPU usage is higher than this threshold for the defined amount of samples
	EventThresholds *EventThresholds `json:"eventThresholds"`
}

func (me *HighCpuSaturationDetectionThresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"cpu_saturation": {
			Type:        hcl.TypeInt,
			Description: "Alert if the CPU usage is higher than this threshold for the defined amount of samples",
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
	}
}

func (me *HighCpuSaturationDetectionThresholds) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"cpu_saturation":   me.CpuSaturation,
		"event_thresholds": me.EventThresholds,
	})
}

func (me *HighCpuSaturationDetectionThresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"cpu_saturation":   &me.CpuSaturation,
		"event_thresholds": &me.EventThresholds,
	})
}
