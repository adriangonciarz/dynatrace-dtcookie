package workload

import "github.com/dtcookie/hcl"

type Settings struct {
	ContainerRestarts        *ContainerRestarts        `json:"containerRestarts"`
	PendingPods              *PendingPods              `json:"pendingPods"`
	WorkloadWithoutReadyPods *WorkloadWithoutReadyPods `json:"workloadWithoutReadyPods"`
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"container_restarts": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(ContainerRestarts).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"pending_pods": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(PendingPods).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"workload_without_ready_pods": {
			Type:        hcl.TypeList,
			Description: "no documentation available",
			Required:    true,
			Elem:        &hcl.Resource{Schema: new(WorkloadWithoutReadyPods).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"container_restarts":          me.ContainerRestarts,
		"pending_pods":                me.PendingPods,
		"workload_without_ready_pods": me.WorkloadWithoutReadyPods,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"container_restarts":          &me.ContainerRestarts,
		"pending_pods":                &me.PendingPods,
		"workload_without_ready_pods": &me.WorkloadWithoutReadyPods,
	})
}
