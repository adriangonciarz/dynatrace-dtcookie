package services

import "github.com/dtcookie/hcl"

type ResponseTimeFixedAll struct {
	DegradationMilliseconds float64 `json:"degradationMilliseconds"` // Alert if the response time degrades beyond this many ms within an observation period of 5 minutes
}

func (me *ResponseTimeFixedAll) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"degradation_milliseconds": {
			Type:        hcl.TypeFloat,
			Description: "Alert if the response time degrades beyond this many ms within an observation period of 5 minutes",
			Required:    true,
		},
	}
}

func (me *ResponseTimeFixedAll) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"degradation_milliseconds": me.DegradationMilliseconds,
	})
}

func (me *ResponseTimeFixedAll) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"degradation_milliseconds": &me.DegradationMilliseconds,
	})
}
