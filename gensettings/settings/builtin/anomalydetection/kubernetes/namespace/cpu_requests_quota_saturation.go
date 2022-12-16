package namespace

import "github.com/dtcookie/hcl"

type CpuRequestsQuotaSaturation struct {
	Configuration *CpuRequestsQuotaSaturationConfig `json:"configuration"` // Alert if
	Enabled       bool                              `json:"enabled"`       // Detect CPU requests quota saturation
}

func (me *CpuRequestsQuotaSaturation) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"configuration": {
			Type:        hcl.TypeList,
			Description: "Alert if",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(CpuRequestsQuotaSaturationConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect CPU requests quota saturation",
			Optional:    true,
		},
	}
}

func (me *CpuRequestsQuotaSaturation) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"configuration": me.Configuration,
		"enabled":       me.Enabled,
	})
}

func (me *CpuRequestsQuotaSaturation) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"configuration": &me.Configuration,
		"enabled":       &me.Enabled,
	})
}
