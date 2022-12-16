package namespace

import "github.com/dtcookie/hcl"

type MemoryLimitsQuotaSaturation struct {
	Configuration *MemoryLimitsQuotaSaturationConfig `json:"configuration"` // Alert if
	Enabled       bool                               `json:"enabled"`       // Detect memory limits quota saturation
}

func (me *MemoryLimitsQuotaSaturation) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"configuration": {
			Type:        hcl.TypeList,
			Description: "Alert if",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(MemoryLimitsQuotaSaturationConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect memory limits quota saturation",
			Optional:    true,
		},
	}
}

func (me *MemoryLimitsQuotaSaturation) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"configuration": me.Configuration,
		"enabled":       me.Enabled,
	})
}

func (me *MemoryLimitsQuotaSaturation) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"configuration": &me.Configuration,
		"enabled":       &me.Enabled,
	})
}
