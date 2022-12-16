package infrastructurehosts

import "github.com/dtcookie/hcl"

type OutOfMemoryDetectionThresholds struct {
	EventThresholds             *StrictEventThresholds `json:"eventThresholds"`
	OutOfMemoryExceptionsNumber int                    `json:"outOfMemoryExceptionsNumber"` // Alert if the number of Java out-of-memory exceptions is at least this value
}

func (me *OutOfMemoryDetectionThresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"event_thresholds": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(StrictEventThresholds).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"out_of_memory_exceptions_number": {
			Type:        hcl.TypeInt,
			Description: "Alert if the number of Java out-of-memory exceptions is at least this value",
			Required:    true,
		},
	}
}

func (me *OutOfMemoryDetectionThresholds) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"event_thresholds":                me.EventThresholds,
		"out_of_memory_exceptions_number": me.OutOfMemoryExceptionsNumber,
	})
}

func (me *OutOfMemoryDetectionThresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"event_thresholds":                &me.EventThresholds,
		"out_of_memory_exceptions_number": &me.OutOfMemoryExceptionsNumber,
	})
}
