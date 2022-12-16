package node

import "github.com/dtcookie/hcl"

type PodsSaturation struct {
	Configuration *PodsSaturationConfig `json:"configuration"` // Alert if
	Enabled       bool                  `json:"enabled"`       // Number of running pods in percent of the node's maximum pod capacity
}

func (me *PodsSaturation) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"configuration": {
			Type:        hcl.TypeList,
			Description: "Alert if",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(PodsSaturationConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Number of running pods in percent of the node's maximum pod capacity",
			Optional:    true,
		},
	}
}

func (me *PodsSaturation) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"configuration": me.Configuration,
		"enabled":       me.Enabled,
	})
}

func (me *PodsSaturation) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"configuration": &me.Configuration,
		"enabled":       &me.Enabled,
	})
}
