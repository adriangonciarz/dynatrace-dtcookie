package databases

import "github.com/dtcookie/hcl"

type ResponseTimeFixedSlowest struct {
	SlowestDegradationMilliseconds float64 `json:"slowestDegradationMilliseconds"` // Threshold
}

func (me *ResponseTimeFixedSlowest) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"slowest_degradation_milliseconds": {
			Type:        hcl.TypeFloat,
			Description: "Threshold",
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
