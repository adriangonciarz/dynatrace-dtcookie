package databases

import "github.com/dtcookie/hcl"

type ResponseTimeFixedAll struct {
	DegradationMilliseconds float64 `json:"degradationMilliseconds"` // Threshold
}

func (me *ResponseTimeFixedAll) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"degradation_milliseconds": {
			Type:        hcl.TypeFloat,
			Description: "Threshold",
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
