package monitoringrule

import "github.com/dtcookie/hcl"

type FlagDefaultTrue struct {
	Enabled bool `json:"enabled"`
}

func (me *FlagDefaultTrue) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "no documentation available",
			Optional:    true,
		},
	}
}

func (me *FlagDefaultTrue) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"enabled": me.Enabled,
	})
}

func (me *FlagDefaultTrue) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"enabled": &me.Enabled,
	})
}
