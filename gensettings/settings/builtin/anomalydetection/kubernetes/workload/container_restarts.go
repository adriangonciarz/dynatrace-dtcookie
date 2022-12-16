package workload

import "github.com/dtcookie/hcl"

type ContainerRestarts struct {
	Configuration *ContainerRestartsConfig `json:"configuration"` // Alert if
	Enabled       bool                     `json:"enabled"`       // Detect container restarts
}

func (me *ContainerRestarts) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"configuration": {
			Type:        hcl.TypeList,
			Description: "Alert if",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ContainerRestartsConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect container restarts",
			Optional:    true,
		},
	}
}

func (me *ContainerRestarts) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"configuration": me.Configuration,
		"enabled":       me.Enabled,
	})
}

func (me *ContainerRestarts) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"configuration": &me.Configuration,
		"enabled":       &me.Enabled,
	})
}
