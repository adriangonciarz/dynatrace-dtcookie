package infrastructurehosts

import "github.com/dtcookie/hcl"

type HighNetworkDetectionThresholds struct {
	ErrorsPercentage int              `json:"errorsPercentage"` // Alert if sent/received traffic utilization is higher than this threshold for the defined amount of samples
	EventThresholds  *EventThresholds `json:"eventThresholds"`
}

func (me *HighNetworkDetectionThresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"errors_percentage": {
			Type:        hcl.TypeInt,
			Description: "Alert if sent/received traffic utilization is higher than this threshold for the defined amount of samples",
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

func (me *HighNetworkDetectionThresholds) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"errors_percentage": me.ErrorsPercentage,
		"event_thresholds":  me.EventThresholds,
	})
}

func (me *HighNetworkDetectionThresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"errors_percentage": &me.ErrorsPercentage,
		"event_thresholds":  &me.EventThresholds,
	})
}
