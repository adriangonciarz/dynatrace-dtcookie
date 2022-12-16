package infrastructureaws

import "github.com/dtcookie/hcl"

// ElbHighConnectionErrorsDetectionThresholds. Alert if the condition is met in 3 out of 5 samples
type ElbHighConnectionErrorsDetectionThresholds struct {
	ConnectionErrorsPerMinute int `json:"connectionErrorsPerMinute"` // Number of backend connection errors is higher than
}

func (me *ElbHighConnectionErrorsDetectionThresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"connection_errors_per_minute": {
			Type:        hcl.TypeInt,
			Description: "Number of backend connection errors is higher than",
			Required:    true,
		},
	}
}

func (me *ElbHighConnectionErrorsDetectionThresholds) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"connection_errors_per_minute": me.ConnectionErrorsPerMinute,
	})
}

func (me *ElbHighConnectionErrorsDetectionThresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"connection_errors_per_minute": &me.ConnectionErrorsPerMinute,
	})
}
