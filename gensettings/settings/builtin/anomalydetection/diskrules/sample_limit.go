package diskrules

import "github.com/dtcookie/hcl"

type SampleLimit struct {
	Samples          int `json:"samples"`          // .. within the last
	ViolatingSamples int `json:"violatingSamples"` // Minimum number of violating samples
}

func (me *SampleLimit) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"samples": {
			Type:        hcl.TypeInt,
			Description: ".. within the last",
			Required:    true,
		},
		"violating_samples": {
			Type:        hcl.TypeInt,
			Description: "Minimum number of violating samples",
			Required:    true,
		},
	}
}

func (me *SampleLimit) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"samples":           me.Samples,
		"violating_samples": me.ViolatingSamples,
	})
}

func (me *SampleLimit) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"samples":           &me.Samples,
		"violating_samples": &me.ViolatingSamples,
	})
}
