package cluster

import "github.com/dtcookie/hcl"

type CpuRequestsSaturation struct {
	Configuration *CpuRequestsSaturationConfig `json:"configuration"` // Alert if
	Enabled       bool                         `json:"enabled"`       // Detect CPU requests saturation
}

func (me *CpuRequestsSaturation) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"configuration": {
			Type:        hcl.TypeList,
			Description: "Alert if",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(CpuRequestsSaturationConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect CPU requests saturation",
			Optional:    true,
		},
	}
}

func (me *CpuRequestsSaturation) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"configuration": me.Configuration,
		"enabled":       me.Enabled,
	})
}

func (me *CpuRequestsSaturation) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"configuration": &me.Configuration,
		"enabled":       &me.Enabled,
	})
}
