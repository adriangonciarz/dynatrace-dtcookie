package externalwebrequest

import "github.com/dtcookie/hcl"

type ValueOverride struct {
	Value string `json:"value"`
}

func (me *ValueOverride) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"value": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
	}
}

func (me *ValueOverride) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"value": me.Value,
	})
}

func (me *ValueOverride) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"value": &me.Value,
	})
}
