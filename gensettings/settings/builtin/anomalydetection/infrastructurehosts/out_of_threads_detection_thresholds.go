package infrastructurehosts

import "github.com/dtcookie/hcl"

type OutOfThreadsDetectionThresholds struct {
	EventThresholds              *StrictEventThresholds `json:"eventThresholds"`
	OutOfThreadsExceptionsNumber int                    `json:"outOfThreadsExceptionsNumber"` // Alert if the number of Java out-of-threads exceptions is at least this value
}

func (me *OutOfThreadsDetectionThresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"event_thresholds": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(StrictEventThresholds).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"out_of_threads_exceptions_number": {
			Type:        hcl.TypeInt,
			Description: "Alert if the number of Java out-of-threads exceptions is at least this value",
			Required:    true,
		},
	}
}

func (me *OutOfThreadsDetectionThresholds) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"event_thresholds":                 me.EventThresholds,
		"out_of_threads_exceptions_number": me.OutOfThreadsExceptionsNumber,
	})
}

func (me *OutOfThreadsDetectionThresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"event_thresholds":                 &me.EventThresholds,
		"out_of_threads_exceptions_number": &me.OutOfThreadsExceptionsNumber,
	})
}
