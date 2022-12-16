package services

import "github.com/dtcookie/hcl"

type ResponseTimeAutoAll struct {
	DegradationMilliseconds float64 `json:"degradationMilliseconds"` // Absolute threshold
	DegradationPercent      float64 `json:"degradationPercent"`      // Relative threshold
}

func (me *ResponseTimeAutoAll) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"degradation_milliseconds": {
			Type:        hcl.TypeFloat,
			Description: "Absolute threshold",
			Required:    true,
		},
		"degradation_percent": {
			Type:        hcl.TypeFloat,
			Description: "Relative threshold",
			Required:    true,
		},
	}
}

func (me *ResponseTimeAutoAll) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"degradation_milliseconds": me.DegradationMilliseconds,
		"degradation_percent":      me.DegradationPercent,
	})
}

func (me *ResponseTimeAutoAll) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"degradation_milliseconds": &me.DegradationMilliseconds,
		"degradation_percent":      &me.DegradationPercent,
	})
}
