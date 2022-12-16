package remote

import "github.com/dtcookie/hcl"

type Settings struct {
	PerformanceProfile PerformanceProfile `json:"performanceProfile"` // Select performance profile for Extension Execution Controller
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"performance_profile": {
			Type:        hcl.TypeString,
			Description: "Select performance profile for Extension Execution Controller",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"performance_profile": me.PerformanceProfile,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"performance_profile": &me.PerformanceProfile,
	})
}
