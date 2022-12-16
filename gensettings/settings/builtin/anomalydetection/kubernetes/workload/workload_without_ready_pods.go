package workload

import "github.com/dtcookie/hcl"

type WorkloadWithoutReadyPods struct {
	Configuration *WorkloadWithoutReadyPodsConfig `json:"configuration"` // Alert if
	Enabled       bool                            `json:"enabled"`       // Detect workloads without ready pods
}

func (me *WorkloadWithoutReadyPods) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"configuration": {
			Type:        hcl.TypeList,
			Description: "Alert if",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(WorkloadWithoutReadyPodsConfig).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Detect workloads without ready pods",
			Optional:    true,
		},
	}
}

func (me *WorkloadWithoutReadyPods) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"configuration": me.Configuration,
		"enabled":       me.Enabled,
	})
}

func (me *WorkloadWithoutReadyPods) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"configuration": &me.Configuration,
		"enabled":       &me.Enabled,
	})
}
