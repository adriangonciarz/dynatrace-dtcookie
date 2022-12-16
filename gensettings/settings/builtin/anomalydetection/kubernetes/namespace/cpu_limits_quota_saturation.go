package namespace

import "github.com/dtcookie/hcl"

type CpuLimitsQuotaSaturation struct {
	Configuration *CpuLimitsQuotaSaturationConfig `json:"configuration"` // Alert if
	Enabled       bool                            `json:"enabled"`       // Detect CPU limits quota saturation
}

func (me *CpuLimitsQuotaSaturation) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"configuration": {
			Type:        hcl.TypeList,
			Description: "Alert if",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(CpuLimitsQuotaSaturationConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect CPU limits quota saturation",
			Optional:    true,
		},
	}
}

func (me *CpuLimitsQuotaSaturation) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"configuration": me.Configuration,
		"enabled":       me.Enabled,
	})
}

func (me *CpuLimitsQuotaSaturation) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"configuration": &me.Configuration,
		"enabled":       &me.Enabled,
	})
}
