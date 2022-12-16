package namespace

import "github.com/dtcookie/hcl"

type PodsQuotaSaturation struct {
	Configuration *PodsQuotaSaturationConfig `json:"configuration"` // Alert if
	Enabled       bool                       `json:"enabled"`       // Detect pods quota saturation
}

func (me *PodsQuotaSaturation) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"configuration": {
			Type:        hcl.TypeList,
			Description: "Alert if",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(PodsQuotaSaturationConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect pods quota saturation",
			Optional:    true,
		},
	}
}

func (me *PodsQuotaSaturation) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"configuration": me.Configuration,
		"enabled":       me.Enabled,
	})
}

func (me *PodsQuotaSaturation) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"configuration": &me.Configuration,
		"enabled":       &me.Enabled,
	})
}
