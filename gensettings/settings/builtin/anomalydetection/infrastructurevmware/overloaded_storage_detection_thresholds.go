package infrastructurevmware

import "github.com/dtcookie/hcl"

// OverloadedStorageDetectionThresholds. Alert if the condition is met in 3 out of 5 samples
type OverloadedStorageDetectionThresholds struct {
	CommandAbortsNumber int `json:"commandAbortsNumber"` // Number of command aborts is higher than
}

func (me *OverloadedStorageDetectionThresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"command_aborts_number": {
			Type:        hcl.TypeInt,
			Description: "Number of command aborts is higher than",
			Required:    true,
		},
	}
}

func (me *OverloadedStorageDetectionThresholds) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"command_aborts_number": me.CommandAbortsNumber,
	})
}

func (me *OverloadedStorageDetectionThresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"command_aborts_number": &me.CommandAbortsNumber,
	})
}
