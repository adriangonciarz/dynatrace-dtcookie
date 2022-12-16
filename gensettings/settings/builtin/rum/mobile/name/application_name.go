package name

import "github.com/dtcookie/hcl"

type ApplicationName struct {
	Name string `json:"name"`
}

func (me *ApplicationName) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "no documentation available",
			Required:    true,
		},
	}
}

func (me *ApplicationName) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"name": me.Name,
	})
}

func (me *ApplicationName) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"name": &me.Name,
	})
}
