package cluster

import "github.com/dtcookie/hcl"

type MemoryRequestsSaturation struct {
	Configuration *MemoryRequestsSaturationConfig `json:"configuration"` // Alert if
	Enabled       bool                            `json:"enabled"`       // Detect memory requests saturation
}

func (me *MemoryRequestsSaturation) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"configuration": {
			Type:        hcl.TypeList,
			Description: "Alert if",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(MemoryRequestsSaturationConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect memory requests saturation",
			Optional:    true,
		},
	}
}

func (me *MemoryRequestsSaturation) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"configuration": me.Configuration,
		"enabled":       me.Enabled,
	})
}

func (me *MemoryRequestsSaturation) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"configuration": &me.Configuration,
		"enabled":       &me.Enabled,
	})
}
