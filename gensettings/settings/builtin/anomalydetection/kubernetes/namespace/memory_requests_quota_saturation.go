package namespace

import "github.com/dtcookie/hcl"

type MemoryRequestsQuotaSaturation struct {
	Configuration *MemoryRequestsQuotaSaturationConfig `json:"configuration"` // Alert if
	Enabled       bool                                 `json:"enabled"`       // Detect memory requests quota saturation
}

func (me *MemoryRequestsQuotaSaturation) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"configuration": {
			Type:        hcl.TypeList,
			Description: "Alert if",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(MemoryRequestsQuotaSaturationConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect memory requests quota saturation",
			Optional:    true,
		},
	}
}

func (me *MemoryRequestsQuotaSaturation) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"configuration": me.Configuration,
		"enabled":       me.Enabled,
	})
}

func (me *MemoryRequestsQuotaSaturation) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"configuration": &me.Configuration,
		"enabled":       &me.Enabled,
	})
}
