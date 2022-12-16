package rumweb

import "github.com/dtcookie/hcl"

type ResponseTimeAutoSlowest struct {
	SlowestDegradationMilliseconds float64 `json:"slowestDegradationMilliseconds"` // Absolute threshold
	SlowestDegradationPercent      float64 `json:"slowestDegradationPercent"`      // Relative threshold
}

func (me *ResponseTimeAutoSlowest) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"slowest_degradation_milliseconds": {
			Type:        hcl.TypeFloat,
			Description: "Absolute threshold",
			Required:    true,
		},
		"slowest_degradation_percent": {
			Type:        hcl.TypeFloat,
			Description: "Relative threshold",
			Required:    true,
		},
	}
}

func (me *ResponseTimeAutoSlowest) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"slowest_degradation_milliseconds": me.SlowestDegradationMilliseconds,
		"slowest_degradation_percent":      me.SlowestDegradationPercent,
	})
}

func (me *ResponseTimeAutoSlowest) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"slowest_degradation_milliseconds": &me.SlowestDegradationMilliseconds,
		"slowest_degradation_percent":      &me.SlowestDegradationPercent,
	})
}
