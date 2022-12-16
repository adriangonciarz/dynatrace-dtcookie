package workload

import "github.com/dtcookie/hcl"

type PendingPods struct {
	Configuration *PendingPodsConfig `json:"configuration"` // Alert if
	Enabled       bool               `json:"enabled"`       // Number of pods in pending phase
}

func (me *PendingPods) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"configuration": {
			Type:        hcl.TypeList,
			Description: "Alert if",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(PendingPodsConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Number of pods in pending phase",
			Optional:    true,
		},
	}
}

func (me *PendingPods) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"configuration": me.Configuration,
		"enabled":       me.Enabled,
	})
}

func (me *PendingPods) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"configuration": &me.Configuration,
		"enabled":       &me.Enabled,
	})
}
