package services

import "github.com/dtcookie/hcl"

type ResponseTimeFixedSlowest struct {
	SlowestDegradationMilliseconds float64 `json:"slowestDegradationMilliseconds"` // Alert if the response time of the slowest 10% degrades beyond this many ms within an observation period of 5 minutes
}

func (me *ResponseTimeFixedSlowest) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"slowest_degradation_milliseconds": {
			Type:        hcl.TypeFloat,
			Description: "Alert if the response time of the slowest 10% degrades beyond this many ms within an observation period of 5 minutes",
			Required:    true,
		},
	}
}

func (me *ResponseTimeFixedSlowest) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"slowest_degradation_milliseconds": me.SlowestDegradationMilliseconds,
	})
}

func (me *ResponseTimeFixedSlowest) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"slowest_degradation_milliseconds": &me.SlowestDegradationMilliseconds,
	})
}
