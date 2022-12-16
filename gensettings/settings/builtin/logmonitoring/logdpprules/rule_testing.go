package logdpprules

import "github.com/dtcookie/hcl"

type RuleTesting struct {
	SampleLog string `json:"sampleLog"` // Sample log in JSON format.
}

func (me *RuleTesting) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"sample_log": {
			Type:        hcl.TypeString,
			Description: "Sample log in JSON format.",
			Required:    true,
		},
	}
}

func (me *RuleTesting) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"sample_log": me.SampleLog,
	})
}

func (me *RuleTesting) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"sample_log": &me.SampleLog,
	})
}
