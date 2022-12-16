package logsongrailactivate

import "github.com/dtcookie/hcl"

type Settings struct {
	Activated bool `json:"activated"` // Activate logs powered by Grail.
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"activated": {
			Type:        hcl.TypeBool,
			Description: "Activate logs powered by Grail.",
			Optional:    true,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"activated": me.Activated,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"activated": &me.Activated,
	})
}
