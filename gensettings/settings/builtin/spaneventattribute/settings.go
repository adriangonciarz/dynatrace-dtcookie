package spaneventattribute

import "github.com/dtcookie/hcl"

type Settings struct {
	Key     string      `json:"key"`     // Key of the span event attribute to store
	Masking MaskingType `json:"masking"` // If this attribute contains confidential data, turn on masking to conceal its value from users
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"key": {
			Type:        hcl.TypeString,
			Description: "Key of the span event attribute to store",
			Required:    true,
		},
		"masking": {
			Type:        hcl.TypeString,
			Description: "If this attribute contains confidential data, turn on masking to conceal its value from users",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"key":     me.Key,
		"masking": me.Masking,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"key":     &me.Key,
		"masking": &me.Masking,
	})
}
